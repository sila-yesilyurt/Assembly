package Functions

//AverageOutDegree takes the adjacency list of a directed network.
//It returns the average outdegree of each node in the network.
func AverageOutDegree(adjList map[string][]string) float64 {
	s := SumOutDegree(adjList)
	return float64(s) / float64(len(adjList))
}

//SumOutDegree takes an adjacency list and sums
//all the "outdegrees" of nodes in the list.
func SumOutDegree(adjList map[string][]string) int {
	sum := 0
	for _, list := range adjList {
		sum += len(list)
	}
	return sum
}
