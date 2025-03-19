package main

import (
	"fmt"
	"os"
	"sort"
	"slices"
	"github.com/jms-guy/pokedexcli/internal/pokeapi"
	"github.com/jms-guy/pokedexcli/internal/filefunctions"
	"github.com/jms-guy/pokedexcli/internal/versionfunctions"
)

//Pokedex command functions//
//Map commands are currently disabled, right now I have no use for the random exploration functions
//To-do functions :  areas -> lists areas in current game version, pokemon -> lists all pokemon in version pokedex

type cliCommand struct {	//Struct for user input commands in the cli
	name		string
	description	string
	callback	func(app *PokedexApp, data pokeapi.APIResponse, agrs []string) error
}

var commandRegistry map[string]cliCommand	//Declaration of Command Registry

func commandCheckVersion(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Returns the current version set by the user
	if app.CurrVersion == ""  {
		fmt.Println("No version currently set.")
		return nil
	}
	fmt.Printf("Game version set to: %s\n", app.CurrVersion)
	return nil
}

func commandSetVersion(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Sets the game version that the Pokedex will use to parse response data
	if len(args) < 2 {
		fmt.Println("Please specify a version")
		return nil
	}
	version := args[1]
	
	_, err := ParseVersion(version)	//Checks version input against an enum to determine it's validity
	if err != nil {
		fmt.Printf("error parsing version input: %s\n", err)
		return nil
	}
	_, exists := app.Version[version]	//Check if version group data already in map
	if !exists {
		versionResults, err := app.Client.GetVersionGroup(app.Cache, "https://pokeapi.co/api/v2/version/"+version)	//If not, fetch it
		if err != nil {
			return err
		}
		app.Version[version] = versionResults	//Adds version details
	}
	app.CurrVersion = version
	fmt.Printf("Version set to %s\n", version)
	return nil
}

func commandVersions(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Versions command function, simply returns the list of strings in the Versions enum in versions.go, representing the supported game versions
	fmt.Println("Currently supported versions:")
	versions := make([]string, 0)	//Creates a string to place the version names into, to make sorting easier (map sorting sucks)
	for _, vName := range versionName {
		versions = append(versions, vName)
	}
	slices.Sort(versions)	//Sorts version names alphabetically, and prints them
	for i, version := range versions {
		if i == len(versions) - 1 {
			fmt.Printf("%s", version)
		} else {
			fmt.Printf("%s : ", version)
		}
	}
	fmt.Println("")
	return nil
}

func commandLoad(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Load command function, loads pokedex data from disk file into program
	filefunctions.LoadPokedex(app)
	return nil
}

func commandSave(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Save command function, saves current pokedex data into a file on disk
	if len(app.UserPokedex) == 0 {
		fmt.Println("You have no pokemon in your Pokedex.")
		return nil
	}
	err := filefunctions.SavePokedex(app)
	if err != nil {
		return fmt.Errorf("error saving pokedex data: %w", err)
	}
	fmt.Println("Pokedex saved!")
	return nil
}
func commandPokedex(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Pokedex command function, displays list of "caught" pokemon available for inspecting
	if len(app.UserPokedex) == 0 {
		fmt.Println("You have no pokemon in your Pokedex.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range app.UserPokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}

func commandInspect(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Inspect command function, returns data for the input pokemon
	pokemonName := args[1]
	pokemon, ok := app.UserPokedex[pokemonName];	//Check if pokemon has been caught
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

func commandCatch(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Catch command function, adding a pokemon to the pokedex for viewing
	if len(args) < 2 {
		return fmt.Errorf(("missing pokemon name or id number"))
	}
	pokemonName := args[1]

	pokemonData, err := app.Client.GetPokemonData(app.Cache, "https://pokeapi.co/api/v2/pokemon/"+pokemonName)	//Fetches pokeapi data
	if err != nil {
		return fmt.Errorf("error getting data for %s: %w", pokemonName, err)
	}
	
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")
	app.UserPokedex[pokemonName] = pokemonData	//Adds data to pokedex
	return nil
}

func commandHelp(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Help command function, returns list of commands in the command registry
	fmt.Println("Welcome to the Pokedex!\nAvailable commands:")
	keys := make([]string, 0, len(commandRegistry))
	for k := range commandRegistry {
		keys = append(keys, k)
	}
	sort.Strings(keys)		//Sorts the command registry, so the commands are always returned in order

	for _, k := range keys {
		cmd := commandRegistry[k]
		fmt.Printf(" %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandFind(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Find command function, returns list of game locations where the input pokemon can be found, in the version set by user
	if len(args) < 2 {
		fmt.Println("Missing pokemon name")
		return nil
	}
	pokemonName := args[1]
	gameVersion := app.CurrVersion

	if _, ok := data.(*pokeapi.EncounterAreas); !ok {
		return fmt.Errorf("command find requires EncounterAreas")
	}
	encounterData, err := app.Client.GetEncounterData(app.Cache, "https://pokeapi.co/api/v2/pokemon/"+pokemonName+"/encounters")	//Gets pokeapi data of encounter locations
	if err != nil {
		return fmt.Errorf("error retrieving %s encounter data: %w", pokemonName, err)
	}

	relevantEncounters := versionfunctions.VersionEncounters(encounterData, gameVersion)	//Sorts encounter data for only version relevant data

	if len(relevantEncounters) == 0 {
		fmt.Println("No location data for this Pokemon found in this version.")
	} else {
		fmt.Printf("%s locations in %s:\n", pokemonName, gameVersion)
	}

	for areaName, details := range relevantEncounters {
		fmt.Printf("--------------- %s ---------------\n", areaName)
		
    	seenEncounters := make(map[string]struct{})	//Unique encounters
		
    	for _, enDetails := range details {
        	for _, encounter := range enDetails.EncounterDetails {
				conditions := extractConditionNames(encounter.ConditionValues)	//Make conditionvalues much more readable
				if conditions == "" || conditions == "[]" {
					conditions = "None"
				}
				// Create a unique key for each encounter
				uniqueKey := fmt.Sprintf("%v-%v-%d-%d-%s",
					enDetails.MaxChance,      // Max chance as part of the key
					conditions, // Conditions as part of the key
					encounter.MinLevel,        // Min level as part of the key
					encounter.MaxLevel,        // Max level as part of the key
					encounter.Method.Name,     // Method name as part of the key
				)

				// Skip if we've already printed this encounter
				if _, exists := seenEncounters[uniqueKey]; exists {
					continue
				}

				// Mark this encounter as seen
				seenEncounters[uniqueKey] = struct{}{}


    	fmt.Printf("***** Chance: %d ***** Conditions: %v ***** Level: %d ***** Method: %s *****\n",
							enDetails.MaxChance,
							conditions,
							encounter.MaxLevel,
							encounter.Method.Name,
							)
			}
		}
	}
	return nil
}

func commandExplore(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Explore command function, listing pokemon available to be caught in the input area
	if len(args) < 2 {
		fmt.Println("missing area name")
		return nil
	}
	areaName := args[1]
	fmt.Printf("Exploring %s...\n", areaName)
	
	if ed, ok := data.(*pokeapi.LocationAreaDetails); ok {
		encounterData, err := app.Client.GetAreaExplorationData(app.Cache, "https://pokeapi.co/api/v2/location-area/"+areaName)	//Fetches pokeapi data of area
		if err != nil {
			return fmt.Errorf("error getting encounter details for area %s: %w", areaName, err)
		}
		ed.Name = encounterData.Name
		ed.PokemonEncounters = encounterData.PokemonEncounters
		fmt.Println("Found Pokemon:")	//Lists pokemon returned in response
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
		areaResults, err := app.Client.GetLocationAreas(app.Cache, configData.Next)
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
		areaResults, err := app.Client.GetLocationAreas(app.Cache, cd.Previous)
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

func commandExit(app *PokedexApp, data pokeapi.APIResponse, args []string) error {	//Exit command function, closes the program
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func init() {	//Initialization of command registry
	commandRegistry = map[string]cliCommand{
		"find":	{
			name: "find",
			description: "Returns a list of locations where the given pokemon can be found in set game version. Limited functionality in more recent versions. -> find _____",
			callback: commandFind,
		},	
		"catch":	{
			name:	"catch",
			description: "Add a pokemon to your pokedex -> catch _____",
			callback: commandCatch,
		},
		"help": {
			name:	"help",
			description:	"Displays a help message.",
			callback:	commandHelp,
		},
		"inspect": {
			name:	"inspect",
			description: "Shows details of a caught pokemon -> inspect _____",
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
			description:	"Exit the Pokedex.",
			callback:	commandExit,
		},
		"explore":	{
			name:	"explore",
			description:	"Explore a location for available pokemon to catch -> explore _____",
			callback:	commandExplore,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays names of all pokemon user has caught.",
			callback: commandPokedex,
		},
		"save": {
			name: "save",
			description: "Saves the current pokedex to a file. Only one save file is currently supported, if you save without loading a previous save, that save will be overwritten.",
			callback: commandSave,
		},
		"load": {
			name: "load",
			description: "Load saved Pokedex data from file. Loading data will not overwrite pokemon currently in Pokedex, it will add to it.",
			callback: commandLoad,
		},
		"set-version": {
			name: "set-version",
			description: "Sets the current pokedex version(red, blue, gold, violet, etc.) -> set-version _____",
			callback: commandSetVersion,
		},
		"check-version":	{
			name: "check-version",
			description: "Returns current Pokedex game version.",
			callback: commandCheckVersion,
		},
		"versions": {
			name: "versions",
			description: "Displays list of currently supported game versions.",
			callback: commandVersions,
		},
	}
}