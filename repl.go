package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var returnSlice []string
	for _, s := range strings.Split(strings.TrimSpace(text), " ") {
		if s == "" {
			continue
		}
		returnSlice = append(returnSlice, strings.TrimSpace(strings.ToLower(s)))
	}
	return returnSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, param string) error
}

func commandExit(config *Config, param string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(config *Config, param string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range config.commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(config *Config, param string) error {
	config.goNext = true
	return getDataFromApi(config)
}

func commandMapb(config *Config, param string) error {
	if config.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	config.goNext = false

	return getDataFromApi(config)
}

func getDataFromApi(config *Config) error {
	areaData, err := getLocationAreaData(config)

	if err != nil {
		return err
	}

	for _, loc := range areaData {
		fmt.Printf("%s\n", string(loc))
	}

	return nil
}

func commandExplore(config *Config, areaName string) error {
	if areaName == "" {
		return fmt.Errorf("Please enter a location to explore.\n")
	}

	config.commandCallbackarg = areaName
	pokemonNames, err := getSpecificLocationPokemonData(config)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonNames {
		fmt.Printf("  - %s\n", pokemon)
	}

	return nil
}

func commandCatch(config *Config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("Please enter a Pokemon to attempt catching.\n")
	}

	config.commandCallbackarg = pokemonName
	pokemon, err := getPokemonData(config)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.Intn(pokemon.BaseExperience) < 50 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.caughtPokemon[pokemon.Name] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)

	return nil
}

func commandInspect(config *Config, pokemonName string) error {
	pokeData, ok := config.caughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	pokeData.printPokeData()
	return nil
}

func commandPokedex(config *Config, param string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range config.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
