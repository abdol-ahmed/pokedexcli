package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abdol-ahmed/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeAPIClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{
		pokeAPIClient: pokeapi.NewClient(5 * time.Second),
	}

	for {
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]

		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		executeCommand(command, cfg, args...)
	}
}

func executeCommand(commandName string, cfg *config, args ...string) {
	if cmd, ok := getCommands()[commandName]; ok {
		err := cmd.callback(cfg, args...)
		if err != nil {
			err2 := fmt.Errorf("%w", err)
			fmt.Println(err2.Error())
		}
	} else {
		fmt.Printf("Unknown command\n")
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Navigate forward to display a Pokémon location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Navigate back to display a Pokémon location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
	}
}
