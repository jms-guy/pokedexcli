package pokeapi

import (
	"net/http"
)

//API response structs, storing different response information depending on the request sent to the PokeAPI//

type Client struct {	//Client struct for http requests
	httpClient	*http.Client
	baseURL		string
}

type Version struct {	//Return struct holding version data from set-version command
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	VersionGroup struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_group"`
}

type RegionData struct {	//Return struct holding region data
	ID        int `json:"id"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"locations"`
	Name  string `json:"name"`
	Pokedexes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokedexes"`
	VersionGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_groups"`
}

type PokedexDetails struct {	//Return struct for pokedex data from pokedex command
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PokemonEntries []struct {
		EntryNumber    int `json:"entry_number"`
		PokemonSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"pokemon_entries"`
	VersionGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_groups"`
}

type VersionGroup struct {	//Return struct holding general version data from set-version command
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID               int    `json:"id"`
	Name             string `json:"name"`
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

type LocationAreaDetails struct {	//Return struct holding location area data for different encounters (explore command)
	Name		string				`json:"name"`
	PokemonEncounters	[]struct {
		Pokemon	struct {
			Name	string	`json:"name"`
			URL		string	`json:"url"`
		}	`json:"pokemon"`
	}	`json:"pokemon_encounters"`
}

type LocationDetails struct {	//Return struct holding location data for explore command
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
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
