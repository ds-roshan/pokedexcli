package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) < 2 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	respPokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	attemptCatch := attemptCatch(respPokemon.BaseExperience)

	if attemptCatch {
		fmt.Printf("%s was caught!\n", respPokemon.Name)
		cfg.caughtPokemon[respPokemon.Name] = respPokemon
	} else {
		fmt.Printf("%s escaped!\n", respPokemon.Name)
	}
	return nil
}

func attemptCatch(baseExperience int) bool {

	chance := max(100-baseExperience, 10)

	roll := rand.Intn(100) + 1
	fmt.Printf("Catch chance %v percentage, Roll %v\n", chance, roll)

	return roll <= chance
}
