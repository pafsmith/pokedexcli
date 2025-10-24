package main

import (
	"errors"
	"fmt"
)



func commandExplore(cfg *config, args []string) error {
 if args[0] == "" {
 	return errors.New("invalid call")
 }
locationExploreResp, err := cfg.pokeapiClient.ExploreLocation(args[0])
if err != nil {
	return err
}
	for _, mon := range locationExploreResp.PokemonEncounters{
		fmt.Println(mon.Pokemon.Name)
	}
	return nil


}
