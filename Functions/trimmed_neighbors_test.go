package Functions

import (
	"testing"
	// "fmt"
)

//Testing GetTrimmedNeighbors
type testpairTransitivity4 struct {
	Pattern  string
	Graph    map[string][]string
	MaxK     int
	Expected []string
}

/* TESTS */

func TestGetTrimmedNeighbors(t *testing.T) {
	//Simple test case
	/*
	   A -> C
	   A -> B -> C -> D -> E
	        B -----------> E
	   A-C, B-E are transitive.
	*/
	graph1 := make(map[string][]string)
	branches1 := [][]string{
		{"A", "B", "C", "D", "E", "G"},
		{"B", "E"},
		{"A", "C"},
	}
	graph1 = BuildGraphFromStrings(branches1, graph1)

	//Neighbor's neighbor
	/*
	   A -> E -> C
	   A -> B -> C
	   A -> D <- C
	*/
	graph2 := make(map[string][]string)
	branches2 := [][]string{
		{"A", "B", "C"},
		{"A", "E", "C"},
		{"A", "D"},
		{"C", "D"},
	}
	graph2 = BuildGraphFromStrings(branches2, graph2)

	//Self-loop edge case
	/*
	   A <- B
	   A -> B -> D
	   A -> C
	   A-C should remain in neighbors.
	*/
	graph3 := make(map[string][]string)
	branches3 := [][]string{
		{"A", "B", "D"},
		{"B", "A"},
		{"A", "C"},
	}
	graph3 = BuildGraphFromStrings(branches3, graph3)

	graph4 := make(map[string][]string)
	branches4 := [][]string{
		{"A", "B", "C", "D"},
		{"C", "B", "A"},
	}
	graph4 = BuildGraphFromStrings(branches4, graph4)

	//Don't over trim.
	graph5 := make(map[string][]string)
	branches5 := [][]string{
		{"A", "B", "C", "D", "E", "F"},
		{"A", "D"},
		{"A", "E"},
		{"A", "F"},
	}
	graph5 = BuildGraphFromStrings(branches5, graph5)

	tests := []testpairTransitivity4{
		{"A", graph1, 3, []string{"B"}},      //testcase1
		{"B", graph1, 2, []string{"C", "E"}}, //testcase2
		{"B", graph1, 3, []string{"C"}},      //testcase3
		{"A", graph2, 4, []string{"B", "E"}}, //testcase4
		{"A", graph3, 2, []string{"B", "C"}}, //testcase5
		{"A", graph4, 3, []string{}},         //testcase6
		{"C", graph4, 5, []string{"D"}},      //testcase7
		{"A", graph5, 2, []string{"B", "D"}}, //testcase8
		{"A", graph5, 3, []string{"B"}},      //testcase9
	}

	for i, test := range tests {
		result := GetTrimmedNeighbors(test.Pattern, test.Graph, test.MaxK)
		if !StrSliceEq(result, test.Expected) {
			t.Error("GetTrimmedNeighbors() failed test case#", i+1, "\n",
				"Failed on:", test.Pattern, test.Graph, test.MaxK, "\n",
				"Expecting:", test.Expected, "Actual result:", result)
		}
	}
}
