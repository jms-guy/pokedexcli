package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/jms-guy/internal/pokecache"
)

type Client struct {
	httpClient	*http.Client
	baseURL		string
}

type ConfigData struct {
	Count		int		`json:"count"`
	Next		*string	`json:"next"`
	Previous	*string	`json:"previous"`
	Results		[]struct {
		Name	string `json:"name"`
		URL		string `json:"url"`
	}	`json:"results"`
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout:	time.Minute,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}

func (c *Client) GetLocationAreas(pageURL *string) (ConfigData, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ConfigData{}, fmt.Errorf("error making request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ConfigData{}, fmt.Errorf("error requesting data: %w", err)
	}
	defer res.Body.Close()

	var areaResults ConfigData
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&areaResults); err != nil {
		return ConfigData{}, fmt.Errorf("error decoding json data: %w", err)
	}
	return areaResults, nil
}
