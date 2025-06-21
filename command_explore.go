package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	fmt.Println("total items: ", len(args))

	if len(args) < 2 {
		return errors.New("you must provide a location name")
	}

	name := args[1]
	locationAreaRes, err := cfg.pokeapiClient.ListLocationArea(name)
	if err != nil {
		return err
	}

	encounters := locationAreaRes.PokemonEncounters
	for _, encounter := range encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
