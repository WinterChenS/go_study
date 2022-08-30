package main

import "fmt"

func main() {
	p := Person{
		name: "John",
	}
	fmt.Println(p.String())
	p.Modify()
	fmt.Println(p.String())
	p.ModifyP()
	fmt.Println(p.String())

	p1 := Point{1, 2}
	q1 := Point{3, 4}
	f := p1.Add
	fmt.Println(f(q1))

}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

type Point struct {
	x, y int
}

type Person struct {
	name string
}

func (p Person) String() string {
	return "Person name is " + p.name
}

func (p Person) Modify() {
	p.name = "Tom"
}

func (p *Person) ModifyP() {
	p.name = "Tom"
}
