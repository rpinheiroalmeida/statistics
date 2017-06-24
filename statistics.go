package statistics

import (
	"fmt"
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

	return Sum(sample) / float64(sample.Len())
}

func Median(sample vector.Vector) float64 {
	check(sample)

	sort.Float64s(sample)

	half := sample.Len() / 2

	if oddSize(sample) {
		return sample[half]
	}

	return Mean(vector.Vector{sample[half-1], sample[half]})
}

func Quantile(sample vector.Vector, percentile float64) float64 {
	pIndex := int(percentile * float64(sample.Len()))

	sort.Float64s(sample)

	return sample[pIndex]
}

func Mode(sample vector.Vector) vector.Vector {
	check(sample)

	counter := NewCounter(sample)
	maxQuantity := counter.MaxValue()

	modes := vector.Vector{}

	for k, v := range counter.Items {
		if v == maxQuantity {
			modes = append(modes, k)
		}
	}
	sort.Reverse(modes)
	return modes
}

func DataRange(sample vector.Vector) float64 {
	return sample.Max() - sample.Min()
}

func DispersionMean(sample vector.Vector) vector.Vector {
	mean := Mean(sample)
	dispersion := vector.Vector{}

	for _, value := range sample {
		dispersion = append(dispersion, value-mean)
	}

	return dispersion
}

func Variance(sample vector.Vector) float64 {
	if sample.Len() <= 1 {
		panic(fmt.Errorf("The (%v) does not have the minimum size (%v)", sample, 2))
	}
	dispersionMean := DispersionMean(sample)

	return dispersionMean.SumOfSquares() / float64(sample.Len()-1)
}

func oddSize(sample vector.Vector) bool {
	return sample.Len()%2 == 1
}

func check(sample vector.Vector) {
	if sample.Empty() {
		panic("Operation Not allowed with empty sample")
	}
}
