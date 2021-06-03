package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func sortSlice(slice []int, wg *sync.WaitGroup) {
	sort.Ints(slice)
	wg.Done()
}
func mergeSlices(s1, s2, s3, s4, s []int) {
	var k1, k2, k3, k4, i int
	var n1, n2, n3, n4 int = len(s1), len(s2), len(s3), len(s4)
	for i < len(s) {
		var mn, mn_index int = 1<<31 - 1, -1
		if k1 < n1 && s1[k1] < mn {
			mn = s1[k1]
			mn_index = 1
		}
		if k2 < n2 && s2[k2] < mn {
			mn = s2[k2]
			mn_index = 2
		}
		if k3 < n3 && s3[k3] < mn {
			mn = s3[k3]
			mn_index = 3
		}
		if k4 < n4 && s4[k4] < mn {
			mn = s4[k4]
			mn_index = 4
		}
		switch mn_index {
		case 1:
			s[i] = s1[k1]
			k1++
		case 2:
			s[i] = s2[k2]
			k2++
		case 3:
			s[i] = s3[k3]
			k3++
		case 4:
			s[i] = s4[k4]
			k4++
		}
		i++
	}
}
func main() {
	start := time.Now()
	// Code to measure
	slice := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Ints(slice)
	for i := 0; i < 100000000; i++ {

	}
	duration := time.Since(start)
	fmt.Println(duration)
}
