package main

import (
	"time"

	"github.com/vystepanenko/pockedexcli/internal/pokedexapi"
)

type config struct {
	pokedexClient        pokedexapi.Client
	nextLocationUrl      *string
	previouseLocationUrl *string
	pokedex              map[string]pokedexapi.Pokemon
}

func main() {
	cfg := config{
		pokedexClient: pokedexapi.NewClient(time.Minute * 5),
		pokedex:       make(map[string]pokedexapi.Pokemon),
	}
	statRepl(&cfg)
}
