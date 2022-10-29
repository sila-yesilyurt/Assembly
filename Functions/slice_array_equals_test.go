package Functions

import "testing"

//Testing FreqMapEquals
type testpairSliceArray struct {
	Map1     map[string]int
	Map2     map[string]int
	Expected bool
}

//Testing FrequencyMap
type testpair2 struct {
	Patterns []string
	Expected map[string]int
}

//Testing StringSliceEquals
type testpair3 struct {
	Patterns1 []string
	Patterns2 []string
	Expected  bool
}

/* TESTING */
var patterns1 []string = []string{
	"ATTCGT",
	"AGGGGC",
	"AGGGCTT",
	"GGCGAA",
	"GGCGAA",
	"GGCGAA",
	"AGGGCTT",
	"TTTTA",
}

var patterns2 []string = []string{
	"TTTTA",
	"GGCGAA",
	"AGGGCTT",
	"GGCGAA",
	"AGGGCTT",
	"AGGGGC",
	"GGCGAA",
	"ATTCGT",
}

var patterns3 []string = []string{
	"ATTCGT",
	"AGGGGC",
	"GGCGAA",
	"GGCGAA",
	"GGCGAA",
	"TTTTA",
}

var patterns4 []string = []string{
	"ATTCGT",
	"AGGGGC",
	"GGCGAA",
	"GGCGAA",
	"GGCGAA",
	"TTTTA",
	"AGGGCTT",
}

/*
//Testing FreqyMapEquals
func TestFreqMapEquals(t *testing.T) {
	map1 := map[string]int{
		"A": 5,
		"B": 4,
		"C": 3,
		"D": 2,
		"E": 11,
	}
	map2 := map[string]int{
		"AA":  4,
		"A":   5,
		"B":   4,
		"CCC": 3,
		"C":   3,
		"D":   2,
		"E":   11,
	}
	map3 := map[string]int{
		"A": 5,
		"D": 2,
		"C": 3,
		"E": 11,
		"B": 4,
	}
	map4 := map[string]int{
		"A": 5,
		"B": 3,
		"C": 3,
		"D": 2,
		"E": 11,
	}
	map5 := map[string]int{
		"A": 5,
		"B": 4,
		"C": 3,
		"E": 11,
	}

	tests := []testpairSliceArray{
		{map1, map1, true},  //testcase1
		{map1, map3, true},  //testcase2
		{map3, map1, true},  //testcase3
		{map1, map2, false}, //testcase4
		{map1, map4, false}, //testcase5
		{map1, map5, false}, //testcase6
	}

	for i, test := range tests {
		result := FreqMapsEquals(test.Map1, test.Map2)
		if result != test.Expected {
			t.Error("FreqMapsEquals failed on testcase#", i+1, "\n",
				"Expecting:", test.Expected, "Actual outcome:", result)
		}
	}
}

//Testing FrequencyMap
func TestFrequencyMap(t *testing.T) {
	ans1 := map[string]int{
		"ATTCGT":  1,
		"AGGGGC":  1,
		"AGGGCTT": 2,
		"GGCGAA":  3,
		"TTTTA":   1,
	}
	ans2 := map[string]int{
		"TTTTA":   1,
		"GGCGAA":  3,
		"ATTCGT":  1,
		"AGGGGC":  1,
		"AGGGCTT": 2,
	}
	ans3 := map[string]int{
		"ATTCGT": 1,
		"AGGGGC": 1,
		"GGCGAA": 3,
		"TTTTA":  1,
	}
	ans4 := map[string]int{
		"ATTCGT":  1,
		"AGGGGC":  1,
		"GGCGAA":  3,
		"TTTTA":   1,
		"AGGGCTT": 1,
	}

	tests := []testpair2{
		{patterns1, ans1}, //testcase1
		{patterns2, ans2}, //testcase2
		{patterns3, ans3}, //testcase3
		{patterns4, ans4}, //testcase4
	}

	for i, test := range tests {
		result := FrequencyMap(test.Patterns)
		if !FreqMapsEquals(result, test.Expected) {
			t.Error("FrequencyMap failed on testcase#", i+1, "\n",
				"Expecting:", test.Expected, "Actual outcome:", result)
		}
	}
}
*/

//Testing StringSliceEquals
func TestStringSliceEquals(t *testing.T) {
	tests := []testpair3{
		{patterns1, patterns1, true},  //testcase 1
		{patterns1, patterns2, true},  //testcase 2
		{patterns1, patterns3, false}, //testcase 3
		{patterns1, patterns4, false}, //testcase 4
		{patterns4, patterns1, false}, //testcase 5
	}

	for i, test := range tests {
		result := StringSliceEquals(test.Patterns1, test.Patterns2)
		if result != test.Expected {
			t.Error("StringSliceEquals failed on testcase#", i+1, "\n",
				"Expecting:", test.Expected, "Actual outcome:", result)
		}
	}
}
