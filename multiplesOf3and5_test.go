package euler

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input int
	want int
}



func TestMultiplesOf3and5(t *testing.T)  {
	TestCases := []TestCase{
		{
			input : 10,
			want : 23,
		},
		{
			input : 49,
			want : 543,
		},
		{
			input : 1000,
			want : 233168,
		},
		{
			input : 8456,
			want : 16687353,
		},
		{
			input : 19564,
			want : 89301183,
		},
	}

	fmt.Println()
	fmt.Printf("Testing Multiples 3 and 5...\n")
	fmt.Println()

	for i, test := range TestCases {
		have := MultiplesOf3and5(test.input)
		if test.want != have {
			t.Errorf("Failed")
		} else {
			fmt.Printf("\033[32m Test Case %v Success \033[0m\n", i)
		}
	}

	fmt.Println()

	
}