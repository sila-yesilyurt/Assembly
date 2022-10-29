package Functions

import "testing"

//Testing AverageOutDegree
type testpairAverageOutdegree struct {
	Graph    map[string][]string
	Expected float64
}

func TestAverageOutDegree(t *testing.T) {
	graph0 := map[string][]string{
		"compBio": {"metagenomics", "alignment", "assembly", "phylogenetics"},
	}
	graph1 := map[string][]string{
		"ATTCCC":  {"ACC", "ATTCA", "AGC", "AGG"},
		"AGGCT":   {"ATCGGA"},
		"AGGCGTA": {"AATCC"},
	}
	graph2 := map[string][]string{
		"A": {"B", "C", "D", "E"},
		"C": {"I", "J", "K", "L"},
		"X": {"Y", "Z"},
		"Z": {"A", "J"},
	}
	graph3 := map[string][]string{
		"red":    {"lily", "rose"},
		"blue":   {"moon", "ocean", "sky"},
		"yellow": {"cheese", "corn"},
		"gray":   {"eyes"},
		"purple": {"umbrella", "uniform", "grapes"},
		"green":  {"leaves"},
	}

	tests := []testpairAverageOutdegree{
		{graph0, 4}, //testcase1
		{graph1, 2}, //testcase2
		{graph2, 3}, //testcase3
		{graph3, 2}, //testcase4
	}

	for i, test := range tests {
		result := AverageOutDegree(test.Graph)
		if result != test.Expected {
			t.Error("AverageOutDegree() failed testcase#", i+1, "\n",
				"Expecting:", test.Expected, "Actual outcome:", result)
		}
	}
}
