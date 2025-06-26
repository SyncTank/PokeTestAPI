package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SyncTank/PokeTestAPI/pokeAPI"
	"github.com/SyncTank/PokeTestAPI/pokeCache"
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
var requestCache Cache

func main() {

	climap = getCommandList()

	const input = "Pokedex > "
	scn := bufio.NewScanner(os.Stdin)
	fmt.Printf(input)
	for scn.Scan() {
		data := scn.Text()
		dataLow := strings.ToLower(data)
		dataList := strings.Fields(dataLow)

		cap, ok := climap[dataList[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			cap.callback()
		}

		fmt.Printf(input)
	}
}

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
	locationMap, err := pokeAPI.GetLocation(climap["map"].settings.nextURL)
	if err != nil {
		fmt.Println("Request Failed %w\n", err)
		return err
	} else {
		//fmt.Println("Request Sucess")
		climap["map"].settings.pastURL = climap["map"].settings.nextURL
		climap["map"].settings.nextURL = locationMap.Next
		//fmt.Println(climap["map"].settings)
	}
	//fmt.Println(locationMap.Results)
	for i := range locationMap.Results {
		fmt.Printf(locationMap.Results[i].Name+" %T %T\n", locationMap.Results[i].Name, ([]byte)(locationMap.Results[i].Name))
		// This conversion is what to cache
		//dataItem := ([]byte)(locationMap.Results[i].Name)
		requestCache.AddCache()
	}
	return nil
}

func commandMapb() error {
	//fmt.Println("Showing the last next 20 Items!")
	locationMap, err := pokeAPI.GetLocation(climap["map"].settings.pastURL)
	if err != nil {
		fmt.Println("Request Failed: \n", err)
		return err
	} else {
		//fmt.Println("Request Sucess")
		climap["map"].settings.nextURL = climap["map"].settings.pastURL
		climap["map"].settings.pastURL = locationMap.Previous
		//fmt.Println(climap["map"].settings)
	}
	for i := range locationMap.Results {
		fmt.Println(locationMap.Results[i].Name)
	}
	return nil
}
