package mathematics

import "testing"

type test struct {
	data     []float64
	expected float64
}

var tests = []test{
	{[]float64{1}, 1},
	{[]float64{1, 2}, 3},
	{[]float64{1, 2, 3}, 6},
}

func TestSum1(t *testing.T) {
	for _, test := range tests {
		sum := Sum1(test.data)
		if sum != test.expected {
			t.Error("Expected", test.expected, "Actual", sum)
		}
	}
}

func TestSum2(t *testing.T) {
	for _, test := range tests {
		sum := Sum2(test.data)
		if sum != test.expected {
			t.Error("Expected", test.expected, "Actual", sum)
		}
	}
}
