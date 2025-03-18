package pokeapi

import (
	"net/http"
)

//API response structs, storing different response information depending on the request sent to the PokeAPI//

type APIResponse interface {	//Interface for return structs
	GetName() string
}

type Client struct {	//Client struct for http requests
	httpClient	*http.Client
	baseURL		string
}

type VersionGroup struct {	//Return struct holding general version data from set-version command
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID               int    `json:"id"`
	MoveLearnMethods []any  `json:"move_learn_methods"`
	Name             string `json:"name"`
	Order            int    `json:"order"`
	Pokedexes        []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokedexes"`
	Regions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"regions"`
	Versions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"versions"`
}

// This structure's command function are currently unused
type ConfigData struct {	//Return struct holding json data returned from http requests for area locations (map/mapb commands) 
	Count		int		`json:"count"`
	Next		*string	`json:"next"`
	Previous	*string	`json:"previous"`
	Results		[]struct {
		Name	string `json:"name"`
		URL		string `json:"url"`
	}	`json:"results"`
}

type LocationAreaDetails struct {	//Return struct holding location area data for different encounters (explore command)
	Name		string				`json:"name"`
	PokemonEncounters	[]struct {
		Pokemon	struct {
			Name	string	`json:"name"`
			URL		string	`json:"url"`
		}	`json:"pokemon"`
	}	`json:"pokemon_encounters"`
}

type PokemonDetails struct {	//Return struct holding detailed pokemon data (catch command)
	BaseExperience int `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Name          string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type EncounterAreas []struct {	//Return struct holding pokemon encounter data (find _____ command)
	LocationArea struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location_area"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int   `json:"chance"`
			ConditionValues []any `json:"condition_values"`
			MaxLevel        int   `json:"max_level"`
			Method          struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"method"`
			MinLevel int `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int `json:"max_chance"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
}

//Struct methods to allow each response structure to satisfy the APIResponse interface, allowing all structs to be passed into the command functions properly//

func (c *ConfigData) GetName() string {		
	return ""
}

func (l *LocationAreaDetails) GetName() string {	
	return ""
}

func (p *PokemonDetails) GetName() string {	
	return ""
}

func (e *EncounterAreas) GetName() string {
	return ""
}