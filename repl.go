package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) < 1 {
			continue
		}

		commandName := input[0]
		command, exists := getCommands()[commandName]

		if exists {
			if err := command.callback(cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: '%s'. Type 'help' or a list of commands.\n", commandName)
		}
	}
}
