package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	name string
}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct {
	name string
}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct {
	name string
}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	var m map[string]Animal = make(map[string]Animal)
	for {
		var command, name, action, animalType string
		fmt.Print(">")
		fmt.Scan(&command, &name, &animalType)

		if command == "newanimal" {
			var animal Animal
			switch animalType {
			case "cow":
				animal = Cow{name}
				m[name] = animal
			case "snake":
				animal = Snake{name}
				m[name] = animal
			case "bird":
				animal = Bird{name}
				m[name] = animal
			}
			fmt.Println("Created it!")
		} else {
			action = animalType
			var animal Animal = m[name]

			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}

		}
	}
}
