package main

type Config struct {
	commandRegistry map[string]cliCommand
	next            string
	prev            string
	currentLocs     []string
	goNext          bool
}
