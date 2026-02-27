package pokeapi

// RespLocations represents the JSON response from the PokeAPI location-area endpoint
type RespLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}