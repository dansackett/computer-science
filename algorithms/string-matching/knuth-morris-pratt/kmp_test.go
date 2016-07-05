package kmp

import "testing"

type TestData struct {
	Text, Pattern string
	Expected      int
}

func TestIt(t *testing.T) {
	testInput := []TestData{
		TestData{
			Text:     "ABACAABACABACAB",
			Pattern:  "ABACAB",
			Expected: 5,
		},
		TestData{
			Text:     "ABCDABD",
			Pattern:  "ABD",
			Expected: 4,
		},
		TestData{
			Text:     "BACBABABAABCBAB",
			Pattern:  "ABABABCA",
			Expected: -1,
		},
	}

	for _, input := range testInput {
		result := KMP(input.Text, input.Pattern)
		if result != input.Expected {
			t.Errorf("Expected index to be %d for pattern '%s' in '%s' found %d", input.Expected, input.Pattern, input.Text, result)
		}
	}
}
