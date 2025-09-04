package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]

		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}
