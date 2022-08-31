package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	fmt.Println(Books{"斗破苍穹","天蚕土豆","小说", 10001})
	fmt.Println(Books{title: "凡人修仙传", author: "位置", book_id: 10002})
	fmt.Println(Books{title: "凡人修仙传", author: "位置"})

	Book1 = Books{"斗破苍穹","天蚕土豆","小说", 10001}
	Book2 = Books{title: "凡人修仙传", author: "位置", book_id: 10002}

	fmt.Println(Book1)
	fmt.Println(Book2)

	Book2.author = "还是不知道"

	fmt.Println(Book2)

	
}