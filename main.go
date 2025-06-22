package main

import (
	"time"

	"github.com/ds-roshan/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.RespPokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
