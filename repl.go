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
		fmt.Print("> Pokedex ")
		sc.Scan()
		inputs := sc.Text()
		inputsSlice := cleanInput(inputs)
		fmt.Printf("Your command was: %s\n", inputsSlice[0])
	}
}

func cleanInput(text string) []string {
	sliceOfTexts := strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
	return sliceOfTexts
}
