package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\n\n")
	fmt.Printf("Usage:\n")
	for k, v := range commandRegistry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
