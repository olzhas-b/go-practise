package main

import (
	"fmt"
	"strings"
	"strconv"
	"sync"
)

func Swap(sli []int, i int) {
	x := sli[i]
	y := sli[i+1]
	sli[i] = y
	sli[i+1] = x
}

func BubbleSort(arr []int, wg *sync.WaitGroup) {
	var count_start int = 1
	var count_end int
	for count_start != count_end {
		count_start = count_end
		for idx, _ := range arr {
			if idx < len(arr) - 1 {
				if arr[idx] > arr[idx + 1] {
					Swap(arr, idx)
					count_end++
				}
			}
		}
	}
	wg.Done()
}

func SplitChunk(arr []int, c chan []int) {
    x := len(arr)
	y := x / 4
	c <- arr[0:y]
	c <- arr[y:2*y]
	c <- arr[2*y:3*y]
	c <- arr[3*y:x]
}

func main() {
	var input string
	chunks := make([]int, 0)
	fmt.Printf("Key in sequence of integers (separated by comma i.e. 3,2,1): ")
	fmt.Scan(&input)
	input_arr := strings.Split(input, ",")
	for _, val := range input_arr {
		val_int, _ := strconv.Atoi(val)
		chunks = append(chunks, val_int)
	}

	// Split chunks using channel to send and receive chunks
	c := make(chan []int)
	go SplitChunk(chunks, c)
	chunk1 := <- c
	chunk2 := <- c
	chunk3 := <- c
	chunk4 := <- c

	// Sort each chunk using wait group
	var wg sync.WaitGroup
	wg.Add(4)
	go BubbleSort(chunk1, &wg)
	go BubbleSort(chunk2, &wg)
	go BubbleSort(chunk3, &wg)
	go BubbleSort(chunk4, &wg)
	wg.Wait()
	fmt.Println("Intermediate chunks: ", chunk1, chunk2, chunk3, chunk4)

	// Combine and sort the whole chunk
	// Without running this code, it is 4 individually sorted chunks
	wg.Add(1)
	BubbleSort(chunks, &wg)
	wg.Wait()

	// Print sorted answer
	fmt.Println("Final chunk: ", chunks)
}
