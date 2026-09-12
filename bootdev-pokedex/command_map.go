package main

import (
	"errors"
	"fmt"
)

func commandMapNext(config *Config) error {
	response, err := config.pokeapiClient.ListLocations(config.nextLocationURL)
	if err != nil {
		return err
	}

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationURL = response.Next
	config.previousLocationURL = response.Previous

	return nil
}

func commandMapPrevious(config *Config) error {
	if config.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	response, err := config.pokeapiClient.ListLocations(config.previousLocationURL)
	if err != nil {
		return err
	}

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationURL = response.Next
	config.previousLocationURL = response.Previous

	return nil
}
