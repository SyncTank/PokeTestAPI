package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/SyncTank/PokeTestAPI/pokeCache"
)

func main() {
	const baseTime = 7 * time.Millisecond
	climap := getCommandList()
	pokecache := cache.NewCache(baseTime)
	fmt.Println(pokecache)

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
