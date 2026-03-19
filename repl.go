package main

import (
	"bufio";
	"fmt";
	"os";
	"strings";
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) < 1 {
			continue
		}

		fmt.Println("Your command was:", input[0])
	}
}
