package main

import (
	"fmt"
	"sync"
)

func setup() {
	fmt.Println("Inits")
}
func hello(on *sync.Once) {
	on.Do(setup)
	fmt.Println("hello world")
}
func main() {
	var on sync.Once
	go hello(&on)
	go hello(&on)
	go hello(&on)
	var input string
	fmt.Scan(&input)
}
