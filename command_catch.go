package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}
	pokemonName := args[0]

	pokemonResponse, err := cfg.pokeapiClient.FindPokemon(pokemonName)
	if err != nil {
		return errors.New("could not find pokemon")
	}

	randInt := rand.Intn(pokemonResponse.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResponse.Name)

	if randInt > 33 {
		fmt.Printf("%s escaped!\n", pokemonResponse.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResponse.Name)

	cfg.pokedex[pokemonResponse.Name] = pokemonResponse
	return nil
}
