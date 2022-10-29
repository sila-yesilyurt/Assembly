package Functions

import (
	"math"
	"strings"
)

func MultAlignContig(contig []string) string {
	match := 3.0
	mismatch := 2.0
	gap := 2.0
	originalSuperGap := 12.0

	if len(contig) == 1 {
		return contig[0]
	}

	start1 := []string{contig[0]}
	start2 := []string{contig[1]}
	align := ProgressiveAlign(start1, start2, match, mismatch, gap, originalSuperGap)

	for i, read := range contig[2:] {
		match = gap * float64(len(align))
		supergap := originalSuperGap + float64(i*i)
		align = ProgressiveAlign(align, []string{read}, match, mismatch, gap, supergap)
		if (i+1)%100 == 0 {
			// every 100 steps, simplify the alignment
			align = []string{consensus(alignToCounts(align), align)}
		}
	}

	return consensus(alignToCounts(align), align)
}

func alignToCounts(alignment []string) [](map[rune]int) {
	l := len(alignment[0])
	for _, s := range alignment {
		if len(s) != l {
			panic("Alignments are different lengths")
		}
	}

	counts := make([](map[rune]int), 0)
	for i := range alignment[0] {
		count := make(map[rune]int)
		for _, s := range alignment {
			count[[]rune(s)[i]] += 1
		}
		counts = append(counts, count)
	}

	return counts
}

func consensus(counts [](map[rune]int), aln []string) string {
	if len(counts) == 0 {
		panic("no counts")
	}
	tot := 0
	for _, v := range counts[0] {
		tot += v
	}

	result := ""

	numEnabled := make([]int, len(counts))
	for _, s := range aln {
		startIdx := 0
		for i, c := range s {
			if c != '-' {
				startIdx = i
				break // found our starting index
			}
		}

		doneIdx := math.Max(float64(strings.LastIndex(s, "A")), float64(strings.LastIndex(s, "C")))
		doneIdx = math.Max(doneIdx, float64(strings.LastIndex(s, "G")))
		doneIdx = math.Max(doneIdx, float64(strings.LastIndex(s, "T")))
		doneIdx = math.Max(doneIdx, float64(strings.LastIndex(s, "N")))

		for i := int(startIdx); i <= int(doneIdx); i++ {
			numEnabled[i] += 1
		}
	}

	for i, count := range counts {
		var maxKey rune
		maxKey = 'A'
		for key, val := range count {
			if key == '-' {
				continue
			}
			if val >= count[maxKey] {
				maxKey = key
			}
		}

		if float64(count[maxKey])/float64(numEnabled[i]) >= 0.1 {
			result += string(maxKey)
		} else if float64(count['-'])/float64(tot+numEnabled[i]) >= 0.9 {
			result += "-"
		} else {
			result += "N"
		}
	}
	return result
}
