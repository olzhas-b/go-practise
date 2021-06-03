package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
func TestMin(t *testing.T) {
	result := Min(1.0, 2.0)
	if result != 1.0 {
		t.Error("Expected 1.0 got", result)
	}
}
