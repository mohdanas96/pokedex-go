package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
	fullUrl := baseUrl + "/pokemon/" + pokemonName
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResponse := RespPokemon{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemonResponse, nil
}
