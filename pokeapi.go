package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2/location-area/"

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreaData(config *Config) ([]string, error) {

	url := config.next
	if !config.goNext {
		url = config.prev
	}

	locData, err := getLocationAreaFromApi(config, url)
	if err != nil {
		return []string{}, err
	}

	areaNames, err := extractLocationAreaData(config, locData)
	
	return areaNames, nil
}

func getLocationAreaFromApi(config *Config, url string) (locationAreaResponse, error) {

	data, ok := config.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locationAreaResponse{}, fmt.Errorf("Error while creating request: %w\n", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return locationAreaResponse{}, fmt.Errorf("Error reading response data %w\n", err)
		}
	}

	config.cache.Add(url, data)

	var locAreaResp locationAreaResponse
	if err := json.Unmarshal(data, &locAreaResp); err != nil {
		return locationAreaResponse{}, fmt.Errorf("Error unmarshalling response data: %w\n", err)
	}

	return locAreaResp, nil
}

func extractLocationAreaData(config *Config, locAreaResp locationAreaResponse) ([]string, error) {
	config.next = locAreaResp.Next
	config.prev = locAreaResp.Previous

	var areaNames []string
	for _, result := range locAreaResp.Results {
		areaNames = append(areaNames, result.Name)
		fmt.Println(result.Name)
	}

	return areaNames, nil
}
