package main

import (
	"fmt"
	"sync"
)

const (
	maxNumber = 100
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	wg.Add(2)

	go consumer(ch)
	go producer(ch)

	wg.Wait()
}

func producer(ch chan int) {
	for i := 1; i <= maxNumber; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func consumer(ch chan int) {
	for num := range ch {
		fmt.Println("Received: ", num)
	}
	wg.Done()
}
