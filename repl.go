package main

import (
	"strings"
	"unicode"
)

func cleanInput (text string) []string {
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
