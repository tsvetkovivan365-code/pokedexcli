package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
	"time"
)

// startRepl initializes and runs the Read-Eval-Print loop for the Pokedex CLI
func startRepl() {
	// create a scanner to read user input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize the shared config with API client and pagination state
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}

	// Main REPL loop - runs until uesr exits
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		cmdName := cleanedInput[0]
		cmd, ok := getSupportedCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// cleanInput normalizes user input by converting tp lower case and splitting into words
func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	splitAndTrim := strings.Fields(lower)
	return splitAndTrim
}

// cliCommand represents a command that can be executed in the REPL
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// getSupportedCommands returns a map of all the available commands
func getSupportedCommands() map[string]cliCommand {
	return map[string]cliCommand {
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
			description: "Get the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 locations",
			callback:    commandMapB,
		},
	}
}

// config holds shared state passed to all commands
type config struct {
	pokeapiClient 	 pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}