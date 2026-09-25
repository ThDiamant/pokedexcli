package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var returnSlice []string
	for _, s := range (strings.Split(strings.TrimSpace(text), " ")) {
		if s == "" {
			continue
		}
		returnSlice = append(returnSlice, strings.TrimSpace(strings.ToLower(s)))
	}
	return returnSlice
}