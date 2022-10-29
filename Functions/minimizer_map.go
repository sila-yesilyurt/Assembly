package Functions

//StringIndex is a type that will map a minimizer string to its list of indices
//in a read set corresponding to reads with this minimizer.
type StringIndex map[string][]int

//MapToMinimizer takes a collection of reads, integer k and integer numIndices.
//It divides each read into numIndices pieces, and identifies the minimizer of each.
//It then returns a map of k-mers to the indices of reads having this k-mer
//as one of their minimizers.
func MapToMinimizer(reads []string, k int, numChunks int) StringIndex {
	//replace this with your code
	dict := make(StringIndex)

	//range over our reads and add each read's minimizers to the map
	for i, currString := range reads {
		// divide the string into numIndices equally sized chunks
		n := len(currString)
		p := (n - k + 1) / numChunks // # k-mers in each chunk
		for j := 0; j < numChunks; j++ {
			var s string
			if j == numChunks-1 {
				//add every symbol left to string
				s = currString[p*j:]
			} else {
				//normal case: I want to add p k-mers starting at position p*j
				s = currString[p*j : p*(j+1)+k-1]
			}

			// now I just find the minimizer of s
			m := Minimizer(s, k)

			// is the minimizer m in the database?
			if IsElementOf(dict[m], i) == false {
				// add the index of the read to the key corresponding to m
				dict[m] = append(dict[m], i)
			}
		}
	}
	return dict
}

func IsElementOf(a []int, j int) bool {
	for _, val := range a {
		if j == val {
			return true
		}
	}

	// j isn't in a if you make it here
	return false
}
