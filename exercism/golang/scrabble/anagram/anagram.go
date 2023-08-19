package main

import (
	"fmt"
	"math/big"
)

func generateAnagrams(prefix, remaining string, results *[]string) {
	if remaining == "" {
		*results = append(*results, prefix)
		return
	}

	for i, char := range remaining {
		newPrefix := prefix + string(char)
		newRemaining := remaining[:i] + remaining[i+1:]
		generateAnagrams(newPrefix, newRemaining, results)
	}
}

func FindAnagrams(word string) []string {
	var results []string
	generateAnagrams("", word, &results)
	return results
}

func main() {
	word := "abcdefghijklmnopqrstuvwxyz"
	anagrams := FindAnagrams(word)
	numAnagrams := countAnagrams(word)
	fmt.Printf("Le nombre d'anagrammes possibles pour le mot %s est : %s\n", word, numAnagrams.String())
	fmt.Println(len(anagrams))
}

func factorial(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(int64(n))
	for i := n - 1; i >= 1; i-- {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func countAnagrams(word string) *big.Int {
	letterCount := make(map[rune]int)
	for _, char := range word {
		letterCount[char]++
	}

	totalLetters := len(word)
	numerator := factorial(totalLetters)
	denominator := big.NewInt(1)

	for _, count := range letterCount {
		denominator.Mul(denominator, factorial(count))
	}

	return new(big.Int).Div(numerator, denominator)
}
