package mathematics

import (
	"fmt"
	"testing"
)

type test struct {
	data     []float64
	expected float64
}

var tests = []test{
	{[]float64{1}, 1},
	{[]float64{1, 2}, 3},
	{[]float64{1, 2, 3}, 6},
}

var items = createItems(1000000)

func createItems(number int) []float64 {
	items := make([]float64, number)
	for i := 0; i < number; i++ {
		items[i] = float64(i)
	}

	return items
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

func ExampleSum1() {
	fmt.Println(Sum1([]float64{1, 2}))
	// Output: 3
}

func ExampleSum2() {
	fmt.Println(Sum2([]float64{1, 2}))
	// Output: 3
}

func BenchmarkSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum1(items)
	}
}

func BenchmarkSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum2(items)
	}
}
