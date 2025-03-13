package filefunctions

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/jms-guy/pokedexcli/internal/pokeapi"
)

type PokedexProvider interface {
	GetUserPokedex() map[string]pokeapi.PokemonDetails
}

func SavePokedex(provider PokedexProvider) error {
	file, err := os.Create("pokedex.json")
	if err != nil {
		return fmt.Errorf("error opening file pokedex.json: %w", err)
	}
	defer file.Close()
	
	pokeData, err := json.Marshal(provider.GetUserPokedex())
	if err != nil {
		return fmt.Errorf("error marshaling pokedex data: %w", err)
	}

	n, err := file.Write(pokeData)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	if n != len(pokeData) {
		return fmt.Errorf("error writing to file: wrote %d bytes, expected %d", n, len(pokeData))
	}
	return nil
}