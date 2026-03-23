package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please specify a Pokemon to catch")
	}

	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemonDetails, err := cfg.pokeApiClient.CatchPokemon(pokemon);
	if err != nil {
		return err
	}

	baseExperience := float64(pokemonDetails.BaseExperience)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	caught := math.Min(r.Float64() * baseExperience, r.Float64() * baseExperience);
	if (caught > baseExperience / 2) {
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemon] = pokemonDetails
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
