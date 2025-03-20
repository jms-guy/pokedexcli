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
	MainGeneration struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"main_generation"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pokedexes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokedexes"`
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

// Command utilizing this struct has been discontinued, might be annoying to remove it
type ConfigData struct {	//Return struct holding json data returned from http requests for area locations (mapb command) 
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

type LocationDetails struct {	//Return struct holding location data for explore command
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
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
