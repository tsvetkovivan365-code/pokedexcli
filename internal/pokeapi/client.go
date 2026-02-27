package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

// Client wraps an HTTP client for making requests to the PokeAPI and caching the results.
type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

// NewClient creates a new Client with a request timeout and a cache reaping interval.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
