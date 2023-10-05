package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go CountInc()
	}

	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}

func CountInc() {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	counter++
}
