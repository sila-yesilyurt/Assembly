package Functions

import (
	"sort"
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := make([][]string, 2)
	tests[0] = []string{"APPLE", "CAR", "HELLO", "JAGUAR"}
	tests[1] = []string{"Hi", "Hi", "Hi", "Yeet", "Yo", "Yo"}

	for i := range tests {
		ans := ShuffleStrings(tests[i])
		sort.Strings(tests[i])
		sort.Strings(ans)
		if StrSliceEq(tests[i], ans) == false {
			t.Error("KmerComposition() failed testcase# ", i, "\n",
				"Failed on:", tests[i], "\n",
				"Actual result:", ans)
		}
	}
}
