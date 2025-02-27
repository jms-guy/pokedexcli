package main

import (
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 	string
		expected []string
	}{
		{
			input:	"  hello  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:	"Broke     and even  ",
			expected: []string{"broke", "and", "even"},
		},
		{
			input:	" Bulbasaur  Chikorita   123  an    hour   ",
			expected: []string{"bulbasaur", "chikorita", "123", "an", "hour"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println(actual)
		if len(actual) != len(c.expected) {
			t.Errorf("Result was unexpected:%v slice length was: %d wanted slice length: %d", actual, len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Result was unexpected: word was: %s wanted word: %s", word, expectedWord)
			}
		}
	}
}