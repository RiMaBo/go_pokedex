package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "\t\tPreceding TaBs and LeAdInG CaRrIaGe returns.\n\n",
			expected: []string{"preceding", "tabs", "and", "leading", "carriage", "returns."},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf(`---------------------------------
Lengths don't match.'
Expecting: %d elements
Actual:    %d elements
Fail
`, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf(`---------------------------------
cleanInput(%v)
Expecting: %s
Actual:    %v
Fail
`, c.input, expectedWord, word)
			}
		}
	}
}
