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
	nextURL    string
	currentURL string
	pastURL    string
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
		nextURL:    pokeAPI.Endpoint,
		currentURL: pokeAPI.Endpoint,
		pastURL:    "",
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

func commandMap() error { // show next 20 items
	fmt.Println(climap["map"].settings.nextURL)
	cachemap, ok := requestCache.GetCache(climap["map"].settings.nextURL)
	if !ok {
		locationMap, err := pokeAPI.GetLocation(climap["map"].settings.nextURL)
		if err != nil {
			fmt.Println("Request Failed %w\n", err)
			return err
		} else {
			results := ""
			for i := range locationMap.Results {
				fmt.Println(locationMap.Results[i].Name)
				results += locationMap.Results[i].Name + "\n"
			}
			// P C N | 0 1 1 | 1 1 2 | 1 2 3 | 2 3 4

			requestCache.AddCache(climap["map"].settings.nextURL, ([]byte)(results))
			climap["map"].settings.pastURL = climap["map"].settings.currentURL
			climap["map"].settings.currentURL = climap["map"].settings.nextURL
			climap["map"].settings.nextURL = locationMap.Next
		}
	} else {
		fmt.Println("CACHED!")
		fmt.Println(cachemap)
		for _, j := range cachemap {
			fmt.Println(string(j))
		}
		climap["map"].settings.pastURL = climap["map"].settings.nextURL
	}
	return nil
}

func commandMapb() error { // show last 20 items
	fmt.Println(climap["map"].settings.pastURL)
	cachemap, ok := requestCache.GetCache(climap["map"].settings.pastURL)
	if !ok {
		locationMap, err := pokeAPI.GetLocation(climap["map"].settings.pastURL)
		if err != nil {
			fmt.Println("Request Failed: \n", err)
			return err
		} else {
			results := ""
			for i := range locationMap.Results {
				fmt.Println(locationMap.Results[i].Name)
				results += locationMap.Results[i].Name + "\n"
			}
			requestCache.AddCache(climap["map"].settings.pastURL, ([]byte)(results))
			climap["map"].settings.nextURL = climap["map"].settings.currentURL
			climap["map"].settings.currentURL = climap["map"].settings.pastURL
			climap["map"].settings.pastURL = locationMap.Previous
		}
	} else {
		fmt.Println("CACHED!")
		fmt.Println(string(cachemap))
		items := strings.Split(string(cachemap), "\n")
		fmt.Println(items)
		for j := range items {
			fmt.Println(string(j))
		}
		climap["map"].settings.nextURL = climap["map"].settings.currentURL
		climap["map"].settings.currentURL = climap["map"].settings.pastURL
	}
	return nil
}
