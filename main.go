package main

import (
	"internal/pokeapi"
	"time"
)

func main() {
	httpTimeout := 10 * time.Second
	cacheInterval := 5 * time.Minute
	pokeClient := pokeapi.NewClient(httpTimeout, cacheInterval)
	cfg := &Config{
		pokeApiClient: pokeClient,
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
