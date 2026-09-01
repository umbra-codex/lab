package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/umbra-codex/lab/bootdev-pokedex/internal/pokeapi"
)

type Config struct {
	Commands            map[string]Command
	PokeapiClient       pokeapi.Client
	NextLocationURL     *string
	PreviousLocationURL *string
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

		command, ok := config.Commands[commandName]
		if ok {
			err := command.Callback(config)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown Command\n")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)
	return output
}

type Command struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMapNext,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapPrevious,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
