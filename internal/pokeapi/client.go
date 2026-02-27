package pokeapi

import (
	"net/http"
	"time"
)

// Client wraps an HTTP client for making requests to the PokeAPI
type Client struct {
	httpClient http.Client
}

// NewClient creates a new Client with the specified request timeout 
func NewClient(timeout time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
