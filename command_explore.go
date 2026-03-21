package main

import "fmt"

func commandExplore(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please provide a location name")
	}

	location := args[0]
	fmt.Printf("Exploring %s...\n", location)

	locationDetails, err := cfg.pokeApiClient.ExploreLocation(location);
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounters := range locationDetails.PokemonEncounters {
		fmt.Println(" - " + encounters.Pokemon.Name)
	}

	return nil
}
