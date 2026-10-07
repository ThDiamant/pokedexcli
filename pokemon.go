package main

import "fmt"

type Pokemon struct {
	Name           string
	Height         int
	Weight         int
	BaseExperience int
	Stats          PokemonStats
	Types          []string
}

type PokemonStats struct {
	Hp             int
	Attack         int
	Defence        int
	SpecialAttack  int
	SpecialDefence int
	Speed          int
}

func (pokemon *Pokemon) printPokeData() {
	fmt.Printf(
		"Name: %s\nHeight: %d\nWeight: %d\nStats:\n -hp: %d\n -attack: %d\n -defence: %d\n -special-attack: %d\n -special-defence: %d\n -speed: %d\nTypes:\n",
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight,
		pokemon.Stats.Hp,
		pokemon.Stats.Attack,
		pokemon.Stats.Defence,
		pokemon.Stats.SpecialAttack,
		pokemon.Stats.SpecialDefence,
		pokemon.Stats.Speed,
	)
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" -%s\n", pokeType)
	}
}
