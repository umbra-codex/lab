package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *Config, _ ...string) error {
	if len(config.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
