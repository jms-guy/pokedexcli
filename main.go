package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/jms-guy/pokedexcli/internal/pokeapi"
	"github.com/jms-guy/pokedexcli/internal/pokecache"
)

type PokedexApp struct {		//Encapsulated shared state struct for functional refinement
	Client		*pokeapi.Client	//Http Client
	Cache		*pokecache.Cache	//Data Cache
	CurrVersion	string
	Version		map[string]pokeapi.VersionGroup			//Stores version of Pokedex for further filtering of data
	UserPokedex	map[string]pokeapi.PokemonDetails	//Pokedex
}

func main() {
	app := &PokedexApp{	//Create shared state struct
		Client: 	pokeapi.NewClient(),		
		Cache: 		pokecache.NewCache(2 * time.Hour),
		CurrVersion: "",
		Version:	make(map[string]pokeapi.VersionGroup),
		UserPokedex: make(map[string]pokeapi.PokemonDetails),
	}
	scanner := bufio.NewScanner(os.Stdin)	//Creates scanner for text input

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())	//Gets user input text
		if len(userInput) == 0 {
			continue
		}
		command, ok := commandRegistry[userInput[0]]	//Searches registry for command
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
	
		err := command.callback(app, userInput)	//Executes command from registry
		if err != nil {
			fmt.Printf("Returned error: %v", err)
			continue
		}		
	}
}


func (app *PokedexApp) GetUserPokedex() *map[string]pokeapi.PokemonDetails {	//Interface function to return pokedex data to filefunctions
	return &app.UserPokedex
}
