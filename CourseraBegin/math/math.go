package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
func Min(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}
