package main

import "fmt"

var f func(string) int


func GenDisplaceFn(a, v, s float64) func(float64) float64 {

	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v * t) + s
	}
}
func main() {
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(2))
}
