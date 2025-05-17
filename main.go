package main

import (
	"fmt"
	"log"
	"os/exec"
)
import "os"
import "bufio"
import "github.com/danicat/simpleansi"

var maze []string

type sprite struct {
	row int
	col int
}

var player sprite

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row, col} //use {} to construct structure
			}
		}
	}
	return nil
}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Println("%c", chr)
			default:
				fmt.Println("")
			}
		}
		fmt.Println()
	}
	simpleansi.MoveCursor(player.row, player.col)
	fmt.Println("P")
	simpleansi.MoveCursor(len(maze)+1, 0)
}
func initialize() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin
	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate the cbreak Modul", err)
	}
}
func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin
	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to recover to cooked module", err)
	}
}
func readinput() (string, error) {
	buffer := make([]byte, 100)
	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}
	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "LEFT", nil
			case 'D':
				return "RIGHT", nil

			}
		}
	}
	return "", nil
}
func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol
	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}
	if maze[newRow][newCol] == '#' {
		newRow, newCol = oldRow, oldCol
	}
	return
}
func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)
}
func main() {
	// initialize game
	initialize()
	defer cleanup()

	// load resources
	err := loadMaze("maze01.txt")
	if err != nil {
		log.Println("failed to open: ", err)
		return
	}

	// game loop
	for {
		// update screen
		printScreen()
		// process input
		input, err := readinput()
		if err != nil {
			log.Print("error reading input", err)
			break
		}

		// process movement
		movePlayer(input)
		// process collisions

		// check game over
		if input == "ESC" {
			break
		}

		// repeat
	}
}
