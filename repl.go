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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
