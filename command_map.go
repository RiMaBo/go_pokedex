package main

import "fmt"

func commandMapForward(cfg *Config, args ...string) error {
	locations, err := cfg.pokeApiClient.FetchLocations(cfg.nextLocationsURL);
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

func commandMapBack(cfg *Config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("You're on the first page")
	}

	locations, err := cfg.pokeApiClient.FetchLocations(cfg.prevLocationsURL);
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
