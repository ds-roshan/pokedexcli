package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
