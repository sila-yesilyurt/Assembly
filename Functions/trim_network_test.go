package Functions

import (
	"testing"
	// "fmt"
)

//Testing TrimGraph
type testpairTrim struct {
	Graph    map[string][]string
	MaxK     int
	Expected map[string][]string
}

/* TESTS */

func TestTrimNetwork(t *testing.T) {
	graph1 := make(map[string][]string)
	graph2 := make(map[string][]string)
	graph3 := make(map[string][]string)

	//Test1: simple cases
	branches1 := [][]string{
		{"A", "B", "C", "D", "E", "F", "G"},
		{"A", "C", "E", "G"},
		{"D", "G"},
	}
	graph1 = BuildGraphFromStrings(branches1, graph1)

	//Test2:
	//Remember to update the adjacency list.
	branches2 := [][]string{
		{"A", "B", "A", "C"},
		{"B", "C"},
	}
	graph2 = BuildGraphFromStrings(branches2, graph2)

	//Test3:
	//Trim everything component in the graph.
	branches3 := [][]string{
		{"A", "B", "C", "D"},
		{"A", "D"},
		{"E", "F", "G", "H", "I"},
		{"E", "H"},
		{"E", "I"},
	}
	graph3 = BuildGraphFromStrings(branches3, graph3)

	//Answers
	ans1 := make(map[string][]string)
	ans2 := make(map[string][]string)
	ans3 := make(map[string][]string)

	branches4 := [][]string{
		{"A", "B", "C", "D", "E", "F", "G"},
	}
	ans1 = BuildGraphFromStrings(branches4, ans1)

	branches5 := [][]string{
		{"A", "B"},
		{"B", "A"},
	}
	ans2 = BuildGraphFromStrings(branches5, ans2)

	branches6 := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H", "I"},
	}
	ans3 = BuildGraphFromStrings(branches6, ans3)

	tests := []testpairTrim{
		{graph1, 2, ans1},
		{graph1, 6, ans1},
		{graph2, 3, ans2},
		{graph3, 5, ans3},
	}
	for i, test := range tests {
		result := TrimNetwork(test.Graph, test.MaxK)
		if !GraphEq(test.Expected, result) {
			t.Error("TrimNetwork() failed testcase#", i+1, "\n",
				"Failed on:", test.Graph, test.MaxK, "\n",
				"Expecting:", test.Expected, "Actual restul:", result)
		}
	}
}
