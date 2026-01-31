package main

import (
	"fmt"
)

type Person struct {
	Name string 
	Age int
}

type Describer interface{
	Describe()
}

func (p Person) Describe() {
	fmt.Println("Name: %v \n, Age: %v", p.Name, p.Age)
}

func getType(i interface{}){

switch x := i.(type){
case Describer:
	x.Describe()
default:
	fmt.Println("Unknown type")
}
}

func main(){
	var q Describer
	q = Person{
		Name: "Ayo",
		Age: 28,
	}
	fmt.Println(q)
}