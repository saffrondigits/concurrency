package main

import "fmt"

var counter int

func main() {

	for i := 0; i < 100; i++ {
		go CountInc()
	}

	fmt.Printf("counter: %v\n", counter)
}

func CountInc() {
	counter++
}
