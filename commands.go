package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"slices"
	"github.com/jms-guy/pokedexcli/internal/filefunctions"
	"github.com/jms-guy/pokedexcli/internal/versionfunctions"
)

//Pokedex command functions//
//To-do functions : update save function to store cache data, handle evolution data?(ugh) what about items?

type cliCommand struct {	//Struct for user input commands in the cli
	name		string
	description	string
	callback	func(app *PokedexApp, agrs []string) error
}

var commandRegistry map[string]cliCommand	//Declaration of Command Registry

func commandPokedex(app *PokedexApp, args []string) error {	//Command to list either all pokedexes in game version, or all pokemon in specified pokedex
	if app.CurrVersion == "" {
		fmt.Println("Please choose your game version using the 'set-version' command.")	//Checks for set game version
		return nil
	}
	group := app.Version[app.CurrVersion]

	if len(args) < 2 {
		fmt.Printf(" ***** Pokedexes in version %s ***** \n", capitalizeString(app.CurrVersion))	//If no pokedex specified, lists all in game version
		for _, pokedex := range group.Pokedexes {
			fmt.Printf(" *** %s ***\n", capitalizeString(pokedex.Name))
		}
		return nil
	}

	dex := args[1]
	dexData, err := app.Client.GetPokedexData(app.Cache, "https://pokeapi.co/api/v2/pokedex/"+dex)	//Gets data for specific dex
	if err != nil {
		return err
	}
	fmt.Printf(" ******** %s Pokedex ********\n", capitalizeString(dexData.Name))
	
	var pokemonSlice []string
	for _, pokemon := range dexData.PokemonEntries {
		pokeData := fmt.Sprintf("%d. %s", pokemon.EntryNumber, capitalizeString(pokemon.PokemonSpecies.Name))	//Lists all pokemon in a more presentable way
		pokemonSlice = append(pokemonSlice, pokeData)
		if len(pokemonSlice) == 10 {
			fmt.Println(strings.Join(pokemonSlice, " --- "))
			pokemonSlice = pokemonSlice[:0]
		}
	}
	if len(pokemonSlice) > 0 {
		fmt.Println(strings.Join(pokemonSlice, " --- "))	//Handles pokemon leftover in slice
	}
	return nil
}

func commandCheckVersion(app *PokedexApp, args []string) error {	//Returns the current version set by the user
	if app.CurrVersion == ""  {
		fmt.Println("No version currently set.")
		return nil
	}
	fmt.Printf("Game version set to: %s\n", capitalizeString(app.CurrVersion))
	return nil
}

func commandSetVersion(app *PokedexApp, args []string) error {	//Sets the game version that the Pokedex will use to parse response data
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
	fmt.Printf("Version set to %s\n", capitalizeString(version))
	return nil
}

func commandVersions(app *PokedexApp, args []string) error {	//Versions command function, simply returns the list of strings in the Versions enum in versions.go, representing the supported game versions
	fmt.Println("Currently supported versions:")
	versions := make([]string, 0)	//Creates a string to place the version names into, to make sorting easier (map sorting sucks)
	for _, vName := range versionName {
		versions = append(versions, vName)
	}
	slices.Sort(versions)	//Sorts version names alphabetically, and prints them
	for i, version := range versions {
		if i == len(versions) - 1 {
			fmt.Printf("%s", capitalizeString(version))
		} else {
			fmt.Printf("%s : ", capitalizeString(version))
		}
	}
	fmt.Println("")
	return nil
}

func commandLoad(app *PokedexApp, args []string) error {	//Load command function, loads pokedex data from disk file into program
	filefunctions.LoadPokedex(app)
	return nil
}

func commandSave(app *PokedexApp, args []string) error {	//Save command function, saves current pokedex data into a file on disk
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
func commandPokemon(app *PokedexApp, args []string) error {	//Pokedex command function, displays list of "caught" pokemon available for inspecting
	if len(app.UserPokedex) == 0 {
		fmt.Println("You have no pokemon in your Pokedex.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range app.UserPokedex {
		fmt.Printf(" - %s\n", capitalizeString(name))
	}
	return nil
}

func commandInspect(app *PokedexApp, args []string) error {	//Inspect command function, returns data for the input pokemon
	if len(args) < 2 {
		fmt.Println("Missing pokemon name or id number.")
		return nil
	}
	pokemonName := args[1]

	if _, ok := app.UserPokedex[pokemonName]; !ok {
		pokemonData, err := app.Client.GetPokemonData(app.Cache, "https://pokeapi.co/api/v2/pokemon/"+pokemonName)	//Fetches pokeapi data
		if err != nil {
			return fmt.Errorf("error getting data for %s: %w", pokemonName, err)
		}
		app.UserPokedex[pokemonName] = pokemonData	//Adds data to pokedex
	}
	pokemon := app.UserPokedex[pokemonName]

	fmt.Printf("Name: %s\n", capitalizeString(pokemon.Name))	//Display results
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		name := stat.Stat.Name
		val := stat.BaseStat
		fmt.Printf("  -%s: %d\n", capitalizeString(name), val)
	}
	fmt.Println("Types:")
	for _, ptype := range pokemon.Types {
		name := ptype.Type.Name
		fmt.Printf("  - %s\n", capitalizeString(name))
	}
	return nil
}

func commandHelp(app *PokedexApp, args []string) error {	//Help command function, returns list of commands in the command registry
	fmt.Println("Welcome to the Pokedex!\n Commands vary in effectiveness based on the game version entered. (Find command has limited use for Sw/Sh/Scar/Viol for example.)")
	fmt.Println(" Most commands are usable by their names, some take additional input, some require additional input. Specified by ***** ******.")
	fmt.Println("Available commands:")
	keys := make([]string, 0, len(commandRegistry))
	for k := range commandRegistry {
		keys = append(keys, k)
	}
	sort.Strings(keys)		//Sorts the command registry, so the commands are always returned in order

	for _, k := range keys {
		cmd := commandRegistry[k]
		fmt.Printf(" - %s: %s\n", capitalizeString(cmd.name), cmd.description)
	}
	return nil
}

func commandFind(app *PokedexApp, args []string) error {	//Find command function, returns list of game locations where the input pokemon can be found, in the version set by user
	if len(args) < 2 {
		fmt.Println("Missing pokemon name")
		return nil
	}
	gameVersion := app.CurrVersion	//Command defaults to current pokedex version
	if len(args) == 3 {	//If another version is specified in command, will use that instead
		if _, err := ParseVersion(args[2]); err != nil {
			fmt.Println("Unknown game version given")
			return nil
		} else {
			gameVersion = args[2]
		}
	} 
	pokemonName := args[1]

	encounterData, err := app.Client.GetEncounterData(app.Cache, "https://pokeapi.co/api/v2/pokemon/"+pokemonName+"/encounters")	//Gets pokeapi data of encounter locations
	if err != nil {
		return fmt.Errorf("error retrieving %s encounter data: %w", pokemonName, err)
	}

	relevantEncounters := versionfunctions.VersionEncounters(encounterData, gameVersion)	//Sorts encounter data for only version relevant data

	if len(relevantEncounters) == 0 {
		fmt.Println("No location data for this Pokemon found.")
	} else {
		fmt.Printf("%s locations in %s:\n", capitalizeString(pokemonName), capitalizeString(gameVersion))
	}

	for areaName, details := range relevantEncounters {
		fmt.Printf("--------------- %s ---------------\n", capitalizeString(areaName))
		
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
					capitalizeString(encounter.Method.Name),     // Method name as part of the key
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
							capitalizeString(encounter.Method.Name),
							)
			}
		}
	}
	return nil
}

func commandExplore(app *PokedexApp, args []string) error {	//Explore command function, listing pokemon available to be caught in the input area
	if len(args) < 2 {
		fmt.Println("missing area name")
		return nil
	}
	areaName := args[1]
	fmt.Printf("Exploring %s...\n", capitalizeString(areaName))
	
	encounterData, err := app.Client.GetAreaExplorationData(app.Cache, "https://pokeapi.co/api/v2/location-area/"+areaName)	//Fetches pokeapi data of area
	if err == nil {	//If argument is a location area, all is well, lists pokemon in location area
		fmt.Println("Found Pokemon:")	//Lists pokemon returned in response
		for _, pokemon := range encounterData.PokemonEncounters {
			fmt.Printf(" - %s\n", capitalizeString(pokemon.Pokemon.Name))
		}
		return nil
	} 
	//If not a location area, fall back and try for location data instead
	locationData, err := app.Client.GetFurtherExplorationData(app.Cache, "https://pokeapi.co/api/v2/location/"+areaName)
	if err != nil {
		fmt.Println("Invalid location name.")
		return nil
	}
	for _, locationArea := range locationData.Areas {
		fmt.Printf(" ***** %s *****\n", capitalizeString(locationArea.Name))
		encounterData, err := app.Client.GetAreaExplorationData(app.Cache, locationArea.URL)
		if err != nil {
			return fmt.Errorf("error getting location area data: %w", err)
		}
		for _, pokemon := range encounterData.PokemonEncounters {
			fmt.Printf(" - %s\n", capitalizeString(pokemon.Pokemon.Name))
		}
	}
	return nil
}

func commandMap(app *PokedexApp, args []string) error {	//Map command function, returns list of locations in region
	group := app.Version[app.CurrVersion]
	fmt.Printf("******** Areas in %s ********\n", capitalizeString(app.CurrVersion))

	for _, region := range group.Regions {
		regionData, err := app.Client.GetRegionData(app.Cache, region.URL)	//Gets regional data 
		if err != nil {
			return err
		}
		fmt.Printf(" ***** %s *****\n", capitalizeString(region.Name))

		var locationSlice []string
		for _, location := range regionData.Locations {	//Parses locations in region
			locationSlice = append(locationSlice, capitalizeString(location.Name))
			if len(locationSlice) == 10 {
				fmt.Println(strings.Join(locationSlice, " --- "))	//Once 10 locations have been parsed, print them to console in a nice, readable form
				locationSlice = locationSlice[:0]
			}
		}
		if len(locationSlice) > 0 {
			fmt.Println(strings.Join(locationSlice, " --- "))	//Handles leftover locations
		}
	}
	return nil
}

func commandExit(app *PokedexApp, args []string) error {	//Exit command function, closes the program
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func init() {	//Initialization of command registry
	commandRegistry = map[string]cliCommand{
		"pokedex":	{
			name: "pokedex",
			description: "If no additional input is given, will list all pokedex's in your current set game version. If a pokedex is specified, will list all pokemon in that pokedex.							***** pokedex -pokedex- *****",	
			callback: commandPokedex,
		},
		"find":	{
			name: "find",
			description: "Returns a list of locations where the given pokemon can be found in set game version. Limited functionality in more recent versions. If no version is specified, current pokedex version is used. 			***** find -pokemon- -version- *****",
			callback: commandFind,
		},	
		"help": {
			name:	"help",
			description:	"Displays a help message.",
			callback:	commandHelp,
		},
		"inspect": {
			name:	"inspect",
			description: "Adds pokemon to pokedex. Shows details of that pokemon.																				***** inspect -pokemon- *****",
			callback: commandInspect,
		},
		"map":	{
			name:	"map",
			description: "Displays all area location names in current pokedex version.",
			callback: commandMap,
		},
		"exit":	{
			name:	"exit",
			description:	"Exit the Pokedex.",
			callback:	commandExit,
		},
		"explore":	{
			name:	"explore",
			description:	"Explore a location for available pokemon to catch. 																					***** explore -area- *****",
			callback:	commandExplore,
		},
		"pokemon": {
			name: "pokemon",
			description: "Displays names of all pokemon in the user's current pokedex.",
			callback: commandPokemon,
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
			description: "Sets the current pokedex version(red, blue, gold, violet, etc.) 																		***** set-version -version- *****",
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