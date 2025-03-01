package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/jms-guy/pokedexproject/internal/pokeapi"
	"github.com/jms-guy/pokedexproject/internal/pokecache"
)

type cliCommand struct {	//Struct for user input commands in the cli
	name		string
	description	string
	callback	func(c *pokeapi.Client, cd *pokeapi.ConfigData, cache *pokecache.Cache) error
}

var commandRegistry map[string]cliCommand	//Declaration of Command Registry

func commandHelp(c *pokeapi.Client, cd *pokeapi.ConfigData, cache *pokecache.Cache) error {	//Help command function
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *pokeapi.Client, cd *pokeapi.ConfigData, cache *pokecache.Cache) error {	//Map command function
	areaResults, err := c.GetLocationAreas(cache, cd.Next)
	if err != nil {
		return fmt.Errorf("error getting area location data: %w", err)
	}
	cd.Next = areaResults.Next
	cd.Previous = areaResults.Previous
	cd.Results = areaResults.Results
	for _, result := range cd.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(c *pokeapi.Client, cd *pokeapi.ConfigData, cache *pokecache.Cache) error {	//Map command function to go backwards
	if cd.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areaResults, err := c.GetLocationAreas(cache, cd.Previous)
	if err != nil {
		return fmt.Errorf("error getting area location data: %w", err)
	}
	cd.Next = areaResults.Next
	cd.Previous = areaResults.Previous
	cd.Results = areaResults.Results
	for _, result := range cd.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExit(c *pokeapi.Client, cd *pokeapi.ConfigData, cache *pokecache.Cache) error {	//Exit command function
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(s string) []string {	//Cleans user input string and returns first word in a lowercase state
	lowerS := strings.ToLower(s)
	results := strings.Fields(lowerS)
	return results
}

func init() {	//Initialization of command registry
	commandRegistry = map[string]cliCommand{	
		"help": {
			name:	"help",
			description:	"Displays a help message",
			callback:	commandHelp,
		},
		"map":	{
			name:	"map",
			description: "Displays 20 area locations in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name:	"mapb",
			description: "Displays the previous 20 area locations in the Pokemon world",
			callback: commandMapb,
		},
		"exit":	{
			name:	"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
	}
}