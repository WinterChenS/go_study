package main

import "fmt"

func main() {
	var array  = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := array[1:3]
	s2 := array[1:]
	s3 := array[:9]
	s4 := array[:]

	fmt.Printf("s1: %v\n", s1);
	fmt.Printf("s2: %v\n", s2);
	fmt.Printf("s3: %v\n", s3);
	fmt.Printf("s4: %v\n", s4);

	a := make([]int, 10)
	b := make([]int, 10, 15)

	fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("b: %v, len: %d, cap: %d\n", b, len(b), cap(b))

	var s []int
	fmt.Println(len(s) == 0, s == nil)
	s = nil
	fmt.Println(len(s) == 0, s == nil)
	s = []int(nil)
	fmt.Println(len(s) == 0, s == nil)
	s = []int{}
	fmt.Println(len(s) == 0, s == nil)

	s5 := append(s4, 1)
	fmt.Println(s5)
	s6 := append(s4, 12, 22, 32)
	fmt.Println(s6)
	s7 := append(s4, s6...)
	fmt.Println(s7)

	s8 := []int{1, 2, 3, 4, 5}
	s9 := []int{5, 4, 3}
	s10 := []int{6}
	copy(s8, s9)
	fmt.Printf("s8: %v\n", s8)
	copy(s10, s9)
	fmt.Printf("s10: %v\n", s10)

	s11 := []int{1, 2, 3, 4, 5}
	modify(s11)
	fmt.Printf("s11: %v\n", s11)
}

func modify(a []int)  {
	a[0] = 100
	fmt.Printf("a: %p, %v\n", &a, a)
}