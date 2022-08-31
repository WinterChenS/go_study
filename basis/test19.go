package main

import "fmt"

func main() {
	L3:
		for i := 0; ; i++ {
			for j := 0; ; j++ {
				if i >= 2 {
					break L3
				}
				if j > 4 {
					break
				}
				fmt.Println(i, j)
			}
		}	
}