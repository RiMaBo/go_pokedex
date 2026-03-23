package main

import (
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please specify a Pokemon to inspect")
	}

	pokemon := args[0]
	pokemonDetails, ok := cfg.pokedex[pokemon];
	if !ok {
		return fmt.Errorf("You have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemonDetails.Name)
	fmt.Printf("Height: %d\n", pokemonDetails.Height)
	fmt.Printf("Weight: %d\n", pokemonDetails.Weight)

	fmt.Println("Stats:")
	for _, stats := range pokemonDetails.Stats {
		fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}

	fmt.Println("Types:")
	for _, types := range pokemonDetails.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}

	return nil
}
