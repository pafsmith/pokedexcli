package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Returns slice of words from test string split on blank space
func cleanInput(text string) []string {
	cleanedText := strings.ToLower(text)
	cleanedText = strings.TrimSpace(cleanedText)
	

	result := strings.Split(cleanedText, " ")
	return result
}


func commandExit() error {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
		return nil
	}


func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}




func main() {
	commands := map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
		"help": {
			name: "help",
			description: "Shows usage guide",
			callback: commandHelp,
		},
}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		inputText := scanner.Text()
		cleanedInputText := cleanInput(inputText)
		command, ok := commands[cleanedInputText[0]]
		if !ok {
			fmt.Println("Unknown command.")
			continue
		}
		err = command.callback()
		if err != nil {
			fmt.Println(err)
		}
		

	}
}
