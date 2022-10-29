package Functions

import "math/rand"

//SimulateReadsClean takes a genome along with a read length and a probability.
//It returns a collection of strings resulting from simulating clean reads,
//where a given position is sampled with the given probability.
func SimulateReadsClean(genome string, readLength int, probability float64) []string {
	n := len(genome)
	reads := make([]string, 0)

	for i := 0; i < n-readLength+1; i++ {
		x := rand.Float64() //gives number in [0,1)
		if x < probability {
			//grab this read
			reads = append(reads, genome[i:i+readLength])
		}
	}

	return reads

}

//SimulateReadsMessy takes as input a genome, a read length, a probability, and error rates for substitutions, insertions, and deletions.
//It returns a collection of reads sampled from the genome where a given position is sampled with the given probability, and errors occur at the respective rates.
func SimulateReadsMessy(genome string, readLength int, probability float64, substitutionErrorRate, insertionErrorRate, deletionErrorRate float64) []string {
	n := len(genome)

	reads := make([]string, 0)

	//range over all k-mer starting positions of the genome
	for i := 0; i < n-readLength+1; i++ {
		x := rand.Float64()
		if x < probability {
			//sample a read from this location!
			currString := "" // this will represent my growing string

			currPosition := i

			for len(currString) < readLength && currPosition < len(genome) {
				// keep adding symbols until read is desired length
				// OR I hit the end of the genome

				//do we mutate the current symbol?
				y := rand.Float64()
				if y < substitutionErrorRate {
					// make a substitution
					sym := RandomDNASymbol()
					// it shouldn't be the same as our current symbol
					for sym == genome[currPosition] {
						//generate a new symbol
						sym = RandomDNASymbol()
					}

					//new symbol found, so append it to string
					currString += string(sym)
					currPosition++
				} else if y < substitutionErrorRate+insertionErrorRate {
					//insertion!
					//add a random symbol and don't move on (don't increment currPosition)
					currString += string(RandomDNASymbol())
				} else if y < substitutionErrorRate+insertionErrorRate+deletionErrorRate {
					//deletion!
					currPosition++ // advance one step
				} else {
					//normal case: take the symbol at the current position of genome
					currString += string(genome[currPosition])
					currPosition++
				}
			}

			// the read has been generated but we need to make sure it
			//didn't prematurely hit the end of the genome
			if len(currString) == readLength {
				reads = append(reads, currString)
			}
		}
	}

	return reads
}
