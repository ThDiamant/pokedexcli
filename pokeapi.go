package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getLocationAreaEndpointBaseUrl() string {
	return "https://pokeapi.co/api/v2/location-area/"
}

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreaData(config *Config) error {
	locData, err := getLocationAreaFromApi(config)
	if err != nil {
		return err
	}

	extractLocationAreaData(config, locData)

	return nil
}

func getLocationAreaFromApi(config *Config) (locationAreaResponse, error) {
	res, err := http.Get(config.next)
	if err != nil {
		return locationAreaResponse{}, fmt.Errorf("Error while creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaResponse{}, fmt.Errorf("Error reading response data %w", err)
	}

	var locAreaResp locationAreaResponse
	if err := json.Unmarshal(data, &locAreaResp); err != nil {
		return locationAreaResponse{}, fmt.Errorf("Error unmarshalling response data: %w", err)
	}

	return locAreaResp, nil
}

func extractLocationAreaData(config *Config, locAreaResp locationAreaResponse) {
	config.next = locAreaResp.Next
	config.prev = locAreaResp.Previous

	var areaNames []string
	for _, result := range locAreaResp.Results {
		areaNames = append(areaNames, result.Name)
	}

	config.currentLocs = areaNames
}
