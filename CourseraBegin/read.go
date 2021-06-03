package main

import (
	"fmt"
	"os"
	"strings"
)

type FirstSecondName struct {
	fname string
	sname string
}

func main() {
	slice := make([]FirstSecondName, 0)
	file, _ := os.Open("text.txt")
	defer file.Close()

	stat, _ := file.Stat()
	bdata := make([]byte, stat.Size())
	file.Read(bdata)
	sdata := strings.Split(string(bdata), "\n")

	for _, value := range sdata {
		fsname := strings.Split(value, " ")
		slice = append(slice, FirstSecondName{fname: fsname[0], sname: fsname[1]})
	}

	for _, value := range slice {
		fmt.Printf("first name is: %s second name is: %s \n", value.fname, value.sname)
	}
}
