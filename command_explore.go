package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]
	locationDetails, err := cfg.pokeAPIClient.GetLocationByName(locationName)

	if err != nil {
		return err
	}

	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
