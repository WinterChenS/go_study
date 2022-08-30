package main

import "fmt"

func main() {
	numbers := []int{0,1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	printSlice(numbers)

	fmt.Printf("numbers == %d\n", numbers)

	fmt.Println("numbers[1:4] ===", numbers[1:4])

	fmt.Println("numbers[:3] ===", numbers[:3])

	fmt.Println("numbers[4:] === ", numbers[4:])

	numbers1 := make([]int,0,5)

	printSlice(numbers1)

	numbers2 := numbers[:2]

	printSlice(numbers2)

	numbers3 := numbers[2:5]

	printSlice(numbers3)


	var numbers4 []int

	printSlice(numbers4)

	numbers4 = append(numbers4, 0)

	printSlice(numbers4)


	numbers4 = append(numbers4, 1)

	printSlice(numbers4)

	numbers4 = append(numbers4, 2,3,4)

	printSlice(numbers4)

	numbers5 := make([]int, len(numbers4), (cap(numbers4))*2)

	printSlice(numbers5)

	copy(numbers5, numbers4)

	printSlice(numbers5)


	
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}