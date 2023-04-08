package euler

import (
	"fmt"
	"testing"
)

func TestPrime(t *testing.T) {
	TestCases := []TestCase{
		{
			input: 2,
			want:  2,
		},
		{
			input: 3,
			want:  3,
		},
		{
			input: 5,
			want:  5,
		},
		{
			input: 7,
			want:  7,
		},
		{
			input: 8,
			want:  2,
		},
		{
			input: 13195,
			want:  29,
		},
	}

	fmt.Println()
	fmt.Printf("Testing Largest Prime Factor...\n")
	fmt.Println()

	for i, test := range TestCases {
		have := LargestPrimeFactor(test.input)
		if test.want != have {
			t.Errorf("\n\033[31mFailed, input : %v\nwant : %v\nhave : %v \033[0m\n", test.input, test.want, have)
		} else {
			fmt.Printf("\033[32m Test Case %v Success\033[0m\n", i)
		}
	}

	fmt.Println()
}
