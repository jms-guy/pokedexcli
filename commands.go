package main

import (
	"fmt"
	"os"
	"pokedexproject/internal/api_commands.go"
)

type cliCommand struct {	//Struct for user input commands in the cli
	name		string
	description	string
	callback	func() error
}

var commandRegistry map[string]cliCommand	//Declaration of Command Registry

func commandHelp(c *config) error {	//Help command function
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *config) error {	//Map command function

}

func commandExit(c *config) error {	//Exit command function
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
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
		"exit":	{
			name:	"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
	}
}