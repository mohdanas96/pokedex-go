package main

import (
	"fmt"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\n\n")
	fmt.Printf("Usage:\n")
	for k, v := range getCommand() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(c *config) error {
	data, err := c.pokeapiClient.GetLocationData(c.Next)
	if err != nil {
		return err
	}

	c.Next = &data.Next
	c.Previous = &data.Previous

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := c.pokeapiClient.GetLocationData(c.Previous)
	if err != nil {
		return err
	}

	c.Next = &data.Next
	c.Previous = &data.Previous

	for _, v := range data.Results {
		fmt.Println(v.Name)
	}

	return nil
}
