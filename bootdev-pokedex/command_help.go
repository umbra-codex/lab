package main

import "fmt"

func commandHelp(config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range config.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
