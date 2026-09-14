package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/umbra-codex/lab/bootdev-pokedex/internal/pokeapi"
)

type Config struct {
	commands            map[string]Command
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := config.commands[commandName]
		if ok {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
		fmt.Println()
		continue
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)
	return output
}

type Command struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapPrevious,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
