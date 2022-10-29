package Functions

//ProgressiveAlignmentScoreTable takes two multiple alignments as well as a collection
//of Clustal scoring parameters. It returns a 2D matrix corresponding to
//the Clustal dynamic programming table for combining the two alignments heuristically
//into a single multiple alignment.
func ProgressiveAlignmentScoreTable(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) [][]float64 {

	// calculate interesting lengths and values
	var strLen1 = len(align1[0])
	var strLen2 = len(align2[0])
	if strLen1 == 0 || strLen2 == 0 {
		panic("Blah")
	}

	// ..cont.
	var numRows = strLen1 + 1
	var numCols = strLen2 + 1

	// initialize penalties
	var gapPenalty = gap           // ADJUSTABLE
	var mismatchPenalty = mismatch // ADJUSTABLE
	var matchBonus = match         // ADJUSTABLE

	// initialize scoring matrix
	scoreTable := make([][]float64, numRows)
	for i := range scoreTable {
		scoreTable[i] = make([]float64, numCols)
	}

	//penalize the 0th row and column as all gaps
	for j := 1; j < numCols; j++ {
		scoreTable[0][j] = float64(j) * (-supergap)
	}
	for i := 1; i < numRows; i++ {
		scoreTable[i][0] = 0
	}

	// traverse through manhattan graph, populate according to longest path recurrence
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {

			// calculate recurrence branches
			var upValue = scoreTable[i-1][j] - supergap   //indel
			var leftValue = scoreTable[i][j-1] - supergap //indel
			var diagValue = scoreTable[i-1][j-1] + SumPairsScore(align1, align2, i-1, j-1,
				matchBonus, mismatchPenalty, gapPenalty) // (mis)match

			// determine max, populate matrix
			scoreTable[i][j] = MaxFloat(upValue, leftValue, diagValue)
		}
	}

	return scoreTable

}
