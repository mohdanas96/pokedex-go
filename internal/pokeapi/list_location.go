package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(pageUrl *string) (RespLocations, error) {
	fullUrl := baseUrl + "/location-area"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	if val, ok := c.cache.Get(fullUrl); ok {
		fmt.Println("RETURNING CACHED DATA")
		locationsResp := RespLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return RespLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(fullUrl, data)
	return locationsResp, nil
}
