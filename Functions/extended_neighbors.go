package Functions

//GetExtendedNeighbors takes in a pattern (read), the overlap graph and maxK.
//It returns the extendedNeighbors list. For each neighbor *n* in this list,
//the distance between n and pattern is between 2 and maxK, inclusively.
func GetExtendedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	extendedNeighbors := []string{}

	currentNeighbors := adjList[pattern] // reachable in one step

	//what is reachable in between 2 and maxK steps?
	for j := 2; j <= maxK; j++ {
		currentNeighbors = AdjacentStrings(currentNeighbors, adjList) // one step away from the current neighbors
		currentNeighbors = RemovePattern(currentNeighbors, pattern)
		//throw out the starting pattern if for some reason we have come back to it in a "cycle"
		extendedNeighbors = append(extendedNeighbors, currentNeighbors...)
	}

	return extendedNeighbors
}

func RemovePattern(patterns []string, pattern string) []string {
	//does pattern occur in the list? if so, delete it

	for i := len(patterns) - 1; i >= 0; i-- {
		if patterns[i] == pattern {
			//remove patterns[i]
			patterns = append(patterns[:i], patterns[i+1:]...)
		}
	}

	return patterns
}

//AdjacentStrings finds all neighbors of a given collection of patterns
//for a given adjacency list.
func AdjacentStrings(currentNeighbors []string, adjList map[string][]string) []string {
	newNeighbors := make([]string, 0)

	for i := range currentNeighbors {
		//what is one away from the current string?
		stringsOneAway := adjList[currentNeighbors[i]]
		//add this to newNeighbors
		newNeighbors = append(newNeighbors, stringsOneAway...)
	}

	return newNeighbors
}
