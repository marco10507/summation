package sum2

import "runtime"

type chunck struct {
	start int
	end   int
}

func Sum(items []float64) float64 {
	chunks, _ := createChunks(len(items), runtime.NumCPU())

	inputChanel := createPipeline(chunks)
	sumChanel := make(chan float64)

	for i := 0; i < runtime.NumCPU(); i++ {
		sumChunk(inputChanel, sumChanel, items)
	}

	var totalSum float64

	for i := 0; i < len(chunks); i++ {
		totalSum += <-sumChanel
	}

	return totalSum
}

func sumChunk(input <-chan chunck, sumChanel chan<- float64, items []float64) {
	go func() {
		for chunck := range input {
			sum := sumRange(items, chunck.start, chunck.end+1)
			sumChanel <- sum
		}
	}()
}

func sumRange(items []float64, start int, end int) float64 {
	var sum float64
	for _, item := range items[start:end] {
		sum += item
	}

	return sum
}

func createPipeline(chunks []chunck) <-chan chunck {
	ch := make(chan chunck)

	go func() {
		for _, chunck := range chunks {
			ch <- chunck
		}

		close(ch)

	}()

	return ch
}

func createChunks(sliceLength int, numberOfChuncks int) ([]chunck, bool) {
	if sliceLength == 0 {
		return nil, false
	}

	var chuncks []chunck
	sliceLastIndex := sliceLength - 1
	chunckLength := (sliceLength / numberOfChuncks) + 1

	start := 0
	end := 0

	if chunckLength > sliceLastIndex {
		end = sliceLastIndex
	} else {
		end = chunckLength
	}

	firstChunck := chunck{
		start,
		end,
	}

	chuncks = append(chuncks, firstChunck)

	if firstChunck.start == firstChunck.end {
		return chuncks, true
	}

	for {
		start = end + 1
		end = start + chunckLength

		if start > sliceLastIndex {
			break
		}

		if end > sliceLastIndex {
			end = sliceLastIndex
		}

		newChunck := chunck{
			start,
			end,
		}

		chuncks = append(chuncks, newChunck)
	}

	return chuncks, true
}
