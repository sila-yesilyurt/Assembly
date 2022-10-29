package Functions

type Alignment []string

//ProgressiveAlign takes two (multiple) alignments as input and
//returns a multiple alignment corresponding to combining the two
//alignments according to the Clustal dynamic programming heuristic.
func ProgressiveAlign(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) Alignment {

	scoreTable := ProgressiveAlignmentScoreTable(align1, align2, match, mismatch, gap, supergap)

	// now generate the backtracking matrix
	backtrack := GenerateBacktrackMatrix(align1, align2, scoreTable, match, mismatch, gap, supergap)

	return OutputAlignment(align1, align2, backtrack)
}

//OutputAlignment takes two multiple alignments and a backtrack matrix resulting
//from applying the Clustal heuristic. It returns the consolidated multiple alignment
//after backtracking.
func OutputAlignment(align1 Alignment, align2 Alignment, backtrack [][]string) Alignment {
	numStr1 := len(align1)
	numStr2 := len(align2)

	//first, initialize alignment

	newAlignment := make(Alignment, numStr1+numStr2)

	//make a 2D array of bytes to hold symbols of the alignment
	allRows := make([][]byte, numStr1+numStr2)
	for i := range allRows {
		allRows[i] = make([]byte, 0)
	}

	//start with row and col indices at the sink
	row := len(backtrack) - 1
	col := len(backtrack[0]) - 1

	//now, backtrack to the source

	for row != 0 || col != 0 {
		if backtrack[row][col] == "UP" {
			//take symbols from first alignment
			for i := 0; i < numStr1; i++ {
				allRows[i] = append([]byte{align1[i][row-1]}, allRows[i]...)
			}

			//add gap symbols for second alignment
			for j := 0; j < numStr2; j++ {
				allRows[numStr1+j] = append([]byte{'-'}, allRows[numStr1+j]...)
			}
			row--
		} else if backtrack[row][col] == "LEFT" {
			// gap symbols for first alignment
			for i := 0; i < numStr1; i++ {
				allRows[i] = append([]byte{'-'}, allRows[i]...)
			}

			//add symbols for second alignment
			for j := 0; j < numStr2; j++ {
				allRows[numStr1+j] = append([]byte{align2[j][col-1]}, allRows[numStr1+j]...)
			}

			col--
		} else if backtrack[row][col] == "DIAG" {
			//take symbols from first alignment
			for i := 0; i < numStr1; i++ {
				allRows[i] = append([]byte{align1[i][row-1]}, allRows[i]...)
			}

			//add symbols for second alignment
			for j := 0; j < numStr2; j++ {
				allRows[numStr1+j] = append([]byte{align2[j][col-1]}, allRows[numStr1+j]...)
			}

			row--
			col--
		} else {
			panic("Backtrack pointers not set correctly.")
		}
	}

	//at end, convert symbol arrays to strings
	for i := range newAlignment {
		newAlignment[i] = string(allRows[i])
	}

	return newAlignment
}

//GenerateBacktrackMatrix takes two multiple alignments as input as well as a score
//table corresponding to Clustal scoring. It returns a 2D array of
//backtrack pointers ("UP", "LEFT", and "DIAG") depending on whether we take
//symbols from the first alignment, the second alignment, or both, respectively.
func GenerateBacktrackMatrix(align1, align2 Alignment, scoreTable [][]float64, match, mismatch, gap, supergap float64) [][]string {

	//initialize our 2-D slice of strings to store all the pointers

	numRows := len(align1[0]) + 1
	numCols := len(align2[0]) + 1

	backtrack := make([][]string, numRows)
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	//set the backtracking pointers based on scores

	//0-th row and column are EZ
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}

	for i := 0; i < numRows; i++ {
		backtrack[i][0] = "UP"
	}

	backtrack[0][0] = "START" // not necessary but gives us a starting point

	// range over inside of table and set based on which edge was used
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			if scoreTable[i][j] == scoreTable[i][j-1]-supergap {
				backtrack[i][j] = "LEFT"
			} else if scoreTable[i][j] == scoreTable[i-1][j]-supergap {
				backtrack[i][j] = "UP"
			} else {
				//hopefully scoreTable[i][j] = sum of pairs score + scoretable[i-1][j-1]
				backtrack[i][j] = "DIAG"
			}
		}
	}

	// the following code is needed to backtrack from the appropriate place
	// in the bottom row. Could alternatively just take the max value and
	// backtrack from there.

	bestScore := -100000000.0
	bestScoreCol := numCols
	for i := numCols - 1; i >= 0; i-- {
		if scoreTable[numRows-1][i] >= bestScore {
			bestScore = scoreTable[numRows-1][i]
			bestScoreCol = i
		}
	}

	for col := bestScoreCol + 1; col < numCols; col++ {
		backtrack[numRows-1][col] = "LEFT"
	}

	return backtrack
}
