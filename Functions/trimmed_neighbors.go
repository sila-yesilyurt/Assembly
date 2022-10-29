package Functions

//GetTrimmedNeighbors takes in a string pattern (read), an adjacency list and maxK.
//It returns all n-transitivity edges in the adjList of the current read (pattern) for n <= maxK.
func GetTrimmedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	neighbors := adjList[pattern]

	extendedNeighbors := GetExtendedNeighbors(pattern, adjList, maxK)

	//is there anything in neighbors that is also in extendedNeighbors?
	//If so, throw it out
	for i := len(neighbors) - 1; i >= 0; i-- {
		if Contains(extendedNeighbors, neighbors[i]) {
			//remove neighbors[i] from neighbors
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
		}
	}

	return neighbors
}

//Contains takes a list of strings and a single string as input.
//It returns true if the list contains the element, and false otherwise.
func Contains(patterns []string, element string) bool {
	for _, pattern := range patterns {
		if pattern == element {
			return true
		}
	}
	//default: survived all tests
	return false
}
