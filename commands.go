package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/jms-guy/pokedexcli/internal/pokeapi"
	"github.com/jms-guy/pokedexcli/internal/pokecache"
)

type cliCommand struct {	//Struct for user input commands in the cli
	name		string
	description	string
	callback	func(c *pokeapi.Client, data pokeapi.APIResponse, cache *pokecache.Cache, agrs []string) error
}

var commandRegistry map[string]cliCommand	//Declaration of Command Registry

func commandHelp(c *pokeapi.Client, data pokeapi.APIResponse, cache *pokecache.Cache, args []string) error {	//Help command function
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExplore(c *pokeapi.Client, data pokeapi.APIResponse, cache *pokecache.Cache, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing location area name")
	}
	
	pageURL := "https://pokeapi.co/api/v2/location-area/" + args[1]
	if ed, ok := data.(*pokeapi.LocationAreaDetails); ok {
		encounterData, err := c.GetAreaExplorationData(cache, pageURL)
		if err != nil {
			return fmt.Errorf("error getting encounter details for area %s: %w", args[1], err)
		}
		ed.Name = encounterData.Name
		ed.Encounters = encounterData.Encounters
		fmt.Println("Found Pokemon:")
		for _, encounter := range ed.Encounters {
			fmt.Println(encounter)
		}
	} else {
		return fmt.Errorf("command explore requires LocationAreaDetails")
	}
	return nil
}

func commandMap(c *pokeapi.Client, data pokeapi.APIResponse, cache *pokecache.Cache, args []string) error {	//Map command function
	if configData, ok := data.(*pokeapi.ConfigData); ok {
		areaResults, err := c.GetLocationAreas(cache, configData.Next)
		if err != nil {
			return fmt.Errorf("error getting area location data: %w", err)
		}
		configData.Next = areaResults.Next
		configData.Previous = areaResults.Previous
		configData.Results = areaResults.Results
		for _, result := range configData.Results {
			fmt.Println(result.Name)
		}
	} else {
		return fmt.Errorf("command map requires ConfigData")
	}
	return nil
}

func commandMapb(c *pokeapi.Client, data pokeapi.APIResponse, cache *pokecache.Cache, args []string) error {	//Map command function to go backwards
	if cd, ok := data.(*pokeapi.ConfigData); ok {
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
	} else {
		return fmt.Errorf("command mapb requires ConfigData")
	}
	return nil
}

func commandExit(c *pokeapi.Client, data pokeapi.APIResponse, cache *pokecache.Cache, args []string) error {	//Exit command function
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
		"explore":	{
			name:	"explore",
			description:	"Explore a location for available pokemon to catch",
			callback:	commandExplore,
		},
	}
}