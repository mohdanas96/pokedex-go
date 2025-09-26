package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(pageUrl *string) (RespLocations, error) {
	fullUrl := baseUrl + "/location-area"
	if pageUrl != nil {
		fullUrl = *pageUrl
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

	return locationsResp, nil
}
