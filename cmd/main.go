package main

import (
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Name  string
	State string
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
	for true {
		var userInput string
		fmt.Print(">> ")
		fmt.Scanln(&userInput)
		userInput = strings.ToLower(userInput)
		// if len(os.Args) < 2 {
		// 	fmt.Println("Usage:", os.Args[0], "<argument>")
		// 	os.Exit(1)
		// }
		// fmt.Println("Program name: ", os.Args[0])
		// for i, arg := range os.Args[1:] {
		// 	fmt.Printf("Arg %d: %s\n", i+1, arg)
		// }
		if userInput == "exit" {
			// Maybe add hello goodbye stuff
			fmt.Println("")
			break
		}

	}
}

func createJson() {
	f, err := os.Create("output.json")
	if err != nil {
		err = os.Remove("output.json")
		f, err = os.Create("output.json")
		defer f.Close()
	}

	defer f.Close()
}
