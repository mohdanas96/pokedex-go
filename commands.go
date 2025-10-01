package main

import (
	"errors"
	"fmt"
	"math/rand"
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
	data, err := c.pokeapiClient.GetLocations(c.next)
	if err != nil {
		return err
	}

	c.next = &data.Next
	c.previous = &data.Previous

	fmt.Println("Next locations:")

	for _, v := range data.Results {
		fmt.Println(" - ", v.Name)
	}

	return nil
}

func commandMapB(c *config, args []string) error {
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := c.pokeapiClient.GetLocations(c.previous)
	if err != nil {
		return err
	}

	c.next = &data.Next
	c.previous = &data.Previous

	fmt.Println("Previous locations:")

	for _, v := range data.Results {
		fmt.Println(" - ", v.Name)
	}

	return nil
}

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("location name must be passed as the second argument")
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

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("pokemon name must be passed as the second argument")
	}

	pokemonName := args[0]

	_, ok := c.pokedex[pokemonName]
	if ok {
		fmt.Printf("%v is already caught!", pokemonName)
		return nil
	}

	data, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	baseExp := data.BaseExperience

	stats := map[string]int{}
	types := []string{}

	for _, v := range data.Stats {
		stats[v.Stat.Name] = v.BaseStat
	}

	for _, v := range data.Types {
		types = append(types, v.Type.Name)
	}

	randN := rand.Intn(baseExp)
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	if randN > 40 {
		fmt.Printf("%v escaped! Try again\n", pokemonName)
		return nil
	}

	c.pokedex[pokemonName] = Pokemon{name: pokemonName, height: data.Height, weight: data.Weight, types: types, stats: stats}
	fmt.Printf("%v was caught!\n", pokemonName)
	fmt.Printf("You may now inspect it with inspect command\n")

	return nil
}

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("pokemon name must be passed as the second argument")
	}

	pokemonName := args[0]
	v, ok := c.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", v.name)
	fmt.Printf("Height: %v\n", v.height)
	fmt.Printf("Weight: %v\n", v.weight)
	fmt.Printf("Stats: \n")
	for k, v := range v.stats {
		fmt.Printf(" -%v: %v\n", k, v)
	}
	fmt.Printf("Types: \n")
	for _, v := range v.types {
		fmt.Printf(" - %v\n", v)
	}

	return nil
}

func commandPokedex(c *config, args []string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("Your pokedex is empty..! Use 'catch' to catch a Pokemon")
	}
	for k := range c.pokedex {
		fmt.Println(k)
	}
	return nil
}
