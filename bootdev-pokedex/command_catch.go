package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	threshold := 45
	result := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if result > threshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("you may now inspect it with the inspect command")
	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
