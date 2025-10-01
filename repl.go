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
	pokedex       map[string]Pokemon
	next          *string
	previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(v *config, args []string) error
}

type Pokemon struct {
	name   string
	height int
	weight int
	types  []string
	stats  map[string]int
}

func startRepl(c *config) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		sc.Scan()
		inputs := sc.Text()
		words := cleanInput(inputs)
		if len(words) <= 0 {
			fmt.Println("err: no commands/inputs are provided")
			continue
		}
		commandName := words[0]
		commandArgs := make([]string, 0)
		if len(words) > 1 {
			commandArgs = words[1:]
		}
		command, exists := getCommand()[commandName]
		if !exists {
			fmt.Println("Unkown command")
			continue
		}
		err := command.callback(c, commandArgs)
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
			name:        "map forward",
			description: "Displays the next 20 locations in pokeworld",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map backward",
			description: "Display the previous 20 location in pokeworld",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays all the Pokemons in a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a Pokemon and stores it",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the info about the Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the pokemon the user have caught",
			callback:    commandPokedex,
		},
	}
}
