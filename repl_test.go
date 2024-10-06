package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%s) == %v, expected %v", c.input, actual, c.expected)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%s) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
