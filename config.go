package main

import (
	"pokedexcli/internal"
)

type Config struct {
	commandRegistry map[string]cliCommand
	next            string
	prev            string
	goNext          bool
	cache           *internal.Cache
}
