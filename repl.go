package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mohdanas96/pokedex-go/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(v *config) error
}

func startRepl() {
	sc := bufio.NewScanner(os.Stdin)
	c := &config{}
	for {
		fmt.Printf("\n> Pokedex ")
		sc.Scan()
		inputs := sc.Text()
		words := cleanInput(inputs)
		if len(words) <= 0 {
			fmt.Println("err: no commands/inputs are provided")
			continue
		}
		commandName := words[0]
		command, exists := getCommand()[commandName]
		if !exists {
			fmt.Println("Unkown command")
			continue
		}
		err := command.callback(c)
		if err != nil {
			fmt.Printf("Could not execute the command, err : %v", err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	sliceOfTexts := strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
	return sliceOfTexts
}

func getCommand() map[string]cliCommand {
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
			description: "Displays the next 20 locations in pokeworld",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Display the previous 20 location in pokeworld",
			callback:    commandMapB,
		}}
}
