package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	if pokemon, ok := cfg.gameState.CaughtPokemons[pokemonName]; ok {
		cfg.gameState.mu.Lock()
		defer cfg.gameState.mu.Unlock()
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats: \n")
		for _, stat := range pokemon.Stats {
			switch stat.Stat.Name {
			case "hp":
				fmt.Printf("\t-hp: %d\n", stat.BaseStat)
			case "attack":
				fmt.Printf("\t-attack: %d\n", stat.BaseStat)
			case "defense":
				fmt.Printf("\t-defense: %d\n", stat.BaseStat)
			case "special-attack":
				fmt.Printf("\t-special-attack: %d\n", stat.BaseStat)
			case "special-defense":
				fmt.Printf("\t-special-defense: %d\n", stat.BaseStat)
			case "speed":
				fmt.Printf("\t-speed: %d\n", stat.BaseStat)
			}

		}
		fmt.Printf("Stats: \n")
		for _, type1 := range pokemon.Types {
			fmt.Printf("\t- %s\n", type1.Type.Name)
		}
	} else {
		return errors.New("you have not caught that pokemon")
	}

	return nil
}
