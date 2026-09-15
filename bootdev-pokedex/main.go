package main

import (
	"time"

	"github.com/umbra-codex/lab/bootdev-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	config := &Config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(config)
}
