package statistics

import (
	"sort"

	"github.com/rpinheiroalmeida/linalg/vector"
)

type Sample vector.Vector

func Sum(sample vector.Vector) float64 {
	total := 0.0
	for _, value := range sample {
		total += value
	}
	return total
}

func Mean(sample vector.Vector) float64 {
	check(sample)

	return Sum(sample) / float64(sample.Size())
}

func Median(sample vector.Vector) float64 {
	check(sample)

	sort.Float64s(sample)

	half := sample.Size() / 2

	if oddSize(sample) {
		return sample[half]
	}

	return Mean(vector.Vector{sample[half-1], sample[half]})
}

func Quantile(sample vector.Vector, percentile float64) float64 {
	pIndex := int(percentile * float64(sample.Size()))

	sort.Float64s(sample)

	return sample[pIndex]
}

func Mode(sample vector.Vector) vector.Vector {
	check(sample)

	counts := count(sample)

	maxQuantitie := maxValue(counts)

	modes := vector.Vector{}

	for k, v := range counts {
		if v == maxQuantitie {
			modes = append(modes, k)
		}
	}

	return modes
}

func DataRange(sample vector.Vector) float64 {
	return sample.Max() - sample.Min()
}

func count(sample vector.Vector) map[float64]int64 {
	counts := map[float64]int64{}

	for _, value := range sample {
		_, ok := counts[value]
		if !ok {
			counts[value] = 0
		}
		counts[value]++
	}

	return counts
}

func maxValue(counts map[float64]int64) int64 {
	var quantities vector.Vector = vector.Vector{}
	for _, v := range counts {
		quantities = append(quantities, float64(v))
	}

	return int64(quantities.Max())
}

func oddSize(sample vector.Vector) bool {
	return sample.Size()%2 == 1
}

func check(sample vector.Vector) {
	if sample.Empty() {
		panic("Operation Not allowed with empty sample")
	}
}
