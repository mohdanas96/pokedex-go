package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetData(url string, v any) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
