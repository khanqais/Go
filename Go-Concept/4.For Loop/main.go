package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fruits := []string{"apple", "orange", "banana"}
	for index, _ := range fruits {
		fmt.Println(index)
	}
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
