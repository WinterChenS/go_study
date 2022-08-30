package main

import "fmt"

func main() {
	
		var a int = 10
		var b int32 = 20

		fmt.Println(a + int(b))

		var c float32 = 10.5

		fmt.Println(int(c))

		fmt.Println(5 % 3)

		var d int = 100
		fmt.Println(a, d)
		a, d = d, a
		fmt.Println(a, d)

		e := 10.6
		fmt.Println(e)

		var x complex64 = 3 + 5i
		var y complex128 = complex(3.5, 10)

		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(real(x), imag(x))
		fmt.Println(real(y), imag(y))

		m := 1
		if m == 1 {
			fmt.Println("m is 1")
		}

		s := `Hello\r\n
		world`

		fmt.Println(s)

		h := "Hello\r\n world"

		fmt.Println(h)

		fmt.Println(len(h))

		fmt.Println(h[:4])

		fmt.Println(h[2:])

		s2 := "Hello 世界"

		for i := 0; i < len(s2); i++ {
			fmt.Println(i, s2[i])
		}

		for i, v := range s2 {
			fmt.Println(i, v)
		}

}	