package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Printf("Your Pokedex:\n")
	for key, _ := range cfg.gameState.CaughtPokemons {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
