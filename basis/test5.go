package main

import "fmt"

func main() {

	var a,b int = 1,100
	result := max(a,b)
	fmt.Printf("转换之前{a=%d,b=%d}\n",a,b)
	fmt.Printf("result: %d\n",result)
	
	
	swap(&a, &b)
	
	fmt.Printf("转换之后{a=%d,b=%d}\n",a,b)
	fmt.Printf("result: %d\n",result)

}


func max(a, b int) int {
	var result int
	if (a > b) {
		result = a
	} else {
		result = b
	}
	return result
}

func swap(a *int, b *int) {

	var temp int
	temp = *a
	*a = *b
	*b = temp
}