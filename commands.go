package main

import (
	"fmt"
	"github.com/SyncTank/PokeTestAPI/pokeAPI"
	"github.com/SyncTank/PokeTestAPI/pokeCache"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	settings    *config
}

type config struct {
	nextURL string
	pastURL string
}

var climap map[string]cliCommand
var requestCache cache.Cache

func cleanInput(text string) []string {
	strList := make([]string, 0)
	results := make([]string, 0)

	strList = strings.Fields(text)

	for _, value := range strList {
		results = append(results, value)
	}
	return results
}

func getCommandList() map[string]cliCommand {
	var nConfig = config{
		nextURL: pokeAPI.Endpoint,
		pastURL: "",
	}
	var result = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a helping message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
			settings:    &nConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the past names of 20 location areas",
			callback:    commandMapb,
			settings:    &nConfig,
		},
	}
	return result
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, item := range climap {
		fmt.Printf("%s : %s\n", item.name, item.description)
	}
	fmt.Println("")
	return nil
}
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap() error {
	//fmt.Println("Showing the next 20 Items!")
	fmt.Println(climap["map"].settings.nextURL)
	cachemap, ok := requestCache.GetCache(climap["map"].settings.nextURL)
	if !ok {
		locationMap, err := pokeAPI.GetLocation(climap["map"].settings.nextURL)
		if err != nil {
			fmt.Println("Request Failed %w\n", err)
			return err
		} else {
			climap["map"].settings.pastURL = climap["map"].settings.nextURL
			climap["map"].settings.nextURL = locationMap.Next
		}
		results := ""
		for i := range locationMap.Results {
			fmt.Println(locationMap.Results[i].Name)
			results += locationMap.Results[i].Name
		}
		requestCache.AddCache(climap["map"].settings.nextURL, ([]byte)(results))
	} else {
		fmt.Println(cachemap)
	}
	return nil
}

func commandMapb() error {
	//fmt.Println("Showing the last next 20 Items!")
	fmt.Println(climap["map"].settings.pastURL)
	locationMap, err := pokeAPI.GetLocation(climap["map"].settings.pastURL)
	if err != nil {
		fmt.Println("Request Failed: \n", err)
		return err
	} else {
		climap["map"].settings.nextURL = climap["map"].settings.pastURL
		climap["map"].settings.pastURL = locationMap.Previous
	}
	results := ""
	for i := range locationMap.Results {
		fmt.Println(locationMap.Results[i].Name)
		results += locationMap.Results[i].Name
	}
	requestCache.AddCache("Mapb", ([]byte)(results))
	return nil
}
