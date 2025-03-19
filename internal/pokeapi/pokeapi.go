package pokeapi

import (
	"github.com/jms-guy/pokedexcli/internal/pokecache"
)

//API call functions, list of functions each making different api requests to PokeAPI. Probably could've used a single function for all calls? Seemed easier to make a new function for each different request.//

func (c *Client) GetEncounterData(cache *pokecache.Cache, url string) (EncounterAreas, error) {	//Function to retrieve encounter locations for user input pokemon, through find command - version specific, based on user input game version
	var encounterResults EncounterAreas
	err := requestAndCacheHandling(c, cache, url, &encounterResults)
	if err != nil {
		return encounterResults, err
	}

	return encounterResults, nil
}

func (c *Client) GetVersionGroup(cache *pokecache.Cache, url string) (VersionGroup, error) {
	var versionData Version
	err := requestAndCacheHandling(c, cache, url, &versionData)	//Gets version data for input version
	if err != nil {
		return VersionGroup{}, err
	}

	verGroup := versionData.VersionGroup.Name
	groupURL := "https://pokeapi.co/api/v2/version-group/"+verGroup	//Creates new url for group that version belongs to

	var versionResults VersionGroup
	secErr := requestAndCacheHandling(c, cache, groupURL, &versionResults)	//Gets version group data
	if secErr != nil {
		return versionResults, secErr
	}
	return versionResults, nil
}

func (c *Client) GetPokemonData(cache *pokecache.Cache, url string) (PokemonDetails, error) {	//Function to return pokemon details through catch command
	var pokemonResults PokemonDetails
	err := requestAndCacheHandling(c, cache, url, &pokemonResults)
	if err != nil {
		return pokemonResults, err
	}
	return pokemonResults, nil
}

func (c *Client) GetAreaExplorationData(cache *pokecache.Cache, url string) (LocationAreaDetails, error) {	//Function to return pokemon encounter details through explore command function
	var encounterResults LocationAreaDetails
	err := requestAndCacheHandling(c, cache, url, &encounterResults)
	if err != nil {
		return encounterResults, err
	}
	return encounterResults, nil
}

func (c *Client) GetLocationAreas(cache *pokecache.Cache, pageURL *string) (ConfigData, error) {	//Function to get area locations through map command functions
	url := c.baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	var areaResults ConfigData	
	err := requestAndCacheHandling(c, cache, url, &areaResults)
	if err != nil {
		return areaResults, err
	}
	return areaResults, nil
}
