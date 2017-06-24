package statistics

import "sort"

type Counter struct {
	Items map[float64]int
}

func NewCounter(sample []float64) *Counter {
	if len(sample) == 0 {
		panic("There is an empty sequence.")
	}
	counts := make(map[float64]int, len(sample))

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
	maxValue := 0
	for _, value := range counter.Items {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

type set map[int]bool

func (counter Counter) Values() []int {
	valuesSet := make(set)
	for _, value := range counter.Items {
		if _, ok := valuesSet[value]; !ok {
			valuesSet[value] = true
		}
	}

	values := make([]int, 0, len(valuesSet))
	for key := range valuesSet {
		values = append(values, key)
	}
	sort.Ints(values)
	return values
}
