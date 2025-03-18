package versionfunctions

import (
	"github.com/jms-guy/pokedexcli/internal/pokeapi"
)

//These are functions related to the command functions that return different results based on the version input by the user. Such as sorting response data relevant to the game version.//

type VersionDetails []struct {	//Return struct slice for storing refined data from find command function
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
}

func VersionEncounters(areaData pokeapi.EncounterAreas, gameVersion string) map[string]VersionDetails {	//Sorts encounter data from the find command function, returns data only relevant to the game version.
	
	relevantData := make(map[string]VersionDetails)	//Map, links the relevant encounter data to the area name
	
	for _, encounterArea := range areaData {
		for _, verDetails := range encounterArea.VersionDetails {
			if verDetails.Version.Name == gameVersion {
				relevantData[encounterArea.LocationArea.Name] = append(relevantData[encounterArea.LocationArea.Name], verDetails)
			}
		}
	}
	return relevantData
}