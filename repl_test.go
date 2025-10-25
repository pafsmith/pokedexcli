package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "basic input with spaces",
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "single word",
			input:    "help",
			expected: []string{"help"},
		},
		{
			name:     "uppercase input",
			input:    "EXPLORE location-area",
			expected: []string{"explore", "location-area"},
		},
		{
			name:     "mixed case",
			input:    "CaTcH Pikachu",
			expected: []string{"catch", "pikachu"},
		},
		{
			name:     "multiple spaces",
			input:    "catch    pikachu",
			expected: []string{"catch", "", "", "", "pikachu"},
		},
		{
			name:     "command with multiple args",
			input:    "command arg1 arg2 arg3",
			expected: []string{"command", "arg1", "arg2", "arg3"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "only spaces",
			input:    "   ",
			expected: []string{""},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("expected length %d, got %d", len(c.expected), len(actual))
				return
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("at index %d: expected %q, got %q", i, expectedWord, word)
				}
			}
		})
	}
}
