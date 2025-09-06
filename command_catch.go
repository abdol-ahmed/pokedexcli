package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/abdol-ahmed/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeAPIClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	if simulateCatch(pokemon) {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.gameState.mu.Lock()
		cfg.gameState.CaughtPokemons[pokemonName] = pokemon
		cfg.gameState.mu.Unlock()
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func simulateCatch(pokemon pokeapi.Pokemon) bool {

	const maxBaseExperience = 608 // Blissey has a high base exp.

	// Calculate a base catch chance.
	// We invert the base experience value relative to the max base experience.
	// This ensures that higher BaseExperience results in a lower catch chance.
	scaledExperience := float64(pokemon.BaseExperience) / maxBaseExperience
	baseCatchChance := 1.0 - scaledExperience

	finalCatchChance := baseCatchChance

	// Clamp the final chance between 0 and 1.
	if finalCatchChance < 0 {
		finalCatchChance = 0
	}
	if finalCatchChance > 1 {
		finalCatchChance = 1
	}

	// Generate a random number to check for a successful catch.
	// In Go 1.20+, the default random number generator is seeded automatically.
	// For older versions, you should seed it with time.Now().UnixNano().
	// For educational purposes, this example shows a seed.
	rand.Seed(time.Now().UnixNano())
	randomRoll := rand.Float64()

	//fmt.Printf("Attempting to catch %s...\n", pokemon.Name)
	//fmt.Printf("Catch chance: %.2f%%\n", finalCatchChance * 100)
	//fmt.Printf("Random roll: %.4f\n", randomRoll)

	return randomRoll < finalCatchChance
}
