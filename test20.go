package main

import "fmt"

func main() {
	L4:
		for i := 0; ; i++ {
			for j := 0; ; j++ {
				if i > 4 {
					break L4
				}
				if i >= 2 {
					continue L4
				}
				if j > 4 {
					break
				}
				fmt.Println(i, j)
			}
		}
}