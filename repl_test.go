package main

import (
	"testing"
	"strings"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	} {
		{
			input: 		"    hello world    ",
			expected: 	[]string{"hello", "world"},
		},
		{
			input: 		"Charmander Bulbasaur PIKACHU",
			expected: 	[]string{"charmander", "bulbasaur", "pikachu"},
		},
		// more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual (%d) and expected (%d) have different lengths", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if strings.Compare(word, expectedWord) != 0 {
				t.Errorf("actual (%s) is different from expected (%s)", word, expectedWord)
			}
		}
	}
}