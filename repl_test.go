package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {	
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "	hello world	",
			expected: []string{"hello", "world"},
		},
		{
			input: " - - - - ",
			expected: []string{"-", "-", "-", "-"},
		},
		{
			input: " ",
			expected: []string{},
		},
		{
			input: "",
			expected: []string{},
		},

	}
	
	// loop over all cases
	for n, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf(
				"Expected length: %d\nActual length: %d", 
				len(c.expected), 
				len(actual),
			)
			continue
		} else {
			fmt.Printf("Passed length case: %d\n", n)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(
					"Expected word: %s\nActual word: %s",
					expectedWord,
					word,
				)
			} else {
				fmt.Printf("Passed equivalency case: %d word: %d\n", n, i)
			} 
		}
	}
}
