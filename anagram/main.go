package anagram

func anagram(s string) int32 {
	if len(s)%2 == 1 {
		return -1
	}

	left := s[:len(s)/2]
	right := s[len(s)/2:]

	var search = func(s string, r rune) int {
		for i, ch := range s {
			if ch == r {
				return i
			}
		}

		return -1
	}

	for _, ch := range left {
		if p := search(right, ch); p != -1 {
			right = right[:p] + right[p+1:]
		}
	}

	return int32(len(right))
}
