package main

import "fmt"

func main() {
	x := [2]int{1, 2}
	modify(x)
	fmt.Printf("x: %p, %v\n", &x, x)
}

func modify(a [2]int) {
	a[0] = 100
	fmt.Printf("a: %p, %v\n", &a, a)
}