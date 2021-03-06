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
	go test1()
	go test2()
	go test3()

	var input string
	fmt.Scanln(&input)
}

func loop(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
	}
}

func test1() {
	<-ch
	fmt.Println("Test 1 got the signal")
	loop("test1")
}

func test2() {
	<-ch
	fmt.Println("Test 2 got the signal")
	loop("test2")
}

func test3() {
	<-ch
	fmt.Println("Test 3 got the signal")
	loop("test3")
}
