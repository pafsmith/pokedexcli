package pokeapi

import (
	"net/http"
	"time"

	"github.com/pafsmith/pokedexcli/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInteval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInteval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
