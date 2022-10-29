package Functions

import "testing"

func StringIndexEq(index1, index2 StringIndex) bool {
	if len(index1) != len(index2) {
		return false
	}
	for i := range index1 {
		if !IntSliceEq(index1[i], index2[i]) {
			return false
		}
	}
	return true
}

/* TESTS */

//Testing Updated Version of MapToMinimizer
func TestMapToMinimizer(t *testing.T) {
	reads1 := []string{
		"ACCGATC",
		"AGGTGGG",
		"CAGAGCC",
		"GGACGTA",
		"TTCGAGG",
		"TTGACCA",
	}

	ans1 := StringIndex{
		"AC": {0, 3, 5},
		"AG": {1, 2, 4},
		"AT": {0},
		"CA": {5},
		"CC": {2},
		"CG": {0, 4},
		"GA": {3},
		"GG": {1},
		"GT": {1, 3},
		"TC": {4},
		"TG": {5},
	}
	ans2 := StringIndex{
		"AC": {0, 3, 5},
		"AG": {1, 2, 4},
		"AT": {0},
		"CG": {3, 4},
		"GA": {5},
		"GG": {1},
	}

	//testcase#1
	result1 := MapToMinimizer(reads1, 2, 3)
	if !StringIndexEq(result1, ans1) {
		t.Error("MapToMinimizer failed testcase1", "Failed on:", reads1, 2, 3,
			"\n", "Expecting:", ans1, "Actual result:", result1)
	}

	//testcase#2
	result2 := MapToMinimizer(reads1, 2, 2)
	if !StringIndexEq(result2, ans2) {
		t.Error("MapToMinimizer failed testcase1", "Failed on:", reads1, 3, 2,
			"\n", "Expecting:", ans2, "Actual result:", result2)
	}
}
