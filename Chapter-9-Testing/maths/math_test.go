package maths

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]float64{1.0, 2.0})
	if v != 1.5 {
		t.Errorf("Expected: %f, got: %f", 1.5, v)
	}
}

func TestMax(t *testing.T) {
	v := Max([]float64{1.0, 2.0, 3.0, 4.0})
	if v != 4.0 {
		t.Errorf("Expected: %f, got: %f", 4.0, v)
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{1.0, 2.0, 3.0, 4.0})
	if v != 1.0 {
		t.Errorf("Expected: %f, got: %f", 1.0, v)
	}
}

type testPair struct {
	values  []float64
	average float64
}

var tests = []testPair{
	{[]float64{1.0, 2.0}, 1.5},
	{[]float64{1.0, 2.0, 3.0}, 2.0},
	{[]float64{1.0, 2.0, 3.0, 4.0}, 2.5},
	{[]float64{}, 0},
}

func TestAverage2(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Errorf("Expected: %f, got: %f", pair.average, v)
		}
	}
}
