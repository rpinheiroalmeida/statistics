package statistics

import (
	"reflect"
	"testing"

	"github.com/rpinheiroalmeida/linalg/vector"
)

func TestSum(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   float64
	}{
		{vector.Vector{7.0}, 7.0},
		{vector.Vector{32.0, 7.0}, 39.0},
		{vector.Vector{}, 0.0},
	}
	for _, c := range cases {
		gotSum := Sum(c.sample)
		if gotSum != c.want {
			t.Errorf("Expected total (%v) summing up (%v) but got (%v)", c.want, c.sample, gotSum)
		}
	}
}

func TestMean(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   float64
	}{
		{vector.Vector{7.0}, 7.0},
		{vector.Vector{13.0, 14.0}, 13.5},
	}
	for _, c := range cases {
		gotMean := Mean(c.sample)

		if gotMean != c.want {
			t.Errorf("Expected mean of (%v) for (%v) but got (%v)", c.want, c.sample, gotMean)
		}
	}
}

func TestMeanPanicsWhenEmptySlice(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected mean panic when empty sample")
		}
	}()

	Mean(vector.Vector{})
}

func TestMedian(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   float64
	}{
		{vector.Vector{7.0}, 7.0},
		{vector.Vector{8.0, 11.0}, 9.5},
		{vector.Vector{7.0, 8.0, 11.0}, 8.0},
		{vector.Vector{7.0, 9.0, 10.0, 17.0}, 9.5},
		{vector.Vector{7.0, 10.0, 17.0, 9.0}, 9.5},
	}
	for _, c := range cases {
		gotMedian := Median(c.sample)

		if gotMedian != c.want {
			t.Errorf("Expected median (%v) for (%v) but got (%v)", c.want, c.sample, gotMedian)
		}
	}
}

func TestMedianPanicsWhenEmptySlice(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected median panic when empty sample")
		}
	}()

	Median(vector.Vector{})
}

func TestQuantile(t *testing.T) {
	cases := []struct {
		sample     vector.Vector
		percentile float64
		want       float64
	}{
		{vector.Vector{7.0}, 0.99, 7.0},
		{vector.Vector{7.0, 9.0, 10.0, 13.0, 17.0}, 0.75, 13.0},
		{vector.Vector{7.0, 9.0, 13.0, 10.0, 17.0}, 0.75, 13.0},
	}

	for _, c := range cases {
		gotQuantile := Quantile(c.sample, c.percentile)
		if gotQuantile != c.want {
			t.Errorf("The expected quantile for (%v) with percentile of (%.2f) was (%.2f) but got (%.2f)", c.sample, c.percentile, c.want, gotQuantile)
		}
	}
}

func TestMode(t *testing.T) {
	cases := []struct {
		sample vector.Vector
		want   vector.Vector
	}{
		{
			vector.Vector{7.0},
			vector.Vector{7.0},
		},
		{
			vector.Vector{7.0, 13.0, 13.0},
			vector.Vector{13.0},
		},
		{
			vector.Vector{17.0, 7.0, 13.0, 17.0, 13.0},
			vector.Vector{17.0, 13.0},
		},
	}

	for _, c := range cases {
		gotMode := Mode(c.sample)

		if !reflect.DeepEqual(gotMode, c.want) {
			t.Errorf("Expected mode (%v) for (%v) but got (%v)", c.want, c.sample, gotMode)
		}
	}
}

func TestModeFailWhenEmptySample(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate mode for empty Sample")
		}
	}()

	Mode(vector.Vector{})
}
