package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl() {
	cfg := &config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: nil,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		inputCommand := words[0]

		command, ok := getCommands()[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg); err != nil {
			fmt.Printf("Error executing command %v with error: %v\n", inputCommand, err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Location areas for previous page",
			callback:    commandMapb,
		},
	}
}
