package main

import "fmt"

func commandHelp(config *Config, _ ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	for _, command := range config.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
