package Functions

//RemoveShortContigs takes a slice of strings as input corresponding to
//contigs as well as a threshold parameter.
//It throws out any contigs shorter than this parameter and returns
//the resulting list.
func RemoveShortContigs(contigs []string, minLength int) []string {
	//remember: things get weird when you remove items from a slice
	//if we start at 0

	for i := len(contigs) - 1; i >= 0; i-- {
		if len(contigs[i]) < minLength {
			//hack it out
			contigs = append(contigs[:i], contigs[i+1:]...)
		}
	}

	return contigs
}
