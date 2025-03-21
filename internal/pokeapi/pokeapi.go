package pokeapi

import (
	"github.com/jms-guy/pokedexcli/internal/pokecache"
)

//API call functions, list of functions each making different api requests to PokeAPI. Probably could've used a single function for all calls? Seemed easier to make a new function for each different request.//

func (c *Client) GetPokedexData(cache *pokecache.Cache, url string) (PokedexDetails, error) {
	var pokedexData PokedexDetails
	err := requestAndCacheHandling(c, cache, url, &pokedexData)
	if err != nil {
		return pokedexData, err
	}

	return pokedexData, nil
}

func (c *Client) GetRegionData(cache *pokecache.Cache, url string) (RegionData, error) {	//Function to get region data for use in map command function
	var regionData RegionData
	err := requestAndCacheHandling(c, cache, url, &regionData)
	if err != nil {
		return regionData, err
	}

	return regionData, nil
}

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

func (c *Client) GetFurtherExplorationData(cache *pokecache.Cache, url string) (LocationDetails, error) {	//Second function for explore command, gets location data instead of locationarea data
	var locationResults LocationDetails
	err := requestAndCacheHandling(c, cache, url, &locationResults)
	if err != nil {
		return locationResults, err
	}
	return locationResults, nil
}

func (c *Client) GetAreaExplorationData(cache *pokecache.Cache, url string) (LocationAreaDetails, error) {	//Function to return pokemon encounter details through explore command function - handles locationarea details
	var encounterResults LocationAreaDetails
	err := requestAndCacheHandling(c, cache, url, &encounterResults)
	if err != nil {
		return encounterResults, err
	}
	return encounterResults, nil
}

