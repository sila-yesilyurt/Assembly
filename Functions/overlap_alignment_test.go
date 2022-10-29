package Functions

import (
	"math"
	"testing"
)

//Testing OverlapScoreTable
type testpairOverlapAlignment1 struct {
	Str0     string
	Str1     string
	Match    float64
	Mis      float64
	Gap      float64
	Expected [][]float64
}

//Testing ScoreOverlapAlignment
type testpairOverlapAlignment2 struct {
	Str0     string
	Str1     string
	Match    float64
	Mis      float64
	Gap      float64
	Expected float64
}

//Global var: minimum int
var negInf float64 = float64(math.MinInt64)

/* HELPER */
func ScoringMatrixEq(mtx1, mtx2 [][]float64) bool {
	if len(mtx1) != len(mtx2) {
		return false
	}
	for i := range mtx1 {
		if len(mtx1[i]) != len(mtx2[i]) {
			return false
		}
		for j := range mtx1[i] {
			if mtx1[i][j] != mtx2[i][j] {
				return false
			}
		}
	}
	return true
}

/* TESTS */
//Testing OverlapScoreTable
func TestOverlapScoreTable(t *testing.T) {
	str1 := "ATCC"
	str2 := "CCAG"
	str3 := "ATTGGCG"
	str4 := "GGCGTAT"
	str5 := "GCAGGAT"

	ans1 := [][]float64{
		[]float64{0, -1, -2, -3, -4},
		[]float64{0, -1, -2, -1, -2},
		[]float64{0, -1, -2, -2, -2},
		[]float64{0, 1, 0, -1, -2},
		[]float64{0, 1, 2, 1, 2},
	}
	ans2 := [][]float64{
		[]float64{0, -5, -10, -15, -20},
		[]float64{0, -2, -7, -9, -14},
		[]float64{0, -2, -4, -9, -11},
		[]float64{0, 1, -1, -6, -11},
		[]float64{0, 1, 2, -3, 2},
	}
	ans3 := [][]float64{
		[]float64{0, -1, -2, -3, -4},
		[]float64{0, 1, 0, -1, -2},
		[]float64{0, 0, 2, 1, 0},
		[]float64{0, -1, 1, 3, 2},
		[]float64{0, -1, 0, 2, 4},
	}
	ans4 := [][]float64{
		[]float64{0, -2, -4, -6, -8, -10, -12, -14},
		[]float64{0, -2, -4, -6, -8, -10, -9, -11},
		[]float64{0, -2, -4, -6, -8, -7, -9, -8},
		[]float64{0, -2, -4, -6, -8, -7, -9, -8},
		[]float64{0, 1, -1, -3, -5, -7, -9, -10},
		[]float64{0, 1, 2, 0, -2, -4, -6, -8},
		[]float64{0, -1, 0, 3, 1, -1, -3, -5},
		[]float64{0, 1, 0, 1, 4, 2, 0, 4},
	}
	ans5 := [][]float64{
		[]float64{0, -4, -8, -12, -16, -20, -24, -28},
		[]float64{0, -3, -7, -7, -11, -15, -19, -23},
		[]float64{0, -3, -6, -10, -10, -14, -18, -18},
		[]float64{0, -3, -6, -9, -13, -13, -17, -17},
		[]float64{0, 1, -3, -7, -8, -12, -16, -20},
		[]float64{0, 1, -2, -6, -6, -7, -11, -15},
		[]float64{0, -3, 2, -2, -6, -9, -10, -14},
		[]float64{0, 1, -2, -1, -1, -5, -9, 1},
	}

	tests := []testpairOverlapAlignment1{
		{str1, str2, 1, 1, 1, ans1}, //testcase1
		{str1, str2, 1, 2, 5, ans2}, //testcase2
		{str1, str1, 1, 1, 1, ans3}, //testcase3
		{str3, str4, 1, 2, 2, ans4}, //testcase4
		{str3, str5, 1, 3, 4, ans5}, //testcase5
	}

	for i, test := range tests {
		result := OverlapScoreTable(test.Str0, test.Str1, test.Match, test.Mis, test.Gap)
		if !ScoringMatrixEq(result, test.Expected) {
			t.Error("OverlapScoreTable failed on testcase#", i+1, "\n",
				"Inputs:", test.Str0, test.Str1, test.Match, test.Mis, test.Gap, "\n",
				"Expecting:", test.Expected, "Actual outcome:", result)
		}
	}
}

//Testing ScoreOverlapAlignment
func TestScoreOverlapAlignment(t *testing.T) {
	str1 := "AAATGCCG"
	str2 := "TCATGCGC"
	str3 := "AATTCGAA"
	str4 := "CCAGGTCCGGACGGAT"
	str5 := "AGGGTCCGGACGGATTTAC"

	tests := []testpairOverlapAlignment2{
		{str1, str2, 4, 1, 5, 13}, //testcase1
		{str2, str1, 4, 1, 5, 13}, //testcase2
		{str1, str3, 2, 2, 5, 3},  //testcase3
		{str1, str2, 2, 3, 5, 0},  //testcase4
		{str3, str4, 2, 1, 5, 2},  //testcase5
		{str4, str5, 2, 1, 5, 24}, //testcase6
		{str5, str4, 2, 1, 5, 4},  //testcase7
	}

	for i, test := range tests {
		result := ScoreOverlapAlignment(test.Str0, test.Str1, test.Match, test.Mis, test.Gap)
		if result != test.Expected {
			t.Error("ScoreOverlapAlignment failed on testcase#", i+1, "\n",
				"Expecting:", test.Expected, "Actual outcome:", result)
		}
	}
}
