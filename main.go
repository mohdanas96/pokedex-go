package main

import (
	"time"

	"github.com/mohdanas96/pokedex-go/internal/pokeapi"
)

func main() {
	startRepl(&config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		pokedex:       make(map[string]Pokemon),
	})
}
