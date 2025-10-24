package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pafsmith/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.RespPokemon
}

func startRepl(cfg *config) {
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

		args := []string{}
		if len(cleanedInputText) > 1 {
			args = cleanedInputText[1:]
		}
		if !ok {
			fmt.Println("Unknown command.")
			continue
		}

		err = command.callback(cfg, args...)
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
	callback    func(*config, ...string) error
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
		"mapb": {
			name:        "mapb",
			description: "Show prevoious map page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
	}
}
