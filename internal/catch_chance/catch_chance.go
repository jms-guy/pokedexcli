package catchchance

import (
	"math/rand"
	"time"
)

// Function to generate a catch rate for pokemon in main function, currently not being used, but can be re-implemented if desired.

func GetCatchBool(baseExp int) bool {	//Get catch chance for pokemon, very simple implmentation
	rand.NewSource(time.Now().UTC().UnixNano())
	randomInt := rand.Intn(100)	//Random number between 0-100
	if baseExp <= 75 {
		if randomInt >= 10 {
			return true
		} else {
			return false
		}
	} else if baseExp <= 150 {
		if randomInt >= 30 {
			return true
		} else {
			return false
		}
	} else if baseExp <= 250 {
		if randomInt >= 50 {
			return true
		} else {
			return false
		}
	} else {
		if randomInt >= 75 {
			return true
		} else {
			return false
		}
	}
}