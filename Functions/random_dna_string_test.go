package Functions

import (
	"fmt"
	"testing"
)

func TestGenerateRandomGenome(t *testing.T) {
	var tests = []struct {
		n int
	}{
		{10000},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.n)
		t.Run(testname, func(t *testing.T) {
			ans := GenerateRandomGenome(tt.n)
			// check length
			if len(ans) != tt.n {
				t.Errorf("Error! Your function returned a string of length %d and should have returned length %d", len(ans), tt.n)
			}
			//check that symbols are valid
			for _, symbol := range ans {
				if symbol != 'A' && symbol != 'C' && symbol != 'G' && symbol != 'T' {
					t.Errorf("Error! Your string contains the symbol %q", symbol)
				}
				//check that we have correct proportion of symbols
				for _, symbol := range []byte{'A', 'C', 'G', 'T'} {
					if SymbolFreq(ans, symbol) < 0.24 || SymbolFreq(ans, symbol) > 0.26 {
						t.Errorf("Error! Frequency of %q is not approximately 0.25", symbol)
					}
				}
			}
		})
	}
}

func SymbolFreq(text string, symbol byte) float64 {
	n := len(text)
	c := 0
	for i := range text {
		if text[i] == symbol {
			c++
		}
	}
	return float64(c) / float64(n)
}
