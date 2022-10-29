package Functions

//SumPairsScore takes two multiple alignments as well as two indices, and scoring
//parameters. It returns the sum of pairs score of the corresponding columns
//in the two alignments, using the specified scoring parameters.
func SumPairsScore(align1 Alignment, align2 Alignment,
	idx1 int, idx2 int, match float64, mismatch float64, gap float64) float64 {

	var totalScore = 0.0
	for _, str1 := range align1 {
		for _, str2 := range align2 {

			var char1 = str1[idx1]
			var char2 = str2[idx2]

			if char1 == '-' && char2 == '-' {
				totalScore += 0
			} else if char1 == '-' || char2 == '-' {
				totalScore -= gap
			} else if char1 == char2 {
				totalScore += match
			} else {
				totalScore -= mismatch
			}

		}
	}
	return totalScore

}
