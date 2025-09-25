package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Next     string
	Previous string
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
		command := words[0]
		v, ok := commandRegistry[command]
		if !ok {
			fmt.Println("Unkown command")
			continue
		}
		err := v.callback(c)
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
