package euler

import(
	"testing"
)



func TestMultiplesOf3and5(t *testing.T)  {
	type TestCase struct {
		input int
		want int
	}
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

	for _, test := range TestCases {
		have := MultiplesOf3and5(test.input)
		if test.want != have {
			t.Errorf("Failed")
		}
	}

	
}