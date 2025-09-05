package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokeAPIClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {

	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeAPIClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
