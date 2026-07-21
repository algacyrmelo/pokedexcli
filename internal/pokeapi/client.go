package pokeapi

import (
	"net/http"
	"time"

	"github.com/algacyrmelo/pokedexcli/internal/pokecache"
)

// Client makes requests to the PokeAPI and caches responses.
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient creates a new Client with the given timeout and cache interval.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
