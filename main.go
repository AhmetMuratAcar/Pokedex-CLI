package main

import (
	"fmt"
	"bufio"
	"os"
)

func main () {
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
