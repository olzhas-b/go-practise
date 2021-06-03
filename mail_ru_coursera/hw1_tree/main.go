package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func getSpace(x int, m *map[int]bool) string {
	var response string
	for i := 0; i < x; i++ {
		if (*m)[i] == true {
			response += "|"
		} else {
			response += " "
		}

	}
	return response
}

func recursion(fileName string, response *string, space int, m *map[int]bool) {
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatal(err)
	}
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	for key, file := range files {

		if len(file.Name()) == 0 {
			break
		}

		if key != len(files)-1 {
			*response += getSpace(space, m) + "├───" + file.Name() + "\n"
			(*m)[space] = true
			if file.IsDir() {
				recursion(fileName+"/"+file.Name(), response, space+4, m)
			}
		} else {
			*response += getSpace(space, m) + "└───" + file.Name() + "\n"
			(*m)[space] = false
			if file.IsDir() {
				recursion(fileName+"/"+file.Name(), response, space+4, m)
			}
		}
		fmt.Println(file.Name(), file.IsDir())
	}
}
func dirTree(out *os.File, path string, printFilse bool) error {
	var response string = ""
	m := make(map[int]bool)
	recursion(".", &response, 0, &m)
	_, err := (*out).WriteString(response)
	return err
}
func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
