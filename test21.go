package main

import "fmt"

func main() {

	fmt.Printf("add(1, 2): %T\n", add(1, 2))
	fmt.Printf("sub(1, 2): %T\n", sub(1, 2))
	fmt.Printf("funcSum(1, 2, 3): %d\n", funcSum(1, 2, 3))
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("funcSum(a...): %d\n", funcSum(a...))
	fmt.Println(swap(1, 2))
	x,_ := swap(1, 2)
	fmt.Println(x)

	num := func(a, b int) int { return a + b }
	fmt.Printf("num(1, 2): %d\n", num(1, 2))
	fmt.Printf("funcSum2(num, 1, 2): %d\n", funcSum2(num, 1, 2))
	f := wrap("add")
	fmt.Printf("funcSum2(f, 1, 2): %d\n", funcSum2(f, 1, 2))
	fmt.Printf("f(1, 2): %d\n", f(1, 2))

}

func wrap(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int { return a + b }
	case "sub":
		return func(a, b int) int { return a - b }
	default:
		return nil
	}
}

func funcSum2(f func(int, int) int, x, y int) int {
	return f(x, y)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) (z int) {
	z = a - b
	return
}

func funcSum(args ...int) (z int) {
	for _, arg := range args {
		z += arg
	}
	return
}

func swap(a, b int) (int, int) {
	return b, a
}