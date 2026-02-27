package pokeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations retrieves a page of location areas, checking the cache before making a network request
func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationsResp RespLocations

	entry, ok := c.cache.Get(url)
	if ok {
		if err := json.Unmarshal(entry, &locationsResp); err != nil {
			return RespLocations{}, err
		}

		return locationsResp, nil
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

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	// Decode JSON response into struct
	if err := json.NewDecoder(tee).Decode(&locationsResp); err != nil {
		return RespLocations{}, err
	}

	// Store bytes in the cache
	c.cache.Add(url, buf.Bytes())

	return locationsResp, nil
}