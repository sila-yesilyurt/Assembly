package Functions

//Minimizer takes a string text and an integer k as input.
//It returns the k-mer of text that is lexicographically minimum.
func Minimizer(text string, k int) string {
	if len(text) < k {
		panic("Error: k is longer than length of string in Minimizer()")
	}
	bestString := text[:k]
	for i := 0; i < len(text)-k+1; i++ {
		currentPattern := text[i : i+k]
		if currentPattern < bestString {
			bestString = currentPattern
		}
	}
	return bestString
}
