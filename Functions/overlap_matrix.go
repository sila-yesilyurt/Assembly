package Functions

//OverlapScoringMatrix takes a collection of reads along with alignment penalties.
//It returns a matrix in which mtx[i][j] is the overlap alignment score of
//reads[i] with reads[j].
func OverlapScoringMatrix(reads []string, match, mismatch, gap float64) [][]float64 {
	numReads := len(reads)
	mtx := InitializeSquareMatrix(numReads)

	for i := range mtx {
		for j := range mtx[i] {
			if i != j { // avoid loops
				mtx[i][j] = ScoreOverlapAlignment(reads[i], reads[j], match, mismatch, gap)
			}
		}
	}
	return mtx
}

//InitializeSquareMatrix takes an integer n and returns a 2-D slice of floats
//with default zero value.
func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n)
	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}
