package main

import (
	"errors"
	"fmt"
)

// Command displays the previous 20 location areas in the Pokemon world
func commandMapB(cfg *config) error {
	// Check if you are already on the first page
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Make an API request to get the previous page of locations
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Update the config with the new pagination URLs from the response
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous
	
	// Loop through and print each location area name
	for i := range resp.Results {
		fmt.Println(resp.Results[i].Name)
	}

	return nil

}