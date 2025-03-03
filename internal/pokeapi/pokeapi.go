package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/jms-guy/pokedexcli/internal/pokecache"
)

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

func NewClient() *Client {	//Creates new client to handle http requests
	return &Client{
		httpClient: &http.Client{
			Timeout:	time.Minute,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}

func (c *Client) GetPokemonData(cache *pokecache.Cache, url string) (PokemonDetails, error) {	//Function to return pokemon details through catch command
	cachedData := checkCache(cache, url)	//Checks cache for data
	if cachedData != nil {
		var pokemonResults PokemonDetails
		err := json.Unmarshal(cachedData, &pokemonResults)
		if err != nil {
			return PokemonDetails{}, fmt.Errorf("error unmarshaling json data from cache: %w", err)
		}
		return pokemonResults, nil
	}

	req, err := http.NewRequest("GET", url, nil)	//Makes http request
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error making request: %w", err)
	}

	res, err := c.httpClient.Do(req)	//Sends request
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error requesting data: %w", err)
	}
	defer res.Body.Close()

	var pokemonResults PokemonDetails
	decoder := json.NewDecoder(res.Body)	//Decodes response
	if err := decoder.Decode(&pokemonResults); err != nil {
		return PokemonDetails{}, fmt.Errorf("error decoding json data: %w", err)
	}

	dataToCache, err := json.Marshal(pokemonResults)	//Marshals data for cache
	if err != nil {
		return pokemonResults, fmt.Errorf("error marshaling data for cache: %w", err)
	}
	cache.Add(url, dataToCache)
	return pokemonResults, nil
}

func (c *Client) GetAreaExplorationData(cache *pokecache.Cache, url string) (LocationAreaDetails, error) {	//Function to return pokemon encounter details through explore command function
	cachedData := checkCache(cache, url)	//Checks cache for data
	if cachedData != nil {
		var encounterResults LocationAreaDetails
		err := json.Unmarshal(cachedData, &encounterResults)
		if err != nil {
			return LocationAreaDetails{}, fmt.Errorf("error unmarshaling json data from cache: %w", err)
		}
		return encounterResults, nil
	}

	req, err := http.NewRequest("GET", url, nil)	//Makes hhtp request
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error making request: %w", err)
	}

	res, err := c.httpClient.Do(req)	//Sends http request
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error requesting data: %w", err)
	}
	defer res.Body.Close()

	var encounterResults LocationAreaDetails
	decoder := json.NewDecoder(res.Body)	//Decodes response data
	if err := decoder.Decode(&encounterResults); err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error decoding json data: %w", err)
	}
	dataToCache, err := json.Marshal(encounterResults)	//Marshals data into cache
	if err != nil {
		return encounterResults, fmt.Errorf("error marshaling data for cache: %w", err)
	}
	cache.Add(url, dataToCache)
	return encounterResults, nil
}

func (c *Client) GetLocationAreas(cache *pokecache.Cache, pageURL *string) (ConfigData, error) {	//Function to get area locations through map command functions
	url := c.baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	var areaResults ConfigData
	cachedData := checkCache(cache, url)	//Checks cache for data before requesting
	if cachedData != nil {
		err := json.Unmarshal(cachedData, &areaResults)
		if err != nil {
			return ConfigData{}, fmt.Errorf("error unmarshaling json data: %w", err)
		}
		return areaResults, nil
	}

	req, err := http.NewRequest("GET", url, nil)	//Creates http request
	if err != nil {
		return ConfigData{}, fmt.Errorf("error making request: %w", err)
	}

	res, err := c.httpClient.Do(req)	//Sends http request
	if err != nil {
		return ConfigData{}, fmt.Errorf("error requesting data: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)	//Decodes response
	if err = decoder.Decode(&areaResults); err != nil {
		return ConfigData{}, fmt.Errorf("error decoding json data: %w", err)
	}

	dataToCache, err := json.Marshal(areaResults)	//Marshals response data into the cache
	if err != nil {
		return areaResults, fmt.Errorf("error marshaling data for cache: %w", err)
	}
	cache.Add(url, dataToCache)
	return areaResults, nil
}

func checkCache(c *pokecache.Cache, pageURL string) []byte {	//Checks cache for existence of data
	val, ok := c.Get(pageURL)
	if !ok {
		return nil
	}
	return val
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