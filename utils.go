package main

import (
	"net/url"
)

func ExtractQueryParamsFromUrl(urlText string) (url.Values, error) {
	u, err := url.Parse(urlText)
	if err != nil {
		return nil, err
	}

	queryParams, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	return queryParams, nil
}
