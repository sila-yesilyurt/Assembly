package Functions

import "math/rand"

//ShuffleStrings takes a collection of strings patterns as input.
//It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
	n := len(patterns)

	//we want to create a permutation of length n
	perm := rand.Perm(n) // creates a permutation of 0 to n-1 integers

	//create a new list of patterns that will represent the changed order
	patterns2 := make([]string, n)

	//set each element of patterns2 according to what the permutation tells us
	for i := range patterns {
		//what is perm[i]?
		index := perm[i]
		patterns2[i] = patterns[index]
	}

	return patterns2
}
