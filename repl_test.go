package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello    world   ",
			expected: []string{"hello", "world"},
		},
	}

	for _, v := range cases {
		has := cleanInput(v.input)

		if len(has) != len(v.expected) {
			t.Errorf("the size of the expected slice is %d and has %d", len(v.expected), len(has))
		}

		for i := range has {
			word := has[i]
			expectedWord := v.expected[i]

			if word != expectedWord {
				t.Errorf("The expectedWord: %s doesn't match current word: %s", expectedWord, word)
			}
		}
	}
}
