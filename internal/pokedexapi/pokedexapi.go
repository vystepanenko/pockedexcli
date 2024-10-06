package pokedexapi

import (
	"net/http"
	"time"

	"github.com/vystepanenko/pockedexcli/internal/pokedexcache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokedexcache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokedexcache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
