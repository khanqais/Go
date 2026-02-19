package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("hii")
	var wg sync.WaitGroup
	//i will wait for 3 go routines
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("Task 1")
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Task 1 is done")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Task 2")
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Task 2 is done")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Task 3")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Task 3 is done")
	}()
	fmt.Println("main:waiting for all tasks to finished")
	wg.Wait()
	fmt.Println("all Task is finished")

}
