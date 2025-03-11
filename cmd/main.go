package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type State int

const (
	New State = iota + 1
	Done
	inProgress
)

type Task struct {
	id    int    `json:"id"`
	name  string `json:"nameTask"`
	state State  `json:"state"`
}

func main() {
	asciiArt := `   ____ _     ___       _____ ___  ___   _     _          _     ___ ____ _____ 
  / ___| |   |_ _|     |_   _/ _ \|   \ | |   | |        | |   |_ _/ ___|_   _|
 | |   | |    | |        | || | | | |\ \| |   | |        | |    | | |     | |  
 | |___| |___ | |        | || |_| | | \ | |___| |___    | |___ | | |___  | |  
  \____|_____|___|       |_|___/|_|__/|_____|_____|    |_____|___|\____| |_|  `

	fmt.Println(asciiArt)
	// call for open file json
	createJson()
	var dynamicInput []string
	for {
		var userInput string
		fmt.Print(">> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // Read the next line of input
		userInput = scanner.Text()
		userInput = strings.ToLower(userInput)
		// userInput = strings.TrimSpace(userInput)
		dynamicInput = strings.Fields(userInput)
		for i := 0; i < len(dynamicInput); i++ {
			fmt.Println(dynamicInput[i])
		}
		if userInput == "add" {

		}

		if userInput == "list" {
			// show all work and its state
		}

		if userInput == "exit" {
			//Check if there is any inProgress or new work
			// func
			// Maybe add hello goodbye stuff
			goodByeUser()
			// break
		}

	}
}

func createJson() {
	// Define the file name
	fileName := "output.json"

	// Check if the file exists
	_, err := os.Stat(fileName)
	if err == nil {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("Error deleting the file:", err)
			return
		}
		fmt.Println("File deleted successfully.")
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
}

func goodByeUser() {
	fmt.Println("Happy working")
}
