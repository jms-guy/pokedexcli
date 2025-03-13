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
	UserPokedex	map[string]pokeapi.PokemonDetails	//Pokedex
}

func main() {
	app := &PokedexApp{	//Create shared state struct
		Client: 	pokeapi.NewClient(),		
		Cache: 		pokecache.NewCache(10 * time.Second),
		UserPokedex: make(map[string]pokeapi.PokemonDetails),
	}
	scanner := bufio.NewScanner(os.Stdin)	//Creates scanner for text input

	configData := &pokeapi.ConfigData{}	//Create data structures for storage use
    locationAreaData := &pokeapi.LocationAreaDetails{}
    pokemonData := &pokeapi.PokemonDetails{}

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
		var data pokeapi.APIResponse	//Creates empty interface to assign data structure to based on command
		if command.name == "explore" {
			data = locationAreaData
		} else if (command.name == "map") || (command.name == "mapb") {
			data = configData
		} else {
			data = pokemonData
		}		
		err := command.callback(app, data, userInput)	//Executes command from registry
		if err != nil {
			fmt.Printf("Returned error: %v", err)
			continue
		}
		
	}
}


func (app *PokedexApp) GetUserPokedex() map[string]pokeapi.PokemonDetails {	//Interface function to return pokedex data to filefunctions
	return app.UserPokedex
}
