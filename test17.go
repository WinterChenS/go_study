package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 4; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	a :=[...]int{10, 20, 30, 40, 50}
	for k := range a {
		fmt.Println(k)
	}

	for k, v := range a {
		fmt.Println(k, v)
	}

	s := []string{"a", "b", "c"}
	for k, v := range s {
		fmt.Println(k, v)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}