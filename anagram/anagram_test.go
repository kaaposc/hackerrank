package anagram

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	var testCases = []struct {
		in       string
		expected int32
	}{
		{in: "abc", expected: -1},
		{in: "xyyx", expected: 0},
		{in: "ab", expected: 1},
		{in: "aaabbb", expected: 3},
		{in: "abccde", expected: 2},
		{in: "mnop", expected: 2},
		{in: "xaxbbbxx", expected: 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("make anagram from `%s`", tc.in), func(t *testing.T) {
			if steps := anagram(tc.in); steps != tc.expected {
				t.Errorf("`%s` to anagram in %d steps, got %d", tc.in, tc.expected, steps)
			}
		})
	}
}
