package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Welcome to files in Golang")
	// content := "This need to go in txt"
	// file, err := os.Create("./myGoFile")
	// if err != nil {
	// 	panic(err)
	// }
	// length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(length)
	readFile()
	// defer file.Close()

}

func readFile() {
	path := "./Test.txt"
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content of file  is " + string(dat))
}
