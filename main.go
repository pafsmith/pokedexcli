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


func main() {
	
	
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
		fmt.Printf("Your command was: %s\n", cleanedInputText[0])
	}
}
