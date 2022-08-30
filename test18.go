package main

import "fmt"

func main() {
	for i := 0; ; i++ {
			if i == 2 {
				goto L1
			}
			fmt.Println(i)
	}
L1:
	fmt.Println("L1")
}