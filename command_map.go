package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokedexClient.GetLocations(cfg.nextLocationUrl)
	if err != nil {
		fmt.Println("Error getting locations:", err)
	}

	fmt.Println("Locations Areas:")
	for _, location := range locations.Results {
		fmt.Printf(" - %s\n", location.Name)
	}

	cfg.nextLocationUrl = locations.Next
	cfg.previouseLocationUrl = locations.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previouseLocationUrl == nil {
		return errors.New("You are on the first page")
	}

	locations, err := cfg.pokedexClient.GetLocations(cfg.previouseLocationUrl)
	if err != nil {
		fmt.Println("Error getting locations:", err)
	}

	fmt.Println("Locations Areas:")
	for _, location := range locations.Results {
		fmt.Printf(" - %s\n", location.Name)
	}

	cfg.nextLocationUrl = locations.Next
	cfg.previouseLocationUrl = locations.Previous

	return nil
}
