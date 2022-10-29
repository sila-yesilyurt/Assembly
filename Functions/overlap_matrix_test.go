package Functions

import (
	"testing"
)

//Test OverlapScoringMatrix
type testpairOverlapMatrix1 struct {
	Reads    []string
	Match    float64
	Mis      float64
	Gap      float64
	Expected [][]float64
}

/* TESTS */
var reads0 []string = []string{
	"ATTGA",
	"TTGAC",
	"TGACA",
	"GACAG",
}

var reads1 []string = []string{
	"ATTGGA",
	"TGGAGG",
	"TGGGAA",
	"CCGGATT",
}

//Note: the 3 reads in this set align well with each other.
var reads2 []string = []string{
	"ATTCGATCCGTC",
	"ATCCGATCCGAG",
	"CGATCCGTC",
}

func TestOverlapScoringMatrix(t *testing.T) {
	//Expected Answers
	ans0 := [][]float64{
		[]float64{0, 4, 3, 2},
		[]float64{2, 0, 4, 3},
		[]float64{1, 2, 0, 4},
		[]float64{0, 0, 2, 0},
	}
	ans1 := [][]float64{
		[]float64{0, 4, 3, 1},
		[]float64{1, 0, 2, 0},
		[]float64{1, 2, 0, 0},
		[]float64{3, 1, 1, 0},
	}
	ans2 := [][]float64{
		[]float64{0, 8, 7, 4},
		[]float64{4, 0, 6, 2},
		[]float64{4, 6, 0, 3},
		[]float64{6, 3, 2, 0},
	}
	ans3 := [][]float64{
		[]float64{0, 6, 9},
		[]float64{6, 0, 5},
		[]float64{6, 6, 0},
	}

	tests := []testpairOverlapMatrix1{
		{reads0, 1, 1, 1, ans0}, //testcase 1
		{reads1, 1, 1, 1, ans1}, //testcase 2
		{reads1, 2, 5, 1, ans2}, //testcase 3
		{reads2, 1, 1, 1, ans3}, //testcase 4
	}

	for i, test := range tests {
		result := OverlapScoringMatrix(test.Reads, test.Match, test.Mis, test.Gap)
		if !ScoringMatrixEq(result, test.Expected) {
			t.Error("OverlapScoringMatrix() failed testcase#", i+1, "\n",
				"Expecting:", test.Expected, "Actual outcome:", result)
		}
	}
}
