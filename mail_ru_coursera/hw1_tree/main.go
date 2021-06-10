package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func getSpace(x int, m *map[int]int) string {
	var response string
	for i := 0; i < x; i++ {
		if (*m)[i] == 1 {
			response += "│"
		} else if (*m)[i] == 0 {
			response += "\t"
		}
	}
	return response
}

func recursion(fileName string, response *string, space int, m *map[int]int, printFile bool) {
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatal(err, "wtf")
	}
	myFiles := files[:]
	myFiles = myFiles[0:0]
	for _, file := range files {
		if file.IsDir() {
			myFiles = append(myFiles, file)
		} else if printFile {
			myFiles = append(myFiles, file)
		}
	}
	files = myFiles
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	for key, file := range files {
		if file.IsDir() == false && printFile == false {
			continue
		}
		if len(file.Name()) == 0 {
			break
		}
		if key != len(files)-1 {
			*response += getSpace(space, m) + "├───" + file.Name()
			(*m)[space] = 1
			if file.IsDir() == false {
				f, errOs := os.Open(fileName+"/"+file.Name())
				if errOs != nil {
					println(errOs)
				}
				stat, _ := f.Stat()
				sz := stat.Size()
				var suffix = " (empty)\n"
				if sz != 0 {
					suffix = fmt.Sprintf(" (%db)\n", sz)
				}
				*response += suffix
			}
			if file.IsDir() {
				*response += "\n"
				recursion(fileName+"/"+file.Name(), response, space + 2, m, printFile)
			}
		} else {
			*response += getSpace(space, m) + "└───" + file.Name()
			(*m)[space] = 2
			if file.IsDir() == false {
				if file.IsDir() == false {
					f, errOs := os.Open(fileName+"/"+file.Name())
					if errOs != nil {
						println(errOs)
					}
					stat, _ := f.Stat()
					sz := stat.Size()
					var suffix = " (empty)\n"
					if sz != 0 {
						suffix = fmt.Sprintf(" (%db)\n", sz)
					}
					*response += suffix
				}
			}
			if file.IsDir() {
				*response += "\n"
				recursion(fileName+"/"+file.Name(), response, space + 2, m, printFile)
			}
		}
	}
}
func dirTree(out *bytes.Buffer, path string, printFile bool) error {
	var response string = ""
	m := make(map[int]int)
	recursion(path, &response, 0, &m, printFile)
	_, err := (*out).Write([]byte(response))
	print(response)
	if err != nil {
		panic(err)
	}
	return nil
}
func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run showMeDir.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		log.Fatal(err, "wtf")
	}
}
