package statistics

import (
	"sort"

	"github.com/rpinheiroalmeida/linalg/vector"
)

type Pair struct {
	Key   float64
	Value int64
}

func (pair *Pair) increment() {
	pair.Value++
}

func Counter(sample vector.Vector) []*Pair {
	counts := map[float64]*Pair{}

	for _, data := range sample {
		_, ok := counts[data]
		if !ok {
			counts[data] = &(Pair{data, 1})
		} else {
			counts[data].increment()
		}
	}

	return values(counts)
}

func values(counts map[float64]*Pair) []*Pair {
	var keys []float64

	for k := range counts {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	var pairs []*Pair
	for _, key := range keys {
		pairs = append(pairs, counts[key])
	}

	return pairs
}
