package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}

	ch := make(chan int)

	go sum(sl[:3], ch)
	go sum(sl[3:], ch)

	total1 := <-ch
	total2 := <-ch

	fmt.Printf("total1: %v\n", total1)
	fmt.Printf("total2: %v\n", total2)

	completeSum := total1 + total2
	fmt.Printf("completeSum: %v\n", completeSum)
}

func sum(arr []int, cha chan int) {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}

	cha <- sum
}
