package main

import (
	"testing"
)

func TestClearInp(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hellos world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}
	for _, cs := range cases {
		actual := clearInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Length Are not equal")
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v is not equal to %v", actualWord, expectedWord)
			}
		}
	}
}
