package main

import (
	"fmt"
	"time"
)

func race(x int) {
	for i := x; i < 10000000000; i += 8 {

	}
}
func main() {
	start := time.Now()
	for j := 0; j < 8; j++ {
		go race(j)
	}
	input := "1"
	fmt.Scan(&input)
	duration := time.Since(start)
	fmt.Println(duration)
}
