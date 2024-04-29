package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/utilities"
)

// generating the starting number on the ascii art file




func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input string>")
		return
	}

	input := os.Args[1] // user input
	// fmt.Println(len(input))

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	input = ascii.HandleBackspace(input)
	input = strings.ReplaceAll(string(input), "\\t", "   ") // handling the tab sequence
	file, err := os.ReadFile("standard.txt")
	ascii.ErrHandler(err)

	input = strings.ReplaceAll(input, "\\n", "\n")
	fileData := strings.Split(string(file), "\n")
	spString := strings.Split(input, "\n") // Split input by "\\n" to handle newline sequence
	// check for  illegal characters in the string
	inputParts, err := ascii.CheckIllegalChar(spString)
	ascii.ErrHandler(err)

	for _, part := range inputParts {
		if part == "" {
			fmt.Println() // Print a newline if the part is empty (i.e., consecutive newline characters)
			continue
		}
		for i := 0; i < 8; i++ { // this loop is responsible for the height of each character
			for _, char := range part { // iterates through each character of the current word
				startingIndex := ascii.GetStartingIndex(int(char)) // obtaining the starting position of the char
				fmt.Print(fileData[startingIndex-1+i])       // printing the character line by line
			}
			fmt.Println() // printing a new line after printing each line of the charcter
		}
	}
}
