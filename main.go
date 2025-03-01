package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/jms-guy/pokedexproject/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient()		//Creating http client
	configData := pokeapi.ConfigData{}	//Creating base ConfigData struct for json data
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
		} else {
			err := command.callback(client, &configData)	//Executes command from registry
			if err != nil {
				fmt.Printf("Returned error: %v", err)
				continue
			}
		}
		
	}
}

