package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SyncTank/pokedex/pokeAPI"
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

func main() {

	climap := getCommandList()

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
		} else if cap.name == "help" {
			commandHelp(climap)
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
	var cmap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a helping message",
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
			settings:    nil,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the past names of 20 location areas",
			callback:    commandMapb,
			settings:    nil,
		},
	}
	return cmap
}

func commandHelp(cmap map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, item := range cmap {
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
	fmt.Println("Showing the next 20 Items!")
	locationMap, err := pokeAPI.GetLocation(pokeAPI.Endpoint)
	if err != nil {
		fmt.Println("Request Failed")
		return err
	} else {
		fmt.Println("Request Sucess")
		fmt.Println(locationMap)
	}
	return nil
}

func commandMapb() error {
	fmt.Println("Showing the last next 20 Items!")
	return nil
}
