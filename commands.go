package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/jms-guy/pokedexcli/internal/pokeapi"
	"github.com/jms-guy/pokedexcli/internal/catch_chance"
)

type cliCommand struct {	//Struct for user input commands in the cli
	name		string
	description	string
	callback	func(app *PokedexApp, data pokeapi.APIResponse, agrs []string) error
}

var commandRegistry map[string]cliCommand	//Declaration of Command Registry

func commandPokedex(app *PokedexApp, data pokeapi.APIResponse, args []string) error {
	if len(app.userPokedex) == 0 {
		fmt.Println("You have no pokemon in your Pokedex.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range app.userPokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}

func commandInspect(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Inspect command function
	pokemonName := args[1]
	pokemon, ok := app.userPokedex[pokemonName];	//Check if pokemon has been caught
	if !ok {
		fmt.Println("User has not caught this pokemon.")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)	//Display results
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		name := stat.Stat.Name
		val := stat.BaseStat
		fmt.Printf("  -%s: %d\n", name, val)
	}
	fmt.Println("Types:")
	for _, ptype := range pokemon.Types {
		name := ptype.Type.Name
		fmt.Printf("  - %s\n", name)
	}
	return nil
}

func commandCatch(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Catch command function
	if len(args) < 2 {
		return fmt.Errorf(("missing pokemon name or id number"))
	}
	pokemonName := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)


	pokemonData, err := app.client.GetPokemonData(app.cache, "https://pokeapi.co/api/v2/pokemon/"+pokemonName)
	if err != nil {
		return fmt.Errorf("error getting data for %s: %w", pokemonName, err)
	}
	expVal := pokemonData.BaseExperience
	if expVal == 0 {
		return fmt.Errorf("missing base experience value")
	}
	if catchchance.GetCatchBool(expVal) {	//Check catch chance
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		app.userPokedex[pokemonName] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func commandHelp(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Help command function
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExplore(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Explore command function
	if len(args) < 2 {
		fmt.Println("missing area name")
		return nil
	}
	areaName := args[1]
	fmt.Printf("Exploring %s...\n", areaName)
	
	if ed, ok := data.(*pokeapi.LocationAreaDetails); ok {
		encounterData, err := app.client.GetAreaExplorationData(app.cache, "https://pokeapi.co/api/v2/location-area/"+areaName)
		if err != nil {
			return fmt.Errorf("error getting encounter details for area %s: %w", areaName, err)
		}
		ed.Name = encounterData.Name
		ed.PokemonEncounters = encounterData.PokemonEncounters
		fmt.Println("Found Pokemon:")
		for _, pokemon := range ed.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
	} else {
		return fmt.Errorf("command explore requires LocationAreaDetails")
	}
	return nil
}

func commandMap(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Map command function
	if configData, ok := data.(*pokeapi.ConfigData); ok {
		areaResults, err := app.client.GetLocationAreas(app.cache, configData.Next)
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

func commandMapb(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Map command function to go backwards
	if cd, ok := data.(*pokeapi.ConfigData); ok {
		if cd.Previous == nil {
			fmt.Println("you're on the first page")
			return nil
		}
		areaResults, err := app.client.GetLocationAreas(app.cache, cd.Previous)
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

func commandExit(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Exit command function
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(s string) []string {	//Cleans user input string for command arguments use
	lowerS := strings.ToLower(s)
	results := strings.Fields(lowerS)
	return results
}

func init() {	//Initialization of command registry
	commandRegistry = map[string]cliCommand{	
		"catch":	{
			name:	"catch",
			description: "Add a pokemon to your pokedex",
			callback: commandCatch,
		},
		"help": {
			name:	"help",
			description:	"Displays a help message",
			callback:	commandHelp,
		},
		"inspect": {
			name:	"inspect",
			description: "Shows details of a caught pokemon",
			callback: commandInspect,
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
		"pokedex": {
			name: "pokedex",
			description: "Displays names of all pokemon user has caught",
			callback: commandPokedex,
		},
	}
}