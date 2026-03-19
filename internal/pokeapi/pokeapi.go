package pokeapi

import (
	"net/http"
	"encoding/json"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
	timeout = 5 * time.Second
)

func FetchLocations(url *string) (ShallowLocations, error) {
	pageUrl := baseURL + "/location-area/"
	if url != nil {
		pageUrl = *url
	}

	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return ShallowLocations{}, err
	}

	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return ShallowLocations{}, err
	}
	defer res.Body.Close()

	var locations ShallowLocations
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return ShallowLocations{}, err
	}

	return locations, nil
}

type ShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
