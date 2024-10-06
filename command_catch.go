package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Pokemon name isn't provided")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokedexClient.GetPokemon(pokemonName)
	if err != nil {
		fmt.Printf("Error getting pokemon %s: %s", pokemonName, err)
	}

	const catchingTrashHold = 50
	difficulty := rand.Intn(pokemon.BaseExperience)

	if difficulty > catchingTrashHold {
		return errors.New("Pokemon has escaped")
	}

	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Printf("You have caught %s\n", pokemon.Name)
	return nil
}
