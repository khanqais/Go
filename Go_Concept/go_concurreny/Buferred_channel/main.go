package main

import "fmt"

func main() {
	//unbuffered -> handshake
	// buffered channel -> capacity -> buffer size
	// make(chan int,4)
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	x := <-ch
	fmt.Println(x)
}
