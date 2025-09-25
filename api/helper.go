package api

import (
	"net/url"

	"github.com/mohdanas96/pokedex-go/api/internal/pokeapi"
)

const (
	BaseUrl = "https://pokeapi.co/api/v2"
)

type Url struct {
	PathParams  string
	QueryParams url.Values
}

func QueryParamsToString(queryParams url.Values) string {
	queryString := ""
	for k, v := range queryParams {
		queryString += k + "="
		for _, s := range v {
			queryString += s + ","
		}
		queryString += "&"
	}
	return queryString
}

func GetLocationData(url Url) (LocationDetails, error) {
	queryString := QueryParamsToString(url.QueryParams)
	fullUrl := BaseUrl + "/location-area/" + url.PathParams + "?" + queryString

	var locationData LocationDetails

	err := pokeapi.GetData(fullUrl, &locationData)
	if err != nil {
		return LocationDetails{}, err
	}

	return locationData, nil
}
