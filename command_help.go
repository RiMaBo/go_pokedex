package main

import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Println("\nUsage:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}
