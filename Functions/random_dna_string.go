package Functions

import "math/rand"

//GenerateRandomGenome takes a parameter length and returns
//a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(length int) string {
	//to prevent lots of string concatenations, we make a slice of symbols
	//that is our desired length
	symbols := make([]byte, length)
	for i := 0; i < length; i++ {
		//add a random symbol (A, C, G, or T) to the i-th value of the slice
		symbols[i] = RandomDNASymbol()
	}
	return string(symbols)
}

//RandomDNASymbol takes no inputs and returns A, C, G, or T with equal probability.
func RandomDNASymbol() byte {
	number := rand.Intn(4)

	switch number {
	case 0:
		return 'A'
	case 1:
		return 'C'
	case 2:
		return 'G'
	case 3:
		return 'T'
	}

	panic("Error: something really weird is happening")
}
