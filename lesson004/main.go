package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(2)
		go GoRoutine(i)
	}

	wg.Wait()
}

func GoRoutine(num int) {
	defer wg.Done()
	fmt.Printf("Goroutine %d is working!\n", num)
}
