package sum1

func Sum(items []float64) float64 {
	inputChan := createPipeline(items)

	c1 := sumChunk(inputChan)
	c2 := sumChunk(inputChan)

	c := fanIn(c1, c2)

	return <-c + <-c
}

func createPipeline(items []float64) <-chan float64 {
	ch := make(chan float64)

	go func() {
		for _, item := range items {
			ch <- item
		}

		close(ch)
	}()

	return ch
}

func sumChunk(in <-chan float64) <-chan float64 {
	ch := make(chan float64)

	go func() {
		var sum float64
		for item := range in {
			sum = sum + item
		}

		ch <- sum

		close(ch)
	}()

	return ch
}

func fanIn(input1, input2 <-chan float64) <-chan float64 {
	ch := make(chan float64)

	go func() {
		for {
			select {
			case sum, ok := <-input1:
				if ok {
					ch <- sum
				}
			case sum, ok := <-input2:
				if ok {
					ch <- sum
				}
			}
		}
	}()

	return ch
}
