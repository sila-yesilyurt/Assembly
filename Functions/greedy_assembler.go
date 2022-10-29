package Functions

//GreedyAssembler takes a collection of strings and returns a genome whose
//k-mer composition is these strings. It makes the following assumptions.
//1. "Perfect coverage" -- every k-mer is detected
//2. No errors in reads
//3. Every read has equal length (k)
//4. DNA is single-stranded
func GreedyAssembler(reads []string) string {
	//create a copy of reads so that we don't delete originals
	reads2 := make([]string, len(reads))
	copy(reads2, reads)

	// greedy algorithm: look for whatever helps me the most (overlap of k-1 symbols).
	if len(reads) == 0 {
		panic("Error: No reads given to GenomeAssembler!")
	}
	// start with arbitrary read
	// first, what is k? length of first read
	k := len(reads2[0])

	genome := reads2[len(reads2)/2] // midpoint k-mer

	// let's throw out everything we have used
	reads2 = Remove(reads2, len(reads2)/2)

	// while we still have reads, try to extend current read
	for len(reads2) > 0 {
		// note: we need to remember to delete any kmer we use or else hit an infinite loop
		for i, kmer := range reads2 {
			// try to extend genome to left and right
			// a hit means that we match k-1 nucleotides to end of genome
			if genome[0:k-1] == kmer[1:] { // extending left
				// update genome by adding first symbol of kmer to left
				genome = kmer[0:1] + genome
				// throw out read
				reads2 = Remove(reads2, i)
				// stop the for loop so we don't have an index out of bounds error
				break // breaks innermost loop you are in
			} else if genome[len(genome)-k+1:len(genome)] == kmer[:k-1] { // extending right
				genome = genome + kmer[k-1:]
				reads2 = Remove(reads2, i)
				break
			}
		}
	}

	return genome
}

//Remove takes a collection of strings and an index.
//It removes the string at the given index and returns the updated array.
func Remove(patterns []string, index int) []string {
	// remember our trick for deleting an element
	patterns = append(patterns[:index], patterns[index+1:]...)
	return patterns
}
