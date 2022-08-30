package main

import (
	"fmt"
)

func main() {
	var code = 123
	var date = "2022-03-09"
	var url = "Code=%d&date=%s"
	var target_url = fmt.Sprintf(url, code, date)

	fmt.Println(target_url)
}