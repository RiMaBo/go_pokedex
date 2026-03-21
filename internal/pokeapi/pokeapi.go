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

func (c *Client) ExploreLocation(location string) (Location, error) {
	fullUrl := baseUrl + "/location-area/" + location

	if cached, ok := c.cache.Get(fullUrl); ok {
		var locationDetails Location
		if err:= json.Unmarshal(cached, &locationDetails); err != nil {
			return Location{}, err
		}

		return locationDetails, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var locationDetails Location
	if err:= json.Unmarshal(data, &locationDetails); err != nil {
		return Location{}, err
	}


	c.cache.Add(fullUrl, data)
	return locationDetails, nil
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

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
