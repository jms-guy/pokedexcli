package api_commands

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var client = http.Client{}

type config struct {
	next		string
	previous 	string
}

type locationAreaResponse struct {
	id		int
	name	string

}

func getLocations(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
}