package Functions

//ALL PENALTIES POSITIVE

//ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
//It returns the maximum score of an overlap alignment in which str0 is overlapped with str1.
//Assume we are overlapping a suffix of str0 with a prefix of str1.
func ScoreOverlapAlignment(str0, str1 string, match, mismatch, gap float64) float64 {
	scoreTable := OverlapScoreTable(str0, str1, match, mismatch, gap)
	numRows := len(scoreTable)
	numCols := len(scoreTable[numRows-1])

	return scoreTable[numRows-1][numCols-1] // lower right element of scoring table is the max.
}

//OverlapScoreTable takes two strings and alignment penalties. It returns a 2-D array
//holding dynamic programming scores for overlap alignment of str0 with str1 with these penalties.
//Assume we are overlapping a suffix of str0 with a prefix of str1.
func OverlapScoreTable(str0, str1 string, match, mismatch, gap float64) [][]float64 {
	if len(str0) == 0 || len(str1) == 0 {
		panic("Error: A string of zero length given to OverlapScoreTable.")
	}
	numRows := len(str0) + 1
	numCols := len(str1) + 1

	scoreTable := InitializeScoreTable(numRows, numCols)

	// we assume that str0 is coded along the rows of the matrix.
	// Thus, we want the 0-th column to all be zero to allow for "free rides" starting
	// at any position in the string. These values are zero by default.

	// next, set the values in the first row to negative infinite to only allow
	// prefix of str1 to match with suffix of str0.
	for j := 1; j < numCols; j++ {
		// scoreTable[0][j] = float64(math.MinInt64) assures no free ride.
		scoreTable[0][j] = scoreTable[0][j-1] - gap
	}

	// now I am ready to range row by row over the interior of table and apply the recurrence relation
	// for overlap alignment.
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//apply the recurrence relation
			upValue := scoreTable[i-1][j] - gap   //indel
			leftValue := scoreTable[i][j-1] - gap //indel
			var diagonalWeight float64
			if str0[i-1] == str1[j-1] { //match!
				diagonalWeight = match
			} else { // mismatch!
				diagonalWeight = -mismatch
			}
			diagValue := scoreTable[i-1][j-1] + diagonalWeight
			scoreTable[i][j] = MaxFloat(upValue, leftValue, diagValue)
		}
	}

	// finally, the lower right value isn't quite correct because it hasn't taken into account the free rides
	// from any element in the final row to the sink.  So we need to adjust it accordingly.

	// we can do so by setting it equal to the max element present in the final row.
	scoreTable[numRows-1][numCols-1] = MaxArrayFloat(scoreTable[numRows-1])

	return scoreTable
}

//InitializeScoreTable takes a number of rows and columns as input.
//It returns a matrix of zeroes as floats with appropriate dimensions.
func InitializeScoreTable(numRows, numCols int) [][]float64 {
	scoreTable := make([][]float64, numRows)
	for i := range scoreTable {
		scoreTable[i] = make([]float64, numCols)
	}
	return scoreTable
}

//MaxArrayFloat takes a slice of integers as input and returns the maximum value in the slice.
func MaxArrayFloat(a []float64) float64 {
	if len(a) == 0 {
		panic("Error: array given to MaxArray has zero length.")
	}
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

//MaxFloat is a variadic function that takes an arbitrary collection of floats.
//It returns the maximum one.
func MaxFloat(nums ...float64) float64 {
	m := 0.0
	// nums gets converted to an array
	for i, val := range nums {
		if val > m || i == 0 {
			// update m
			m = val
		}
	}
	return m
}
