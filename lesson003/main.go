package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(4)
	go SayHello("World")
	go SayHello("Dave")
	go SayHello("Kalidas")
	go SayHello("Ravi")

	wg.Wait()
}

func SayHello(name string) {
	defer wg.Done()
	fmt.Printf("Hello %s!\n", name)
}
