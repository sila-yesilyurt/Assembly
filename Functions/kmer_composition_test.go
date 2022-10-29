package Functions

import "testing"

type testpair struct {
	Genome   string
	K        int
	Expected []string
}

func TestKmerComposition(t *testing.T) {
	genomes := []string{
		"AGCCGTCGAA",
		"ATTTTTTTTTTTGGGGGCTAAAGTGGGGGGGGGGGTTTTAAA",
		"GGTGCCAGTCGAGTGACTGGCCCCCAA",
	}

	answers := [][]string{
		{"AGCC", "GCCG", "CCGT", "CGTC", "GTCG", "TCGA", "CGAA"},
		{"AGCCGT", "GCCGTC", "CCGTCG", "CGTCGA", "GTCGAA"},
		{"ATTTTTTTTTTTGGGGGCTAAAGTGGGGGGGGGGGTTTTAA", "TTTTTTTTTTTGGGGGCTAAAGTGGGGGGGGGGGTTTTAAA"},
		{"GGTGCCAG", "GTGCCAGT", "TGCCAGTC", "GCCAGTCG", "CCAGTCGA", "CAGTCGAG",
			"AGTCGAGT", "GTCGAGTG", "TCGAGTGA", "CGAGTGAC", "GAGTGACT", "AGTGACTG",
			"GTGACTGG", "TGACTGGC", "GACTGGCC", "ACTGGCCC", "CTGGCCCC", "TGGCCCCC",
			"GGCCCCCA", "GCCCCCAA"},
	}

	tests := []testpair{
		{genomes[0], 4, answers[0]},
		{genomes[0], 6, answers[1]},
		{genomes[1], 41, answers[2]},
		{genomes[2], 8, answers[3]},
	}

	for i, test := range tests {
		result := KmerComposition(test.Genome, test.K)
		if !StrSliceEq(result, test.Expected) {
			t.Error("KmerComposition() failed testcase# ", i+1, "\n",
				"Failed on:", test.Genome, test.K, "\n",
				"Expecting", test.Expected, "Actual result:", result)
		}
	}
}
