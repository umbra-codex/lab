package main

import (
	"time"

	"github.com/umbra-codex/lab/bootdev-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &Config{
		Commands:      GetCommands(),
		PokeapiClient: pokeClient,
	}
	startRepl(config)
}
