package statistics

import (
	"reflect"
	"testing"

	"github.com/rpinheiroalmeida/linalg/vector"
)

func TestCount(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   []*Pair
	}{
		{vector.Vector{1.0, 2.0, 3.0},
			[]*Pair{&Pair{1.0, 1.0}, &Pair{2.0, 1.0}, &Pair{3.0, 1.0}},
		},
		{vector.Vector{1.0, 2.0, 1.0},
			[]*Pair{&Pair{1.0, 2.0}, &Pair{2.0, 1.0}},
		},
		{vector.Vector{1.0, 1.0},
			[]*Pair{&Pair{1.0, 2.0}},
		},
	}

	for _, c := range cases {
		gotCounter := Counter(c.sample)

		if !reflect.DeepEqual(gotCounter, c.want) {
			t.Errorf("Counter(%v) want: %v; got: %v",
				c.sample, c.want, gotCounter)
		}
	}
}

func TestCount_WhenSampleIsEmpty(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   []*Pair
	}{
		{vector.Vector{}, []*Pair{}},
	}

	for _, c := range cases {
		gotCount := Counter(c.sample)
		if len(gotCount) != len(c.want) && len(c.want) > 0 {
			t.Errorf("Counter(%v) is not empty.", c.sample)
		}
	}
}
