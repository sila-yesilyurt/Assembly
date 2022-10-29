package Functions

//MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
//It returns the adjacency list of the overlap network of the reads; edges are only included
//in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	adjacencyList := make(map[string][]string)

	//generate overlap scoring matrix and then binarize it
	mtx := OverlapScoringMatrix(reads, match, mismatch, gap)
	b := BinarizeMatrix(mtx, threshold)

	//range over our reads and set each read's corresponding set of reads that it overlaps with

	for i := range b {
		for j := range b[i] {
			//what are the reads that overlap well with reads[i]?
			if b[i][j] == 1 {
				//append reads[j] to the slice corresponding to reads[i]
				adjacencyList[reads[i]] = append(adjacencyList[reads[i]], reads[j])
			}
		}
	}

	return adjacencyList
}
