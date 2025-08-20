package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/SyncTank/PokeTestAPI/pokeAPI"
	"github.com/SyncTank/PokeTestAPI/pokeCache"
)

type cliCommands struct {
	name        string
	description string
	callback    func() error
	settings    *config
}

type config struct {
	nextURL    string
	currentURL string
	pastURL    string
	argv       string
}

var dex pokeAPI.PokeDex
var climap map[string]cliCommands
var requestCache *cache.Cache

func cleanInput(text string) []string {
	strList := make([]string, 0)
	results := make([]string, 0)

	strList = strings.Fields(text)

	for _, value := range strList {
		results = append(results, value)
	}
	return results
}

func getCommandList() map[string]cliCommands {
	var nConfig = config{
		nextURL:    pokeAPI.Endpoint + pokeAPI.LoctionEndpoint,
		currentURL: pokeAPI.Endpoint + pokeAPI.LoctionEndpoint,
		pastURL:    "",
		argv:       "",
	}
	var result = map[string]cliCommands{
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
		"explore": {
			name:        "explore",
			description: "Displays list of pokemon in a certain location",
			callback:    commandExplore,
			settings:    &nConfig,
		},
		"catch": {
			name:        "catch",
			description: "Attmepts to catch a pokemon",
			callback:    commandCatch,
			settings:    &nConfig,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints the name, height, weight, stats and type(s) of the Pokemon",
			callback:    commandInspect,
			settings:    &nConfig,
		},
	}
	return result
}

func commandInspect() error {
	fmt.Println("Throwing a Pokeball at " + "...")

	return nil
}

func commandCatch() error {
	pokemon, err := pokeAPI.GetPokemon(pokeAPI.Endpoint + pokeAPI.PokemonEndpoint + climap["catch"].settings.argv)
	isCaught := false
	if err != nil {
		fmt.Println("Request Failed", err)
		return err
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		ratio := ((48 / ((math.Sqrt(float64(pokemon.Base_experience))) - (99 / float64(pokemon.Weight*pokemon.Height)) + 1)) + 1 + (10 / (math.Sqrt(float64(pokemon.Base_experience))))) * 10
		chance := r.Intn(50)
		//fmt.Println("Random : ", ratio, chance)
		if chance+int(ratio) > 50 {
			isCaught = true
		} else {
			isCaught = false
		}
	}
	fmt.Println("Throwing a Pokeball at " + climap["catch"].settings.argv + "...")
	if isCaught {
		fmt.Println(climap["catch"].settings.argv + " was caught!")
		if dex.Pokedex == nil {
			dex.Pokedex = make(map[string]pokeAPI.Pokemon)
		}
		dex.Pokedex[climap["catch"].settings.argv] = pokemon
	} else {
		fmt.Println(climap["catch"].settings.argv + " escaped!")
	}
	return nil
}

func commandExplore() error {
	fmt.Println("Exploring " + climap["explore"].settings.argv + "...")
	fmt.Println("Found Pokemon:")
	cachemap, ok := requestCache.GetCache(climap["explore"].settings.argv)
	if !ok {
		pokemonMap, err := pokeAPI.GetPokemons(pokeAPI.Endpoint + pokeAPI.LoctionEndpoint + climap["explore"].settings.argv)
		if err != nil {
			fmt.Println("Request Failed", err)
			return err
		} else {
			results := ""
			for i := range pokemonMap.PokemonEncounters {
				fmt.Println(" - " + pokemonMap.PokemonEncounters[i].Pokemon.Name)
				results += pokemonMap.PokemonEncounters[i].Pokemon.Name + "\\n"
			}
			requestCache.AddCache(climap["explore"].settings.argv, ([]byte)(results))
		}
	} else {
		items := strings.Split(string(cachemap), "\\n")
		for i := range len(items) - 1 {
			fmt.Println(" - " + items[i])
		}
	}
	return nil
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
	//fmt.Println(climap["map"].settings.nextURL)
	cachemap, ok := requestCache.GetCache(climap["map"].settings.nextURL)
	if !ok {
		locationMap, err := pokeAPI.GetLocation(climap["map"].settings.nextURL)
		if err != nil {
			fmt.Println("Request Failed", err)
			return err
		} else {
			results := ""
			for i := range locationMap.Results {
				fmt.Println(locationMap.Results[i].Name)
				results += locationMap.Results[i].Name + "\\n"
			}

			requestCache.AddCache(climap["map"].settings.nextURL, ([]byte)(results))
			climap["map"].settings.pastURL = climap["map"].settings.currentURL
			climap["map"].settings.currentURL = climap["map"].settings.nextURL
			climap["map"].settings.nextURL = locationMap.Next
		}
	} else {
		items := strings.Split(string(cachemap), "\\n")
		for i := range len(items) - 1 {
			fmt.Println(items[i])
		}
		climap["map"].settings.pastURL = climap["map"].settings.currentURL
		climap["map"].settings.currentURL = climap["map"].settings.nextURL
	}
	return nil
}

func commandMapb() error { // show last 20 items
	//fmt.Println(climap["map"].settings.pastURL)
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
				results += locationMap.Results[i].Name + "\\n"
			}
			requestCache.AddCache(climap["map"].settings.pastURL, ([]byte)(results))
			climap["map"].settings.nextURL = climap["map"].settings.currentURL
			climap["map"].settings.currentURL = climap["map"].settings.pastURL
			climap["map"].settings.pastURL = locationMap.Previous
		}
	} else {
		items := strings.Split(string(cachemap), "\\n")
		for i := range len(items) - 1 {
			fmt.Println(items[i])
		}
		climap["map"].settings.nextURL = climap["map"].settings.currentURL
		climap["map"].settings.currentURL = climap["map"].settings.pastURL
	}
	return nil
}
