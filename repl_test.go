package main

import (
	"testing"
	"strings"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		name string
		input string
		expected []string
	} {
		{
			name:		"1",
			input: 		"    hello world    ",
			expected: 	[]string{"hello", "world"},
		},
		{
			name:		"2",
			input: 		"Charmander Bulbasaur PIKACHU",
			expected: 	[]string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name:		"3",
			input:		"OnE    Two  THREE",
			expected:	[]string{"one", "two", "three"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				"%s: actual (%d) and expected (%d) have different lengths",
				c.name, len(actual), len(c.expected),
			)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if strings.Compare(word, expectedWord) != 0 {
				t.Errorf(
					"%s: actual (%s) is different from expected (%s)",
					c.name,
					word,
					expectedWord,
				)
			}
		}
	}
}