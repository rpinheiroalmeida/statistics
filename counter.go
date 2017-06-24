package statistics

import "math"

type Counter struct {
	counts map[float64]int
}

var counts map[float64]int

func NewCounter(sample []float64) *Counter {
	counts = make(map[float64]int, len(sample))

	for _, data := range sample {
		if _, ok := counts[data]; !ok {
			counts[data] = 1
		} else {
			counts[data]++
		}
	}
	return &Counter{counts}
}

func (counter Counter) MaxValue() int {
	maxValue := math.MinInt64
	for _, value := range counter.counts {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
