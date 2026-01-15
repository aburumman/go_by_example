package main

import (
	"fmt"
	"strconv"
)

type Reshape interface {
	Area() string
	Perim() string
}

type Circle struct {
	Name     string
	Radius   int
	Diameter int
}

type Quad struct {
	Name    string
	Length  int
	Breadth int
}

func (r Quad) Area() string {
	return string(r.Length * r.Breadth)
}

func (r Quad) Perim() string {
	return string((r.Length * 2) + (r.Breadth * 2))
}

func (r Circle) Area() string {
	area := r.Radius * r.Radius
	str_area := strconv.Itoa(area)
	//return string(r.Radius * r.Radius)
	return str_area
}

func (r Circle) Perim() string {
	//return string(r.Radius * r.Radius)
	area := r.Radius * r.Radius
	str_area := strconv.Itoa(area)
	//return string(r.Radius * r.Radius)
	return str_area
}

func GetGeom(r Reshape) string {
	return fmt.Sprintf("The Area is: %v. \n The Perimeter is: %v", r.Area(), r.Perim())
}

func main() {
	myCircle := Circle{
		Name:     "Oval",
		Radius:   34,
		Diameter: 17,
	}
	fmt.Println(GetGeom(myCircle))
}
