package Functions

import "sort"

//CopyGraph returns a copy of the input graph.
func CopyGraph(graph map[string][]string) map[string][]string {
	newGraph := make(map[string][]string)

	for key, value := range graph {
		newArr := make([]string, len(value))
		n := copy(newArr, value)
		if n != len(value) {
			panic("Something wrong happend when copying graph.")
		}
		newGraph[key] = newArr
	}
	return newGraph
}

//MtxEq takes two matrices and returns true if they are equal
//and false otherwise.
func MtxEq(mtx1, mtx2 [][]int) bool {
	if len(mtx1) != len(mtx2) {
		return false
	}
	for i := range mtx1 {
		if len(mtx1[i]) != len(mtx2[i]) {
			return false
		}
		for j := range mtx1[i] {
			if mtx1[i][j] != mtx2[i][j] {
				return false
			}
		}
	}
	return true
}

//GraphEq takes 2 graphs graph1 and graph2.
//It returns if graph1 and graph2 has the same keys and values.
func GraphEq(graph1, graph2 map[string][]string) bool {
	for key1, value1 := range graph1 {
		value2, found := graph2[key1]
		if !found {
			return false
		}
		if !StrSliceEq(value1, value2) {
			return false
		}
	}
	for key2, value2 := range graph2 {
		value1, found := graph1[key2]
		if !found {
			return false
		}
		if !StrSliceEq(value1, value2) {
			return false
		}
	}
	return true
}

//StrSliceEq takes 2 string slices slice1 and slice2.
//It returns if slice1 == slice2.
func StrSliceEq(slice1, slice2 []string) bool {
	sort.Strings(slice1)
	sort.Strings(slice2)
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

//SumLength takes a collection of strings and returns their total length.
func SumLength(patterns []string) int {
	c := 0

	for i := range patterns {
		c += len(patterns[i])
	}

	return c
}
