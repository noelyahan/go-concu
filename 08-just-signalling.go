package main

import (
	"fmt"
	"time"
)

// This example shows the use case of
// channel signalling withod data via unbuffered channel

// create a channel with empty struct
var ch chan struct{}

func main() {
	ch = make(chan struct{})
	go func() {
		// lets add a delay to notify
		// delay represents the blocking
		// the processing time
		time.Sleep(time.Second * 3)
		close(ch)
	}()

	// all these methods will be notify when the channel is close
	test1()
	test2()
	test3()
}

func test1() {
	<-ch
	fmt.Println("Test 1 got the signal")
}

func test2() {
	<-ch
	fmt.Println("Test 2 got the signal")
}

func test3() {
	<-ch
	fmt.Println("Test 3 got the signal")
}
