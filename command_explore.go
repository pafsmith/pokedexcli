package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a location name")
	}
	locationName := args[0]

	locationExploreResp, err := cfg.pokeapiClient.ExploreLocation(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring: %s\n", locationExploreResp.Name)
	fmt.Println("Pokemon found:")
	for _, mon := range locationExploreResp.PokemonEncounters {
		fmt.Println(mon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
