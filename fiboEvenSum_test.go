package euler

import (
	"fmt"
	"testing"
)

func TestFibo(t *testing.T) {
	TestCases := []TestCase{
		{
			input : 10,
			want : 10,
		},
		{
			input : 34,
			want : 44,
		},
		{
			input : 60,
			want : 44,
		},
		{
			input : 1000,
			want : 798,
		},
		{
			input : 100000,
			want : 60696,
		},
		{
			input : 4000000,
			want : 4613732,
		},
	}

	fmt.Println()
	fmt.Printf("Testing FIbonacciEvenSum...\n")
	fmt.Println()

	for i, test := range TestCases {
		have := FibonacciSum(test.input)
		if test.want != have {
			t.Errorf("\n\033[31m ❌Failed, input : %v\nwant : %v\nhave : %v \033[0m\n", test.input, test.want, have)
		} else {
			fmt.Printf("\033[32m ✅Test Case %v Success\033[0m\n", i)
		}
	}

	fmt.Println()
}