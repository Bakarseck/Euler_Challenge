package euler

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input int
	want  int
}

func TestMultiplesOf3and5(t *testing.T) {
	TestCases := []TestCase{
		{
			10,
			23,
		},
		{
			49,
			543,
		},
		{
			1000,
			233168,
		},
		{
			8456,
			16687353,
		},
		{
			19564,
			89301183,
		},
	}

	fmt.Println()
	fmt.Printf("Testing Multiples 3 and 5...\n")
	fmt.Println()

	for i, test := range TestCases {
		have := MultiplesOf3and5(test.input)
		if test.want != have {
			t.Errorf("\n\033[31m ❌Failed, input : %v\nwant : %v\nhave : %v \033[0m\n", test.input, test.want, have)
		} else {
			fmt.Printf("\033[32m ✅Test Case %v Success \033[0m\n", i)
		}
	}

	fmt.Println()

}
