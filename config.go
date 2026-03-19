package main

type Config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

func initConfig() Config {
	return Config{
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}
}
