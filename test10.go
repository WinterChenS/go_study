package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println(a[0])
	fmt.Println(len(a))
	var b [3]int = [3]int{1, 2, 3}
	var c [3]int = [3]int{1, 2}
	fmt.Println(b)
	fmt.Println(c[2])
	d := [...]int{1, 2, 3}
	fmt.Printf("%T\n", d)

	e := [5]int {1, 4: 4}
	f := [...]int {1, 5: 6}
	fmt.Println(e)
	fmt.Println(f)
	//多维数据
	var g [4][2]int
	h := [3][2]int{{1, 2}, {4, 5}, {7, 8}}
	i := [4][2]int{1:{4,5}, 3:{8,9}}
	j := [...][2]int{1:{4,5}, 3:{8,9}}
	fmt.Println(g, h, i, j)

	for n, v := range h {
		fmt.Println("打印打印")
		fmt.Println(n, v)
		for n2, v2 := range v {
			fmt.Println(n2, v2)
		}
	}

	a1 := [2]int{1, 2}
	a2 := [...]int{1, 2}
	a3 := [2]int{1, 3}
	fmt.Println(a1 == a2, a1 == a3, a2 == a3)

	x := [2]int{10, 20}
	y := x

	fmt.Printf("x: %p, %v\n", &x, x)
	fmt.Printf("y: %p, %v\n", &y, y)

	test(x)

}

func test(a [2]int) {
	fmt.Printf("a: %p, %v\n", &a, a)
}