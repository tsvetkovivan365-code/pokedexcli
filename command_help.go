package main

import (
	"fmt"
)

// commandHelp prints the supported commands name and description
func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	
	for _, cmd := range getSupportedCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}