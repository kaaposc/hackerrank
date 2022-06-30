// Package betweenTwoSets solves Hackerrank challenge "Between Two Sets"[1].
//
// [1] https://www.hackerrank.com/challenges/between-two-sets/problem
package betweenTwoSets

// lowestCommon calculates lowest common denominator of two integers.
func lowestCommon(a int32, b int32) int32 {
	// LCD(x, 1) => x
	if a == 1 {
		return b
	}

	if b == 1 {
		return a
	}

	var l, r int32 = a, b
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
func getTotalX(a []int32, b []int32) int32 {
	// if at least one int from set A is greater than any of set B ints return 0
	for _, i := range a {
		for _, j := range b {
			if i > j {
				return 0
			}
		}
	}

	var (
		lCA int32 = 1
		gCB int32 = b[0]
	)

	for _, n := range a {
		lCA = lowestCommon(lCA, n)
	}

	for _, n := range b {
		gCB = greatestCommon(gCB, n)
	}

	if lCA > gCB || gCB%lCA != 0 {
		return 0
	}

	var (
		i int32
		n int32 = lCA
	)

	for n <= gCB {
		if greatestCommon(n, gCB) == n {
			i++
		}

		n += lCA
	}

	return i
}
