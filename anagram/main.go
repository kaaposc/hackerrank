package anagram

import (
	"strings"
)

func anagram(s string) int32 {
	if len(s)%2 == 1 {
		return -1
	}

	left := strings.Split(s[:len(s)/2], "")
	right := strings.Split(s[len(s)/2:], "")

	var search = func(s []string, needle string) int {
		for i, ch := range s {
			if ch == needle {
				return i
			}
		}

		return -1
	}

	for _, ch := range left {
		if p := search(right, ch); p != -1 {
			right = append(right[:p], right[p+1:]...)
		}
	}

	return int32(len(right))
}
