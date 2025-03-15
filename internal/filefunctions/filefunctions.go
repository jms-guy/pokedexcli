package filefunctions

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/jms-guy/pokedexcli/internal/pokeapi"
)

//Currently only one save file is supported, if you have previous saved data and do not load it upon restarting the program, the save data will be overwritten on another save.

type PokedexProvider interface {
	GetUserPokedex() *map[string]pokeapi.PokemonDetails
}

func LoadPokedex(provider PokedexProvider) error {	//File function to load pokedex save data from file into shared state struct pokedex map
	if _, err := os.Stat("pokedex.json"); err != nil {		//Checks for existence of save data
		if os.IsNotExist(err) {
			return fmt.Errorf("no save data available to load")
		} else {
			return fmt.Errorf("error loading save data: %w", err)
		}
	}
	
	body, err := os.ReadFile("pokedex.json")	//Reads save data file
	if err != nil {
		return fmt.Errorf("error reading save file: %w", err)
	}

	if err := json.Unmarshal(body, provider.GetUserPokedex()); err != nil {	//Unmarshals json data from save file into struct
		return fmt.Errorf("error unmarshaling save data into pokedex: %w", err)
	}
	return nil
}

func SavePokedex(provider PokedexProvider) error {	//Save file function for saving data from the current pokedex into a file, to be read on a new program cycle
	file, err := os.Create("pokedex.json")	//Creates file to save to
	if err != nil {
		return fmt.Errorf("error opening file pokedex.json: %w", err)
	}
	defer file.Close()
	
	pokeData, err := json.Marshal(*provider.GetUserPokedex())	//Marshals pokedex data into json
	if err != nil {
		return fmt.Errorf("error marshaling pokedex data: %w", err)
	}

	n, err := file.Write(pokeData)	//Writes save file
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	if n != len(pokeData) {
		return fmt.Errorf("error writing to file: wrote %d bytes, expected %d", n, len(pokeData))
	}
	return nil
}