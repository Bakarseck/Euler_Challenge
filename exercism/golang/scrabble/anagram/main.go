package main

import (
	"strings"
)

func Detect(subject string, candidates []string) (result []string) {
	for _, word := range candidates {
		if HasSameCharacter(subject, word) {
			result = append(result, word)
		}
	}
	return
}

func HasSameCharacter(str, candidate string) bool {

	str, candidate = strings.ToUpper(str), strings.ToUpper(candidate)

	if str == candidate {
		return false
	}

	if len(str) != len(candidate) {
		return false
	}

	charCount := make(map[rune]int)
	for _, v := range str {
		charCount[v]++
	}

	for _, v := range candidate {
		if count, found := charCount[v]; found && count > 0 {
			charCount[v]--
		} else {
			return false
		}
	}

	return true
}
