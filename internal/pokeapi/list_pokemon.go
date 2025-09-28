package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationPokemon(locationName string) (RespPokemons, error) {
	fullUrl := baseUrl + "/location-area/" + locationName

	if v, ok := c.cache.Get(fullUrl); ok {
		fmt.Println("RETURNING CACHED DATA")
		pokemonResponse := RespPokemons{}
		err := json.Unmarshal(v, &pokemonResponse)
		if err != nil {
			return RespPokemons{}, err
		}
		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return RespPokemons{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemons{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemons{}, err
	}

	pokemonResponse := RespPokemons{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return RespPokemons{}, err
	}

	c.cache.Add(fullUrl, data)
	return pokemonResponse, nil
}
