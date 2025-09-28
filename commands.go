package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, args []string) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args []string) error {
	fmt.Printf("\nWelcome to the Pokedex!\n\n")
	fmt.Printf("Usage:\n")
	for k, v := range getCommand() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(c *config, args []string) error {
	data, err := c.pokeapiClient.GetLocationData(c.Next)
	if err != nil {
		return err
	}

	c.Next = &data.Next
	c.Previous = &data.Previous

	fmt.Println("Next locations:")

	for _, v := range data.Results {
		fmt.Println(" - ", v.Name)
	}

	return nil
}

func commandMapB(c *config, args []string) error {
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

	fmt.Println("Previous locations:")

	for _, v := range data.Results {
		fmt.Println(" - ", v.Name)
	}

	return nil
}

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("location name must be passed as the second argument")
	}
	locationName := args[0]

	data, err := c.pokeapiClient.GetLocationPokemon(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v... \n", data.Name)
	fmt.Println("Found Pokemon:")

	for _, v := range data.PokemonEncounters {
		fmt.Println(" - ", v.Pokemon.Name)
	}

	return nil
}
