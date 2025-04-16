package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const Endpoint string = "https://pokeapi.co/api/v2/location-area"

type Area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Area `json:"results"`
}

func GetLocation(url string) (Location, error) {

	res, err := http.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("Error creating request: %w", err)
	}
	defer res.Body.Close()

	var result Location
	data, err := io.ReadAll(res.Body)
	if err = json.Unmarshal(data, &result); err != nil {
		return Location{}, err
	}

	return result, nil
}
