package main

import (
	"strings"
	"fmt"
    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

//Helper functions for command functions in commands.go

func capitalizeString(s string) string {    //Simple function to just title-case a string input 'Like This' - all lowercase is kinda ugly.
    return cases.Title(language.Und, cases.NoLower).String(s)
}

func cleanInput(s string) []string {	//Cleans user input string for command arguments use
	lowerS := strings.ToLower(s)
	results := strings.Fields(lowerS)
	return results
}

func ParseVersion(input string) (Version, error) {	//Checks version input used in set-version command against enum struct in versions.go
	for v, name := range versionName {
		if name == input {
			return v, nil
		}
	}
	return 0, fmt.Errorf("unknown version: %s", input)
}

func extractConditionNames(conditions []any) string {	
    var conditionNames []string 

    for _, condition := range conditions {
        // Assert that each item is a map[string]interface{}
        if conditionMap, ok := condition.(map[string]any); ok {
            
            if name, exists := conditionMap["name"].(string); exists {
                conditionNames = append(conditionNames, capitalizeString(name)) // Add name to the list
            }
        }
    }

    // Join the condition names into a single, user-friendly string
    return fmt.Sprintf("[%s]", joinStrings(conditionNames, ", "))
}

func joinStrings(elements []string, separator string) string {
    joined := ""
    for i, element := range elements {
        if i > 0 {
            joined += separator
        }
        joined += element
    }
    return joined
}