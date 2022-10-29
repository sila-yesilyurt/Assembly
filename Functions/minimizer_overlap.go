package Functions

import "fmt"

func MakeOverlapNetworkMinimizers(reads []string, minimizerIndex StringIndex, match, mismatch, gap float64, threshold float64) map[string][]string {
	overlapGraph := make(map[string][]string)

	numReads := len(reads)
	mtx := InitializeSquareMatrix(numReads)

	//let's set all values equal to some horrible negative value

	bigNegative := -10000000.0
	for i := range mtx {
		for j := range mtx[i] {
			mtx[i][j] = bigNegative
		}
	}

	//range through the minimizer map and perform all appropriate alignments

	counter := 0
	for _, readIndices := range minimizerIndex {
		counter++
		if counter%100 == 0 {
			fmt.Println("Currently processing element", counter, "of the minimizer map.")
		}

		//perform all alignments that haven't been done
		for i := range readIndices {
			for j := i + 1; j < len(readIndices); j++ {
				//where are these reads in the read dataset?
				index1 := readIndices[i]
				index2 := readIndices[j]

				read1 := reads[index1]
				read2 := reads[index2]

				//have I performed an alignment?
				if mtx[index1][index2] == bigNegative {
					//we have not aligned and we should!
					mtx[index1][index2] = ScoreOverlapAlignment(read1, read2, match, mismatch, gap)
				}
				// the same goes for the symmetric value across the main diagonal
				if mtx[index2][index1] == bigNegative {
					mtx[index2][index1] = ScoreOverlapAlignment(read2, read1, match, mismatch, gap)
				}
			}
		}
	}

	//after we have built the matrix, we should binarize it
	binaryMatrix := BinarizeMatrix(mtx, threshold)

	//we know how to make the overlap graph -- this code we already wrote
	for i := range binaryMatrix {
		currRead := reads[i]
		//which reads do I connect reads[i] to? The ones with a 1 in binary mtx
		for j := range binaryMatrix[i] {
			if binaryMatrix[i][j] == 1 {
				overlapGraph[currRead] = append(overlapGraph[currRead], reads[j])
			}
		}
	}

	return overlapGraph
}
