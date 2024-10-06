package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) < 1 {
		return errors.New("You haven't caught any pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
