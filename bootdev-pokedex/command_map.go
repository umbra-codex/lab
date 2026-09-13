package main

import (
	"errors"
	"fmt"
)

func commandMapNext(config *Config, _ ...string) error {
	listLocations, err := config.pokeapiClient.ListLocations(config.nextLocationURL)
	if err != nil {
		return err
	}

	for _, location := range listLocations.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationURL = listLocations.Next
	config.previousLocationURL = listLocations.Previous

	return nil
}

func commandMapPrevious(config *Config, _ ...string) error {
	if config.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	listLocations, err := config.pokeapiClient.ListLocations(config.previousLocationURL)
	if err != nil {
		return err
	}

	for _, location := range listLocations.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationURL = listLocations.Next
	config.previousLocationURL = listLocations.Previous

	return nil
}
