package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const Endpoint string = "https://pokeapi.co/api/v2/location-area"

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

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

func GetPokemons(url string) (LocationArea, error) {
	res, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error creating request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return LocationArea{}, fmt.Errorf("HTTP failed %s", string(body))
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error Reading data")
	}

	var result LocationArea
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error Decoding Data %s", string(data))
	}

	fmt.Println(result)

	return result, nil
}
