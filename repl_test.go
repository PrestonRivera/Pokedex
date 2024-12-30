package main

import (
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Pichu, Bulbasaur, Squirtle ",
			expected: []string{"pichu", "bulbasaur", "squirtle"},
		},
		{
			input: " ",
			expected: []string{},
		},
		{
			input: "ONE, TWO, THREE",
			expected: []string{"one", "two", "three"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected length %d, got %d for input %q", len(c.expected), len(actual), c.input)
		}
		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf("expected %v, got %v for input %q", c.expected, actual, c.input)
				break
			}
		}
	}
}