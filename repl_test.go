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
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "good Morning world ",
			expected: []string{"good", "morning", "world"},
		},
		{
			input: "apples are cheap",
			expected: []string{"apples", "are", "cheap"},
		},
	}

	for _, c:= range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Input \"%v\" wasn't cleaned successfully", actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not match %v", word, expectedWord)
			}
		}
	}
}