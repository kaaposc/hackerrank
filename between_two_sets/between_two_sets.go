// Package betweenTwoSets solves Hackerrank challenge "Between Two Sets"[1].
//
// [1] https://www.hackerrank.com/challenges/between-two-sets/problem
package betweentwosets

// lowestCommon calculates lowest common denominator of two integers.
func lowestCommon(a int32, b int32) int32 {
	// LCD(x, 1) => x
	if a == 1 {
		return b
	}

	if b == 1 {
		return a
	}

	l, r := a, b
	for l != r {
		if l < r {
			l += a
		} else {
			r += b
		}
	}

	return l
}

// greatestCommon calculates greatest common divisor of two integers.
func greatestCommon(a int32, b int32) int32 {
	// GCD(x, 1) => 1
	if a == 1 || b == 1 {
		return 1
	}

	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}

		// early exit as soon as any one drops to 1
		if a == 1 || b == 1 {
			return 1
		}
	}

	return a
}

// getTotalX counts integers between two sets that are both LCD of a and GCD of b.
func getTotalX(setA []int32, setB []int32) int32 {
	// if at least one int from set A is greater than any of set B ints return 0
	for _, i := range setA {
		for _, j := range setB {
			if i > j {
				return 0
			}
		}
	}

	lCA := int32(1)
	gCB := setB[0]

	for _, n := range setA {
		lCA = lowestCommon(lCA, n)
	}

	for _, n := range setB {
		gCB = greatestCommon(gCB, n)
	}

	if lCA > gCB || gCB%lCA != 0 {
		return 0
	}

	var (
		i int32
		n = lCA
	)

	for n <= gCB {
		if greatestCommon(n, gCB) == n {
			i++
		}

		n += lCA
	}

	return i
}
