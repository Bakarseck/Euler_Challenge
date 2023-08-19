package main

import "strings"

var equivalence = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) (result string) {
	for _, v := range dna {
		result += string(equivalence[v])
	}
	return
}

// toRNA returns the RNA complement of the input DNA strand.
func ToRNA_(dna string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		case 'T':
			return 'A'
		case 'A':
			return 'U'
		}
		return r
	}, dna)
}
