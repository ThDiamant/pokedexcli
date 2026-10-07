package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal"
	"time"
)

// TODO: General refactor
// TODO: Modify catch probabilities so catching a pokemon sometimes fails

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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores the given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
	}
}

func repl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanInput := cleanInput(scanner.Text())

		command, ok := config.commandRegistry[cleanInput[0]]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}

		var param string
		if len(cleanInput) > 1 {
			param = cleanInput[1]
		}

		err := command.callback(config, param)
		if err != nil {
			fmt.Printf("Error while running command %s: %v", command.name, err)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s\n", err)
		}
	}
}

func main() {
	const interval = 5 * time.Second

	config := Config{
		commandRegistry: getCommands(),
		next:            locationBaseURL,
		cache:           internal.NewCache(interval),
		caughtPokemon:   make(map[string]Pokemon),
	}
	repl(&config)
}
