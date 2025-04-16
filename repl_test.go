package main

import (
	"github.com/SyncTank/pokedex/pokeAPI"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello World ",
			expected: []string{"Hello", "World"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Length of slice doesn't match. Expected %d, got %d", len(c.expected), len(actual))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match at index %d. Expected '%s', got '%s'", i, expectedWord, word)
				return
			}
		}
	}
}

func TestPokeEndPoint(t *testing.T) {
	_, err := pokeAPI.GetLocation(pokeAPI.Endpoint)
	if err != nil {
		t.Errorf("Error reaching endpoint:")
		return
	}
}
