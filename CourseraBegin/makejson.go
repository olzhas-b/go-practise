package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Address string
}

func main() {
	p := new(Person)

	fmt.Print("please enter Name: ")
	fmt.Scanf("%s\n", &p.Name)

	fmt.Print("please enter Adress: ")
	fmt.Scanf("%s\n", &p.Address)

	barr, err := json.Marshal(p)
	if err != nil {
		panic("error")
	}
	fmt.Println(string(barr))
}
