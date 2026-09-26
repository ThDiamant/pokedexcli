package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanInput[0])

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s\n", err)
		}
	}
}
