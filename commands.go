package main

import (
	"os"
	"fmt"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var COMMANDS map[string]cliCommand
func init() {
	COMMANDS = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "exits the REPL loop",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "displays command names and their uses",
			callback: commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	
	for _, c := range COMMANDS {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	
	fmt.Println()

	return nil
}
