package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Expected exectly one argument")
	}

	areaName := args[0]
	location, err := cfg.pokedexClient.GetArea(areaName)
	if err != nil {
		fmt.Printf("Error getting location %s: %s", areaName, err)
	}

	fmt.Println("Pokemon list:")
	for _, location := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", location.Pokemon.Name)
	}

	return nil
}
