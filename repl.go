package main

import (
	"fmt"
	"log"
	"strings"
	"os"
	"bufio"
)

func startRepl() {
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
		command, ok := getCommands()[cleanedInputText[0]]
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

// Returns slice of words from test string split on blank space
func cleanInput(text string) []string {
	cleanedText := strings.ToLower(text)
	cleanedText = strings.TrimSpace(cleanedText)

	result := strings.Split(cleanedText, " ")
	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Shows usage guide",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows map",
			callback:    commandMap,
		},
	}
}
