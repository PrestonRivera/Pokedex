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
			input: "I really LIKE PIZza",
			expected: []string{"i", "really", "like", "pizza"},
		},
		{
			input: "have a good day son",
			expected: []string{"have", "a", "good", "day", "son"},
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