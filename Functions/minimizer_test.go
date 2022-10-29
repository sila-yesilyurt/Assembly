package Functions

import "testing"

type testpairMinimizer struct {
	Text     string
	K        int
	Expected string
}

/* HELPER */
func IntSliceEq(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

/* TESTS */
func TestMinimizer(t *testing.T) {
	texts := []string{
		"TTTGACCGG",
		"ATCACTCCCTCAGGGAGTTACCC",
		"CCGTGATGATTACCGA",
		"AATAAAGAGGCCTTC",
	}

	tests := []testpairMinimizer{
		{texts[0], 3, "ACC"},
		{texts[0], 6, "GACCGG"},
		{texts[1], 4, "ACCC"},
		{texts[1], 8, "ACTCCCTC"},
		{texts[2], 5, "ACCGA"},
		{texts[3], 4, "AAAG"},
	}

	for i, test := range tests {
		result := Minimizer(test.Text, test.K)
		if result != test.Expected {
			t.Error("Maximizer() failed testcase#", i+1, "\n",
				"Failed on:", test.Text, test.K, "\n",
				"Expected:", test.Expected, "Actual result:", result)
		}
	}
}
