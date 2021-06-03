package main

import (
	"fmt"
	"time"
)

type Animal struct {
	food       string
	voice      string
	locomotion string
}

func (a Animal) Speak() {
	fmt.Println(a.voice)
}
func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func main() {
	var m map[string]Animal = map[string]Animal{
		"cow":   Animal{food: "grass", voice: "moo", locomotion: "walk"},
		"bird":  Animal{food: "worms", voice: "peep", locomotion: "fly"},
		"snake": Animal{food: "mice", voice: "hsss", locomotion: "slither"},
	}
	for {
		var (
			animalType, action string
		)
		fmt.Print(">")
		fmt.Scan(&animalType)
		time.Sleep(1 * time.Second)
		fmt.Printf(">")
		fmt.Scan(&action)
		if action == "eat" {
			m[animalType].Eat()
		} else if action == "move" {
			m[animalType].Move()
		} else {
			m[animalType].Speak()
		}

	}
}
