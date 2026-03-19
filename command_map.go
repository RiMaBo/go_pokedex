package main

import (
	"internal/pokeapi"
	"fmt"
)

func commandMapForward(cfg *Config) error {
	locations, err := pokeapi.FetchLocations(cfg.nextLocationsURL);
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(cfg *Config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("You're on the first page")
	}

	locations, err := pokeapi.FetchLocations(cfg.prevLocationsURL);
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
