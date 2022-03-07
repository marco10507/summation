package sum0

func Sum(items []float64) float64 {
	var sum float64

	for _, item := range items {
		sum += item
	}

	return sum
}
