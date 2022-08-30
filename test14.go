package main

import "fmt"


type Person struct {
	name string
	age int
}

type admin struct {
	p Person
	isAdmin bool
}

type admin1 struct {
	Person
	isAdmin bool
}

type leader struct {
	p Person
	isLeader bool
}

func main() {
	u1 := Person{"小明", 18}
	fmt.Println(u1)	

	u2 := Person {
		name: "小红",
		age: 20,
	}
	fmt.Println(u2)
	fmt.Println(u2.name)
	fmt.Println(u2.age)

	u3 := admin{
		p: u1,
		isAdmin: true,
	}
	fmt.Println(u3)
	fmt.Println(u3.p.name)
	u4 := admin1{
		Person: u1,
		isAdmin: true,
	}

	fmt.Println(u4)
	fmt.Println(u4.name)
}
