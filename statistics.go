package statistics

import (
	"fmt"
	"math"
	"sort"

	"github.com/rpinheiroalmeida/collections"
	"github.com/rpinheiroalmeida/linalg/vector"
)

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
	counter := collections.NewCounter(sample)
	maxQuantity := counter.MaxValue()

	modes := make(vector.Vector, 0, maxQuantity)

	for k, v := range counter.Items {
		if v == maxQuantity {
			modes = append(modes, k)
		}
	}

	sort.Sort(sort.Reverse(modes))

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
	checkMinimumSize(sample.Len(), 1)
	dispersionMean := DispersionMean(sample)

	return dispersionMean.SumOfSquares() / float64(sample.Len()-1)
}

func StandardDeviation(sample vector.Vector) float64 {
	return math.Sqrt(Variance(sample))
}

func InterQuantileRange(sample vector.Vector) float64 {
	return Quantile(sample, 0.75) - Quantile(sample, 0.25)
}

func Covariance(x, y vector.Vector) float64 {
	n := x.Len()
	checkMinimumSize(n, 1)
	return (DispersionMean(x).Dot(DispersionMean(y))) / float64(n-1)
}

func checkMinimumSize(value, minimum int) {
	if value <= minimum {
		panic(fmt.Errorf("The minimum size was not obeyed - %d", minimum))
	}
}

func oddSize(sample vector.Vector) bool {
	return sample.Len()%2 == 1
}

func check(sample vector.Vector) {
	if sample.Empty() {
		panic("Operation Not allowed with empty sample")
	}
}
