package pokeapi

import (
	"net/http"
	"time"
	"github.com/jms-guy/pokedexcli/internal/pokecache"
)

//API call functions, list of functions each making different api requests to PokeAPI. Probably could've used a single function for all calls? Seemed easier to make a new function for each different request.//

func NewClient() *Client {	//Creates new client to handle http requests
	return &Client{
		httpClient: &http.Client{
			Timeout:	time.Minute,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}

func (c *Client) GetEncounterData(cache *pokecache.Cache, url string) (EncounterAreas, error) {	//Function to retrieve encounter locations for user input pokemon, through find command - version specific, based on user input game version
	var encounterResults EncounterAreas
	found, err := getCachedData(cache, url, &encounterResults)	//Checks cache
	if err != nil {
		return encounterResults, err
	}
	if found {
		return encounterResults, nil
	}

	//No cache data found
	httpErr := makeHttpRequest(c, url, &encounterResults)	//Http request
	if httpErr != nil {
		return encounterResults, httpErr
	}

	cacheErr := storeIntoCache(cache, url, &encounterResults)	//Stores data in cache
	if cacheErr != nil {
		return encounterResults, cacheErr
	}

	return encounterResults, nil
}

func (c *Client) GetVersionGroup(cache *pokecache.Cache, url string) (VersionGroup, error) {
	return VersionGroup{}, nil
}

func (c *Client) GetPokemonData(cache *pokecache.Cache, url string) (PokemonDetails, error) {	//Function to return pokemon details through catch command
	var pokemonResults PokemonDetails
	found, err := getCachedData(cache, url, &pokemonResults)	//Checks cache
	if err != nil {
		return pokemonResults, err
	}
	if found {
		return pokemonResults, nil
	}

	//No cache data found
	httpErr := makeHttpRequest(c, url, &pokemonResults)	//Http request
	if httpErr != nil {
		return pokemonResults, httpErr
	}

	cacheErr := storeIntoCache(cache, url, &pokemonResults)	//Stores data in cache
	if cacheErr != nil {
		return pokemonResults, cacheErr
	}

	return pokemonResults, nil
}

func (c *Client) GetAreaExplorationData(cache *pokecache.Cache, url string) (LocationAreaDetails, error) {	//Function to return pokemon encounter details through explore command function
	var encounterResults LocationAreaDetails
	found, err := getCachedData(cache, url, &encounterResults)	//Checks cache
	if err != nil {
		return encounterResults, err
	}
	if found {
		return encounterResults, nil
	}

	//No cache data found
	httpErr := makeHttpRequest(c, url, &encounterResults)	//Http request
	if httpErr != nil {
		return encounterResults, httpErr
	}

	cacheErr := storeIntoCache(cache, url, &encounterResults)	//Stores data in cache
	if cacheErr != nil {
		return encounterResults, cacheErr
	}

	return encounterResults, nil
}

func (c *Client) GetLocationAreas(cache *pokecache.Cache, pageURL *string) (ConfigData, error) {	//Function to get area locations through map command functions
	url := c.baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	var areaResults ConfigData	
	found, err := getCachedData(cache, url, &areaResults)	//Checks cache
	if err != nil {
		return areaResults, err
	}
	if found {
		return areaResults, nil
	}

	//No cache data found
	httpErr := makeHttpRequest(c, url, &areaResults)	//Http request
	if httpErr != nil {
		return areaResults, httpErr
	}

	cacheErr := storeIntoCache(cache, url, &areaResults)	//Stores data in cache
	if cacheErr != nil {
		return areaResults, cacheErr
	}

	return areaResults, nil
}


