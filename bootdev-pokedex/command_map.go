package main

import (
	"errors"
	"fmt"
)

func commandMapNext(config *Config) error {
	locationResp, err := config.PokeapiClient.ListLocations(config.NextLocationURL)
	if err != nil {
		return err
	}

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	config.NextLocationURL = locationResp.Next
	config.PreviousLocationURL = locationResp.Previous

	return nil
}

func commandMapPrevious(config *Config) error {
	if config.PreviousLocationURL == nil {
		return errors.New("you're on the first page\n")
	}

	locationResp, err := config.PokeapiClient.ListLocations(config.PreviousLocationURL)
	if err != nil {
		return err
	}

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	config.NextLocationURL = locationResp.Next
	config.PreviousLocationURL = locationResp.Previous

	return nil
}
