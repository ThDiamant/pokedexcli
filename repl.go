package main

import (
	"fmt"
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
	callback    func(config *Config) error
}

func commandExit(config *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(config *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range config.commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(config *Config) error {
	if err := getLocationAreaData(config); err != nil {
		return err
	}

	for _, loc := range config.currentLocs {
		fmt.Printf("%s\n", loc)
	}

	return nil
}
