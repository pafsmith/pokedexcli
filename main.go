package main

import (
	"fmt"
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
	fmt.Println("Hello, World!")
}
