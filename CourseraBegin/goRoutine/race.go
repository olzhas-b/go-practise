package main

import (
	"fmt"
	"time"
)

var firstFinished, secondFinished bool

func firstPlayer(x int) {
	for {
		x++
		fmt.Printf("first players distance %d\n", x)
		time.Sleep(13 * time.Millisecond)
		if x > 100 {
			if !secondFinished {
				fmt.Println("first Player Won")
			}
			firstFinished = true
			break
		}
	}
}

func secondPlayer(x int) {
	for {
		x++
		fmt.Printf("second players distance %d\n", x)
		time.Sleep(13 * time.Millisecond)
		if x > 100 {
			if !firstFinished {
				fmt.Println("second Player Won")
			}
			secondFinished = true
			break
		}
	}
}

func main() {
	// we run first and second player, if someone reach higher than 100, he win
	go firstPlayer(0)
	go secondPlayer(0)

	// for chance to see result in complier
	var input string
	fmt.Scan(&input)
}
