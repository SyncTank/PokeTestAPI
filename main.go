package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/SyncTank/pokedex/pokeAPI"
)

type cliCommand struct {
	name string
	description string
	callback func() error
	settings *config
}

type config struct {
	nextURL string
	pastURL string
}

var cliMap = map[string]cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a helping message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the past names of 20 location areas",
			callback: commandMapb,
		},
	}

func main(){
	fmt.Println(poke.Test)

	const input = "Pokedex >"
	scn := bufio.NewScanner(os.Stdin)
	fmt.Printf(input)
	for scn.Scan() {
		data := scn.Text()
		dataLow := strings.ToLower(data)
		dataList := strings.Fields(dataLow)

		cap, ok := cliMap[dataList[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			cap.callback()
		}

		fmt.Printf(input)
	}
}

func cleanInput(text string) [] string {
	strList := make([]string, 0);
	results := make([]string, 0);

	strList = strings.Fields(text)

	for _, value := range strList {
		results = append(results, value)
	}
	return results;
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandMap() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandMapb() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
