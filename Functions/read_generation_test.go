package Functions

import "testing"

type testpairSimulateReadsClean struct {
	Genome string
	Len    int
	Prob   float64
}

/* HELPER */
//Checkes if all reads generated are of length k.
func IsCorrectLength(reads []string, k int) bool {
	for _, read := range reads {
		if len(read) != k {
			return false
		}
	}
	return true
}

//Returns a map of kmer: occurance (index) in the genome.
func MakeKMerSet(genome string, k int) map[string][]int {
	kmers := make(map[string][]int)
	for i := 0; i < len(genome)-k+1; i++ {
		kmer := genome[i : i+k]
		kmers[kmer] = append(kmers[kmer], i)
	}
	return kmers
}

//Checks if all reads are substrings of the genome
func IsSubString(reads []string, genome string, len int) bool {
	kmers := MakeKMerSet(genome, len)
	for _, read := range reads {
		_, found := kmers[read]
		if !found {
			return false
		}
	}
	return true
}

func TestSimulateReadsClean(t *testing.T) {
	tests := []testpairSimulateReadsClean{
		{"ATCCCTGTCCGAGGCTTAGGGGGACGGAGTCCGATCCAGGGGAG", 10, 0.5}, //small
		{GenerateRandomGenome(500), 150, 0.9},                     //medium
		{GenerateRandomGenome(10000), 500, 0.9},                   //large
	}

	for _, test := range tests {
		reads := SimulateReadsClean(test.Genome, test.Len, test.Prob)
		//Check if number of reads generated is approx. correct.
		ratio := float64(len(reads)) / float64(len(test.Genome)-test.Len+1)
		window := 0.15
		if ratio <= test.Prob-window || ratio >= test.Prob+window {
			t.Error("Error: # of reads generated is too far from desired probability.")
		}

		//Check if length of all reads is correct.
		result := IsCorrectLength(reads, test.Len)
		if !result {
			t.Error("Error: not all reads are of desired length.")
		}

		//Check if all reads are substrings of the genome.
		result = IsSubString(reads, test.Genome, test.Len)
		if !result {
			t.Error("Error: some reads is not a substring of the original genome.")
		}
	}
}
