package Functions

import (
	"testing"
)

//Testing TestGetExtendedNeighbors
type testpairExtendedNeighbors struct {
	Pattern  string
	AdjList  map[string][]string
	MaxK     int
	Expected []string
}

/* HELPERS */
func BuildGraphFromStrings(branches [][]string, graph map[string][]string) map[string][]string {
	for _, path := range branches {
		for i := 0; i < len(path)-1; i++ {
			curr := path[i]
			next := path[i+1]
			graph[curr] = append(graph[curr], next)
		}
	}
	return graph
}

func SetEqual(set1 []string, set2 []string) bool {
	for _, elem := range set1 {
		//By the time we call SetEqual(), we already know Contains() works.
		if !Contains(set2, elem) {
			return false
		}
	}
	return true
}

/* TESTS */
func TestGetExtendedNeighbors(t *testing.T) {
	graph1 := make(map[string][]string)
	branches1 := [][]string{
		{"A", "B", "C", "D", "E"},
		{"B", "F", "G"},
	}
	graph1 = BuildGraphFromStrings(branches1, graph1)

	graph2 := make(map[string][]string)
	branches2 := [][]string{
		{"A", "B"},
		{"M", "N"},
		{"P", "Q"},
	}
	graph2 = BuildGraphFromStrings(branches2, graph2)

	//Test if function removes pattern from the extended neighbors
	graph3 := make(map[string][]string)
	branches3 := [][]string{
		{"A", "B", "C", "D", "E"},
		{"D", "A"},
		{"C", "B", "A"},
	}
	graph3 = BuildGraphFromStrings(branches3, graph3)

	results := [][]string{
		{"F", "C"},
		{"F", "C", "G", "D", "E"},
		{},
		{"C", "D", "E", "B"},
	}

	tests := []testpairExtendedNeighbors{
		{"A", graph1, 2, results[0]},
		{"A", graph1, 4, results[1]},
		{"A", graph1, 7, results[1]},
		{"B", graph2, 2, results[2]},
		{"A", graph3, 5, results[3]},
	}

	for i, test := range tests {
		outcome := GetExtendedNeighbors(test.Pattern, test.AdjList, test.MaxK)
		if !SetEqual(outcome, test.Expected) {
			t.Error("GetExtendedNeighbors() failed on testcase# ", i+1, "!")
		}
	}
}
