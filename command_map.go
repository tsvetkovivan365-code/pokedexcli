package main

import "fmt"

func commandMap(cfg *config) error {
	// Make an API request to get the next page of locations
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
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