package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/jms-guy/pokedexcli/internal/pokecache"
)

func NewClient() *Client {	//Creates new client to handle http requests
	return &Client{
		httpClient: &http.Client{
			Timeout:	time.Minute,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}

func requestAndCacheHandling(c *Client, cache *pokecache.Cache, url string, dataStruct any) error {	//Function that combines the other helper functions into one, that handles getting cached data, making http requests, and storing data to cache
	found, err := getCachedData(cache, url, dataStruct)	//Checks cache
	if err != nil {
		return err
	}
	if found {
		return nil
	}

	//No cache data found
	httpErr := makeHttpRequest(c, url, dataStruct)	//Http request
	if httpErr != nil {
		return httpErr
	}

	cacheErr := storeIntoCache(cache, url, dataStruct)	//Stores data in cache
	if cacheErr != nil {
		return cacheErr
	}

	return nil
}

func makeHttpRequest(c *Client, url string, dataStruct any) (error) {	//Makes http request and handles resposne
	req, err := http.NewRequest("GET", url, nil)	//Creates http request
	if err != nil {
		return  fmt.Errorf("error making request: %w", err)
	}

	res, err := c.httpClient.Do(req)	//Sends http request
	if err != nil {
		return  fmt.Errorf("error requesting data: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)	//Decodes response
	if err = decoder.Decode(dataStruct); err != nil {
		return  fmt.Errorf("error decoding json data: %w", err)
	}
	return nil
}

func getCachedData(cache *pokecache.Cache, url string, dataStruct any) (bool, error) {	//Retrieves cache data
	cachedData := checkCache(cache, url)
	if cachedData == nil {
        return false, nil // Cache miss
    }
    
    err := json.Unmarshal(cachedData, dataStruct)
    if err != nil {
        return false, fmt.Errorf("error unmarshaling json data: %w", err)
    }
    return true, nil // Cache hit
}

func storeIntoCache(cache *pokecache.Cache, url string, datastruct any) error {	//Marshals data and stores it in cache
	dataToCache, err := json.Marshal(datastruct)
	if err != nil {
		return fmt.Errorf("error marshaling data for cache: %w", err)
	}
	cache.Add(url, dataToCache)
	return nil
}

func checkCache(c *pokecache.Cache, pageURL string) []byte {	//Checks cache for existence of data
	val, ok := c.Get(pageURL)
	if !ok {
		return nil
	}
	return val
}
