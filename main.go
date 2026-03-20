package main

import (
	"internal/pokeapi"
	"time"
)

func main() {
	httpTimeout := 5 * time.Second
	cacheInterval := 5 * time.Minute
	pokeClient := pokeapi.NewClient(httpTimeout, cacheInterval)
	cfg := &Config{
		pokeApiClient: pokeClient,
	}

	startRepl(cfg)
}
