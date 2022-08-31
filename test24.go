package main

import "fmt"

func main() {
	a := 1
	test(a)
	b := "hello"
	test(b)
	c := []int{1, 2, 3}
	test(c)
	var n interface{} = 55
	assert(n)
	var m interface{} = "hello"
	assert(m)
	searchType(n)
	searchType(m)
}

func test(i interface{}) {
	fmt.Printf("Type = %T, Value = %v\n", i, i)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func searchType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case []int:
		fmt.Println("[]int")
	default:
		fmt.Println("unknown")
	}
}