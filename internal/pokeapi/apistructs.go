package pokeapi

import (
	"net/http"
)

//API response structs//

type APIResponse interface {	//Interface for return structs
	GetName() string
}

type Client struct {	//Client struct for http requests
	httpClient	*http.Client
	baseURL		string
}

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
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
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

func (c *ConfigData) GetName() string {		//Struct method to satisfy interface definition
	return ""
}

func (l *LocationAreaDetails) GetName() string {	//Struct method to satisfy interface definition
	return ""
}

func (p *PokemonDetails) GetName() string {	//Struct method to sastisfy interface definition
	return ""
}