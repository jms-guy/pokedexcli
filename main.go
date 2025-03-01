package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/jms-guy/pokedexproject/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient()		//Creating http client
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
			err := command.callback()	//Executes command from registry
			if err != nil {
				fmt.Printf("Returned error: %v", err)
				continue
			}
		}
		
	}
}

func cleanInput(s string) []string {	//Cleans user input string and returns first word in a lowercase state
	lowerS := strings.ToLower(s)
	results := strings.Fields(lowerS)
	return results
}
