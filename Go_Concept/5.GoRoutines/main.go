// package main

// import "fmt"

// func main() {
//     go greeter("Hello")
//     greeter("world")
// }

// func greeter(s string) {
//     for i := 0; i < 6; i++ {
//        fmt.Println(s)
//     }
// }

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		greeter("Hello")
	}()

	greeter("world") // Runs synchronously, no WaitGroup needed

	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
