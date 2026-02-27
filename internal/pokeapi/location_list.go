package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ListLocations fetches a page of location areas from the PokeAPI 
func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return  RespLocations{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	// Decode JSON response into struct
	var locationsResp RespLocations
	if err := json.NewDecoder(resp.Body).Decode(&locationsResp); err != nil {
		return RespLocations{}, err
	}

	return locationsResp, nil
}