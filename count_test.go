package statistics

import (
	"reflect"
	"testing"

	"github.com/rpinheiroalmeida/linalg/vector"
)

func TestNewCounter(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   *Counter
	}{
		{vector.Vector{1.0, 2.0}, &Counter{map[float64]int{1.0: 1, 2.0: 1}}},
		{vector.Vector{1.0, 1.0}, &Counter{map[float64]int{1.0: 2}}},
	}

	for _, c := range cases {
		gotCounter := NewCounter(c.sample)
		if !reflect.DeepEqual(gotCounter, c.want) {
			t.Errorf("NewCounter(%v) want: %v but got: %v",
				c.sample, c.want, gotCounter)
		}
	}
}

func TestNewCounter_WhenSampleIsEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected NewCounter panic when empty sample")
		}
	}()

	NewCounter(vector.Vector{})
}

func TestMaxValues(t *testing.T) {
	cases := []struct {
		sample []float64
		want   int
	}{
		{[]float64{1.0, 2.0, 3.0, 1.0}, 2.0},
		{[]float64{1.0, 1.0}, 2.0},
		{[]float64{1.0, 2.0}, 1.0},
	}
	for _, c := range cases {
		gotMaxValue := NewCounter(c.sample).MaxValue()
		if c.want != gotMaxValue {
			t.Errorf("Counter(%v).MaxValue() want: %v but got %v.",
				c.sample, c.want, gotMaxValue)
		}
	}
}

func TestMaxValues_WhenCounterIsEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected Counter.MaxValue() panic when empty sample")
		}
	}()

	NewCounter(vector.Vector{}).MaxValue()
}

func TestValues(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   []int
	}{
		{vector.Vector{1.0, 2.0, 3.0, 4.0}, []int{1}},
		{vector.Vector{1.0, 2.0, 1.0, 4.0}, []int{1, 2}},
		{vector.Vector{1.0, 2.0, 1.0, 1.0}, []int{1, 3}},
		{vector.Vector{1.0, 1.0, 1.0, 1.0}, []int{4}},
	}

	for _, c := range cases {
		gotValues := NewCounter(c.sample).Values()
		if !reflect.DeepEqual(gotValues, c.want) {
			t.Errorf("Counter(%v).Values() want: %v but got %v",
				c.sample, c.want, gotValues)
		}
	}
}

func TestValues_WhenCounterIsEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected Counter.Values() panic when empty sample")
		}
	}()

	NewCounter(vector.Vector{}).Values()
}
