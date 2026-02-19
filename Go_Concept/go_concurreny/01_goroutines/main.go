package main

import (
	"fmt"
	"time"
)

func main() {
	

	start := time.Now()
	go func() {
		time.Sleep(300 * time.Microsecond)
		fmt.Println("goroutines A: Finished simulated API at", time.Since(start))
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("goroutines A: Finished simulated API at", time.Since(start))
	}()
	//main ->no waiting
	fmt.Println("main: started two go routines at", time.Since(start))

	//small work-> any logic
	fmt.Println("main:doing step 1", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main:doing step 2", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main:doing step 3", time.Since(start))

	//time sleep time
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main: existing at", time.Since(start))
}
