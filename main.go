package main

import (
	"Assembly/Functions"
	"fmt"
)

func main() {
	fmt.Println("Genome assembly!")

	//RunGreedyAssembly()

	//SARSOverlapNetwork()

	//SARSOverlapNetworkMinimizers()

	//SARSOverlapNetworkMinimizersTrim()

	//SARSOverlapNetworkMinimizersTrimMessy()

	//SARSOverlapNetworkMinimizersTrimMessyContigs()

	//SARSOverlapNetworkMinimizersTrimMessyContigsConsensus()

	FinalCountdown()
}

func FinalCountdown() {
	//read in contigs as our reads

	fmt.Println("Read in contigs.")

	reads := ReadStringsFromFASTA("Output/contigs.fasta")

	fmt.Println("Contigs read. We have", len(reads), "reads.")

	fmt.Println("Make overlap network.")

	//now let's establish the alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//we also need a threshold for quality of overlap scores to retain
	//when we binarize
	threshold := 40.0

	adjList := Functions.MakeOverlapNetwork(reads, match, mismatch, gap, threshold)

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Let's trim the transitivity out of our network.")

	maxK := 3

	adjList = Functions.TrimNetwork(adjList, maxK)

	fmt.Println("Network trimmed.")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Now generating maximal non-branching paths.")

	contigs := Functions.GenerateContigs(adjList)

	fmt.Println("Non-branching paths generated!")

	fmt.Println("We have", len(contigs), "total contigs.")

	fmt.Println("Now we would like to align all of our contigs.")

	sequencedContigs := make([]string, len(contigs))

	for i := range contigs {
		fmt.Print("Aligning ", i+1, " contigs\r")
		sequencedContigs[i] = Functions.MultAlignContig(contigs[i])
	}

	fmt.Println()

	minLength := 300

	fmt.Println("Removing short contigs that are shorter than", minLength, "nucleotides.")

	sequencedContigs = Functions.RemoveShortContigs(sequencedContigs, minLength)

	fmt.Println("We are down to", len(sequencedContigs), "contigs.")

	fmt.Println("Writing final contigs to file.")

	outFilenameContigs := "Output/final_contigs.fasta"

	WriteContigsToFileFASTA(sequencedContigs, outFilenameContigs)

	fmt.Println("Pipeline complete!")
}

func RunGreedyAssembly() {
	genomeLength := 1000
	k := 10

	//generate a genome
	originalGenome := Functions.GenerateRandomGenome(genomeLength)

	//chop up into k-mers and shuffle

	reads := Functions.KmerComposition(originalGenome, k)

	reads = Functions.ShuffleStrings(reads)

	//now our algorithm won't know original order

	assembledGenome := Functions.GreedyAssembler(reads)

	//now, did we do good? chop up the assembled genome into k mers

	kmers := Functions.KmerComposition(assembledGenome, k)

	//do these kmers match our original kmers? If so, great!

	if Functions.StringSliceEquals(reads, kmers) == true {
		fmt.Println("Mission accomplished!")
	} else {
		fmt.Println("Something is wrong.")
	}
}

func SARSOverlapNetwork() {
	fmt.Println("Import the SARS-CoV genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV_genome.fasta")

	fmt.Println("Genome read. Let's simulate reads.")

	readLength := 150
	probability := 0.1

	reads := Functions.SimulateReadsClean(genome, readLength, probability)

	fmt.Println("Reads simulated! Now, generate the overlap network.")

	//now let's establish the alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//we also need a threshold for quality of overlap scores to retain
	//when we binarize
	threshold := 40.0

	adjList := Functions.MakeOverlapNetwork(reads, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))
}

func SARSOverlapNetworkMinimizers() {
	fmt.Println("Import the SARS-CoV genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV_genome.fasta")

	fmt.Println("Genome read. Let's simulate reads.")

	readLength := 150
	probability := 0.1

	reads := Functions.SimulateReadsClean(genome, readLength, probability)

	fmt.Println("Reads simulated! Now, generate the minimizer map.")

	k := 10
	numChunks := 4

	minimizerDict := Functions.MapToMinimizer(reads, k, numChunks)

	fmt.Println("Minimizer map is made. Map contains", len(minimizerDict), "total keys.")
	fmt.Println("Now let's make the overlap network.")

	//now let's establish the alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//we also need a threshold for quality of overlap scores to retain
	//when we binarize
	threshold := 40.0

	adjList := Functions.MakeOverlapNetworkMinimizers(reads, minimizerDict, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))
}

func SARSOverlapNetworkMinimizersTrim() {
	fmt.Println("Import the SARS-CoV genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV_genome.fasta")

	fmt.Println("Genome read. Let's simulate reads.")

	readLength := 150
	probability := 0.1

	reads := Functions.SimulateReadsClean(genome, readLength, probability)

	fmt.Println("Reads simulated! Now, generate the minimizer map.")

	k := 10
	numChunks := 4

	minimizerDict := Functions.MapToMinimizer(reads, k, numChunks)

	fmt.Println("Minimizer map is made. Map contains", len(minimizerDict), "total keys.")
	fmt.Println("Now let's make the overlap network.")

	//now let's establish the alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//we also need a threshold for quality of overlap scores to retain
	//when we binarize
	threshold := 40.0

	adjList := Functions.MakeOverlapNetworkMinimizers(reads, minimizerDict, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Let's trim the transitivity out of our network.")

	maxK := 3

	adjList = Functions.TrimNetwork(adjList, maxK)

	fmt.Println("Network trimmed.")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))
}

func SARSOverlapNetworkMinimizersTrimMessy() {
	fmt.Println("Import the SARS-CoV genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV_genome.fasta")

	fmt.Println("Genome read. Let's simulate (messy) reads.")

	readLength := 150
	probability := 0.1

	substitutionErrorRate := 0.003
	insertionErrorRate := 0.003
	deletionErrorRate := 0.003

	reads := Functions.SimulateReadsMessy(genome, readLength, probability, substitutionErrorRate, insertionErrorRate, deletionErrorRate)

	fmt.Println("Reads simulated! Now, generate the minimizer map.")

	k := 10
	numChunks := 4

	minimizerDict := Functions.MapToMinimizer(reads, k, numChunks)

	fmt.Println("Minimizer map is made. Map contains", len(minimizerDict), "total keys.")
	fmt.Println("Now let's make the overlap network.")

	//now let's establish the alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//we also need a threshold for quality of overlap scores to retain
	//when we binarize
	threshold := 40.0

	adjList := Functions.MakeOverlapNetworkMinimizers(reads, minimizerDict, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Let's trim the transitivity out of our network.")

	maxK := 3

	adjList = Functions.TrimNetwork(adjList, maxK)

	fmt.Println("Network trimmed.")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))
}

func SARSOverlapNetworkMinimizersTrimMessyContigs() {
	fmt.Println("Import the SARS-CoV genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV_genome.fasta")

	fmt.Println("Genome read. Let's simulate (messy) reads.")

	readLength := 150
	probability := 0.1

	substitutionErrorRate := 0.003
	insertionErrorRate := 0.003
	deletionErrorRate := 0.003

	reads := Functions.SimulateReadsMessy(genome, readLength, probability, substitutionErrorRate, insertionErrorRate, deletionErrorRate)

	fmt.Println("Reads simulated! Now, generate the minimizer map.")

	k := 10
	numChunks := 4

	minimizerDict := Functions.MapToMinimizer(reads, k, numChunks)

	fmt.Println("Minimizer map is made. Map contains", len(minimizerDict), "total keys.")
	fmt.Println("Now let's make the overlap network.")

	//now let's establish the alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//we also need a threshold for quality of overlap scores to retain
	//when we binarize
	threshold := 40.0

	adjList := Functions.MakeOverlapNetworkMinimizers(reads, minimizerDict, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Let's trim the transitivity out of our network.")

	maxK := 3

	adjList = Functions.TrimNetwork(adjList, maxK)

	fmt.Println("Network trimmed.")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Now generating maximal non-branching paths.")

	contigs := Functions.GenerateContigs(adjList)

	fmt.Println("Non-branching paths generated!")

	fmt.Println("We have", len(contigs), "total contigs.")
}

func SARSOverlapNetworkMinimizersTrimMessyContigsConsensus() {
	fmt.Println("Import the SARS-CoV-2 genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV-2_genome.fasta")

	fmt.Println("Genome read. Let's simulate (messy) reads.")

	readLength := 150
	probability := 0.1

	substitutionErrorRate := 0.003
	insertionErrorRate := 0.003
	deletionErrorRate := 0.003

	reads := Functions.SimulateReadsMessy(genome, readLength, probability, substitutionErrorRate, insertionErrorRate, deletionErrorRate)

	fmt.Println("Reads simulated! Now, generate the minimizer map.")

	k := 10
	numChunks := 6

	minimizerDict := Functions.MapToMinimizer(reads, k, numChunks)

	fmt.Println("Minimizer map is made. Map contains", len(minimizerDict), "total keys.")
	fmt.Println("Now let's make the overlap network.")

	//now let's establish the alignment parameters
	match := 1.0
	mismatch := 1.0
	gap := 5.0

	//we also need a threshold for quality of overlap scores to retain
	//when we binarize
	threshold := 50.0

	adjList := Functions.MakeOverlapNetworkMinimizers(reads, minimizerDict, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Let's trim the transitivity out of our network.")

	maxK := 3

	adjList = Functions.TrimNetwork(adjList, maxK)

	fmt.Println("Network trimmed.")

	fmt.Println("The network has", len(reads), "total reads.")

	fmt.Println("The average outdegree is:", Functions.AverageOutDegree(adjList))

	fmt.Println("Now generating maximal non-branching paths.")

	contigs := Functions.GenerateContigs(adjList)

	fmt.Println("Non-branching paths generated!")

	fmt.Println("We have", len(contigs), "total contigs.")

	fmt.Println("Now we would like to align all of our contigs.")

	sequencedContigs := make([]string, len(contigs))

	for i := range contigs {
		fmt.Print("Aligning ", i+1, " contigs\r")
		sequencedContigs[i] = Functions.MultAlignContig(contigs[i])
	}

	fmt.Println()

	minLength := 300

	fmt.Println("Removing short contigs that are shorter than", minLength, "nucleotides.")

	sequencedContigs = Functions.RemoveShortContigs(sequencedContigs, minLength)

	fmt.Println("We are down to", len(sequencedContigs), "contigs.")

	fmt.Println("Writing final contigs to file.")

	outFilenameContigs := "Output/contigs.fasta"

	WriteContigsToFileFASTA(sequencedContigs, outFilenameContigs)

	fmt.Println("Pipeline complete!")
}
