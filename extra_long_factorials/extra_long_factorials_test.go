package extralongfactorials

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in       int32
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "6"},
		{4, "24"},
		{5, "120"},
		{10, "3628800"},
		{15, "1307674368000"},
		{20, "2432902008176640000"},
		{25, "15511210043330985984000000"},
		{30, "265252859812191058636308480000000"},
		{50, "30414093201713378043612608166064768844377641568960512000000000000"},
		{100, "93326215443944152681699238856266700490715968264381" +
			"62146859296389521759999322991560894146397615651828" +
			"6253697920827223758251185210916864000000000000000000000000"},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("%d! = %s", tc.in, tc.expected), func(t *testing.T) {
			t.Parallel()
			if n := extraLongFactorials(tc.in); n != tc.expected {
				t.Errorf("%d! should eq %s, got %s", tc.in, tc.expected, n)
			}
		})
	}
}
