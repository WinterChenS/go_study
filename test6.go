package main

import "fmt"
import "unsafe"

func main() {
	var array = [...]int{2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(unsafe.Sizeof(array))
	one := array[9]
	fmt.Println(one)

	var array2 [10] int

	for i := 0; i < 10; i++ {
		array2[i] = i + 100
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, array2[j])
	}
}