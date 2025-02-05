package main

import (
	"strings"
	"unicode"
	"bufio"
	"os"
	"fmt"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			text := scanner.Text()
			formattedInput := cleanInput(text)
			fmt.Println("Your command was:", formattedInput[0])
		}
	}
}

func cleanInput(text string) []string {
	var list []string
	var newWord string
	lowerStr := strings.ToLower(text)
	
	// iterating over every character in the input
	for _, char := range lowerStr {
		isWhiteSpace := unicode.IsSpace(char)
		if !isWhiteSpace {
			newWord += string(char)
		} else if isWhiteSpace && len(newWord) > 0 {
			list = append(list, newWord)
			newWord = ""
		}
	}
	
	// leftover word
	if len(newWord) > 0 {
		list = append(list, newWord)
	}

	return list
}
