package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello hi how are you doing",
			expected: []string{"hello", "hi", "how", "are", "you", "doing"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println(actual)
		if len(actual) != len(c.expected) {
			t.Errorf("error while seperating and formatting inputs: %v != %v\n check your inputs only use 1 space to separate inputs", len(actual), len(c.expected))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("result of cleanInput() function doesn't match expected result %v != %v", word, expectedWord)
				t.Fail()
			}
		}
	}
}
