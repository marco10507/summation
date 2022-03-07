package main

import (
	"fmt"
	"log"
	"summation/sums/sum0"
	"summation/sums/sum2"
	"time"
)

func main() {
	items := []float64{}

	for i := 0; i < 5*1000000; i++ {
		items = append(items, float64(i))
	}

	var b1 benchMark
	var b3 benchMark

	for i := 0; i < 100; i++ {
		func() {
			defer b1.timeTrack(time.Now(), "sum0", false)
			sum0.Sum(items)
		}()

		func() {
			defer b3.timeTrack(time.Now(), "sum2", false)
			sum2.Sum(items)
		}()
	}

	fmt.Printf("average sum0: %fms\n", b1.average()/1000)
	fmt.Printf("average sum2: %fms\n", b3.average()/1000)
	fmt.Printf("speed up of sum2: %fms\n", (b1.average() / b3.average()))
	fmt.Printf("delta: %fms\n", (b1.average()-b3.average())/1000)
}

type benchMark struct {
	results []time.Duration
	counter int
}

func (b *benchMark) timeTrack(start time.Time, name string, printTrack bool) {
	elapsed := time.Since(start)
	if printTrack {
		log.Printf("%s took %s", name, elapsed)
	}
	b.results = append(b.results, elapsed)
	b.counter = b.counter + 1
}

func (b benchMark) average() float64 {
	var sum float64

	for _, item := range b.results {
		sum += float64(item.Microseconds())
	}

	return sum / float64(b.counter)
}
