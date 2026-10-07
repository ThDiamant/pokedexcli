package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const locationBaseURL = "https://pokeapi.co/api/v2/location-area/"
const pokemonBaseURL = "https://pokeapi.co/api/v2/pokemon/"

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type specificAreaResponse struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func getSpecificLocationPokemonData(config *Config) ([]string, error) {
	url := locationBaseURL + config.commandCallbackarg

	pokLocData, err := getLocationPokemonFromApi(config, url)
	if err != nil {
		return []string{}, err
	}

	areaNames, err := extractPokemonAreaData(pokLocData)

	return areaNames, nil
}

func getLocationPokemonFromApi(config *Config, url string) (specificAreaResponse, error) {
	data, ok := config.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return specificAreaResponse{}, fmt.Errorf("Error while creating request: %w\n", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return specificAreaResponse{}, fmt.Errorf("Error reading response data %w\n", err)
		}
	}

	config.cache.Add(url, data)

	var specificLocResp specificAreaResponse
	if err := json.Unmarshal(data, &specificLocResp); err != nil {
		return specificAreaResponse{}, fmt.Errorf("Error unmarshalling response data: %w\n", err)
	}

	return specificLocResp, nil
}

func extractPokemonAreaData(specificLocResp specificAreaResponse) ([]string, error) {
	pokemonNames := []string{}

	for _, pokemon := range specificLocResp.PokemonEncounters {
		pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
	}

	return pokemonNames, nil
}

func getPokemonData(config *Config) (Pokemon, error) {
	url := pokemonBaseURL + config.commandCallbackarg

	pokeApiData, err := getPokemonDataFromApi(config, url)
	if err != nil {
		return Pokemon{}, err
	}

	pokeData, err := extractPokemonData(pokeApiData)
	if err != nil {
		return Pokemon{}, err
	}

	return pokeData, nil
}

func getPokemonDataFromApi(config *Config, url string) (pokemonResponse, error) {
	data, ok := config.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return pokemonResponse{}, fmt.Errorf("Error while creating request: %w\n", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return pokemonResponse{}, fmt.Errorf("Error reading response data %w\n", err)
		}
	}

	config.cache.Add(url, data)

	var pokemonResp pokemonResponse
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return pokemonResponse{}, fmt.Errorf("Error unmarshalling response data: %w\n", err)
	}

	return pokemonResp, nil
}

func extractPokemonData(pokeResp pokemonResponse) (Pokemon, error) {

	stats := make(map[string]int)
	for _, stat := range pokeResp.Stats {
		stats[stat.Stat.Name] = stat.BaseStat
	}

	types := []string{}
	for _, pokeType := range pokeResp.Types {
		types = append(types, pokeType.Type.Name)
	}

	return Pokemon{
		Name:           pokeResp.Name,
		Height:         pokeResp.Height,
		Weight:         pokeResp.Weight,
		BaseExperience: pokeResp.BaseExperience,
		Stats: PokemonStats{
			Hp:             stats["hp"],
			Attack:         stats["attack"],
			Defence:        stats["defence"],
			SpecialAttack:  stats["special-attack"],
			SpecialDefence: stats["special-defence"],
			Speed:          stats["speed"],
		},
		Types: types,
	}, nil
}
