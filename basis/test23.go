package main

import "fmt"

type Duck interface {
	Eat()
}

type Duck1 interface {
	Eat()
	Walk()
}

type Dog struct {

}

func (d Dog) Eat() {
	fmt.Println("Dog eat")
}

func (d Dog) Walk() {
	fmt.Println("Dog walk")
}

type Cat struct {
}

func (c Cat) Eat() {
	fmt.Println("Cat eat")
}

func main() {

	var d Duck = &Dog{}
	d.Eat()
	var c Duck = &Cat{}
	c.Eat()

	var e Duck = d
	e.Eat()

	s := []Duck{d, c}
	for _, v := range s {
		v.Eat()
	}

	var d1 Duck1 = &Dog{}

	d1.Walk()

	var d2 Duck = d1

	d2.Eat()

	// var d3 Duck1 = c
	// d3.Eat()
}