package main

import "fmt"

func main() {
	var m = map[string]int{"a": 1, "b": 2}
	fmt.Println(m)

	m1 := make(map[string]int)
	fmt.Println(m1)

	l := make([]int, 0, 10)
	fmt.Println(l)

	m2 := make(map[string]int, 10)
	fmt.Println(m2)
	fmt.Printf("m2: %p, %v\n", &m2, m2)

	var m3 map[string]int
	fmt.Println(m3 == nil, len(m3) == 0)
	m3 = make(map[string]int)
	m3["a"] = 1
	fmt.Println(m3)

	m["c"] = 3
	m["d"] = 4
	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["k"])

	delete(m, "a")
	delete(m, "k")
	fmt.Println(m)
	fmt.Println(len(m))

	if value,ok := m["d"];ok {
		fmt.Println("ok:")
		fmt.Println(value)
	}
	//遍历
	for k,v := range m {
		fmt.Println(k,v)
	}

	modify(m)
	fmt.Println(m)

}

func modify(m map[string]int) {
	m["a"] = 100
	fmt.Printf("m: %p, %v\n", &m, m)
}