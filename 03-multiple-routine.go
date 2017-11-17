package main

import "fmt"

func print() {
	for i := 0; ; i++ {
		fmt.Println("print:", i)
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		go print()
	}
	var input string
	fmt.Scanln(&input)
}
