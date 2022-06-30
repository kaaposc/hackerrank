// Package anagram solves Hackerrank challenge "Anagram"[1].
//
// [1] https://www.hackerrank.com/challenges/anagram/problem
package anagram

import "strings"

// anagram takes string argument and returns steps needed to make anagram.
func anagram(str string) int32 {
	if len(str)%2 == 1 {
		return -1
	}

	left := str[:len(str)/2]
	right := str[len(str)/2:]

	for _, ch := range left {
		if b, a, found := strings.Cut(right, string(ch)); found {
			right = b + a
		}
	}

	return int32(len(right))
}
