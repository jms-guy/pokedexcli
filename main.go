package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", userInput[0])
	}
}

func cleanInput(s string) []string {
	lowerS := strings.ToLower(s)
	results := strings.Fields(lowerS)
	return results
}