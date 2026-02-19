package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	//run work concurrenlty + collect result
	// pipe ->send values between goroutines
	// one send data-> ch<- value
	// another recieve value:= <-value
	type User struct {
		ID   int
		Name string
	}
	ch := make(chan User)
	//worker go routines
	go func() {
		//simulate small work
		time.Sleep(200 * time.Millisecond)
		//Send: block until main Recieve
		// unbuffred channel, send + receive is a handshake
		ch <- User{ID: 100, Name: "Qais"}

	}()
	fmt.Println("Waiting to reciever user..")
	// Main now blocks and waits untill data is realeased
	u := <-ch
	fmt.Println("Main:now go user", u, u.ID, u.Name)
}
