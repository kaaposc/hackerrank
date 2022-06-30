package betweenTwoSets

import (
	"fmt"
	"testing"
)

func TestLowestCommon(t *testing.T) {
	var testCases = []struct{ inA, inB, expected int32 }{
		{1, 1, 1},
		{1, 2, 2},
		{2, 3, 6},
		{3, 5, 15},
		{4, 6, 12},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("LCD of %d and %d", tc.inA, tc.inB), func(t *testing.T) {
			if n := lowestCommon(tc.inA, tc.inB); n != tc.expected {
				t.Errorf("LCD of %d and %d is %d, got %d", tc.inA, tc.inB, tc.expected, n)
			}
		})
	}
}

func TestGreatestCommon(t *testing.T) {
	var testCases = []struct{ inA, inB, expected int32 }{
		{1, 1, 1},
		{1, 2, 1},
		{2, 2, 2},
		{2, 4, 2},
		{4, 6, 2},
		{5, 7, 1},
		{10, 16, 2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("GCD of %d and %d", tc.inA, tc.inB), func(t *testing.T) {
			if n := greatestCommon(tc.inA, tc.inB); n != tc.expected {
				t.Errorf("GCD of %d and %d is %d, got %d", tc.inA, tc.inB, tc.expected, n)
			}
		})
	}
}

func TestSets(t *testing.T) {
	var testCases = []struct {
		inA      []int32
		inB      []int32
		expected int32
	}{
		{
			inA:      []int32{1},
			inB:      []int32{2},
			expected: 2,
		},
		{
			inA:      []int32{2, 3},
			inB:      []int32{12, 18},
			expected: 1,
		},
		{
			inA:      []int32{2, 4},
			inB:      []int32{16, 32, 64},
			expected: 3,
		},
		{
			inA:      []int32{100, 99, 98, 97, 96, 95, 94, 93, 92, 91},
			inB:      []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("between %v and %v", tc.inA, tc.inB), func(t *testing.T) {
			if n := getTotalX(tc.inA, tc.inB); n != tc.expected {
				t.Errorf("between %v and %v are %d ints, got %d", tc.inA, tc.inB, tc.expected, n)
			}
		})
	}
}
