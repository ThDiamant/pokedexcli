package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)

		command, ok := config.commandRegistry[cleanInput[0]]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}

		err := command.callback(config)
		if err != nil {
			fmt.Printf("Error while running command %s: %v", command.name, err)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s\n", err)
		}
	}
}

func main() {
	config := Config{
		commandRegistry: map[string]cliCommand{
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
		},
	}
	repl(&config)
}
