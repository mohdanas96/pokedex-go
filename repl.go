package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	sc := bufio.NewScanner(os.Stdin)
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
		err := v.callback()
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
