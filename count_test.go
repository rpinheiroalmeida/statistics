package statistics

import "testing"

func TestMaxValues(t *testing.T) {
	cases := []struct {
		sample []float64
		want   int
	}{
		{[]float64{1.0, 2.0, 3.0, 1.0}, 2.0},
	}
	for _, c := range cases {
		gotMaxValue := NewCounter(c.sample).MaxValue()
		if c.want != gotMaxValue {
			t.Errorf("Counter(%v).MaxValue() want: %v but got %v.",
				c.sample, c.want, gotMaxValue)
		}
	}
}
