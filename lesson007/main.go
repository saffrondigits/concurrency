package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 7
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 9
	}()

	select {
	case result := <-ch:
		fmt.Println("Received 1: ", result)
	case result := <-ch2:
		fmt.Println("Received 2: ", result)
	case <-time.After(3 * time.Second):
		fmt.Println("Time out happened!")
	}
}
