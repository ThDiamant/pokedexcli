package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandRegistry := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)
		
		command, ok := commandRegistry[cleanInput[0]]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("Error while running command %s: %v", command.name, err)
		}


		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s\n", err)
		}
	}
}

