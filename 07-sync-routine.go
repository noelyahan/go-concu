package main

import (
	"fmt"
	"sync"
	"time"
)

// It's really important to close all go routines
// unless there is a special case to not close go routines (has to have a valid reason)
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("func 1", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("func 2", i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finished!")

}
