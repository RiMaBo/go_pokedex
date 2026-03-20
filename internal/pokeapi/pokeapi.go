package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

func (c *Client) FetchLocations(pageUrl *string) (ShallowLocations, error) {
	fullUrl := baseUrl + "/location-area/"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	if cached, ok := c.cache.Get(fullUrl); ok {
		var locations ShallowLocations
		if err := json.Unmarshal(cached, &locations); err != nil {
			return ShallowLocations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return ShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ShallowLocations{}, err
	}

	var locations ShallowLocations
	if err := json.Unmarshal(data, &locations); err != nil {
		return ShallowLocations{}, err
	}

	c.cache.Add(fullUrl, data)
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
