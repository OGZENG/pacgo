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
	return nil
}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		fmt.Println(line)
	}
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
	}
	return "", nil
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
		if input == "ESC" {
			break
		}

		// process movement

		// process collisions

		// check game over

		// Temp: break infinite loop
		break

		// repeat
	}
}
