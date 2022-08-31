package main

import (
		"fmt"
)

func add(x int, y int, ch chan int) {
	z := x + y
	ch <- z
}

func main() {

	ch := make(chan int, 5)
	for i := 0; i < 10; i++ {
		go add(i, i, ch)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}


