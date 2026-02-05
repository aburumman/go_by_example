package main

import (
	"fmt"
)

//var s interface{}
var m int = 68

func main() {
	var s interface{} = 89
	fmt.Println(s)

	t, ok := s.(int)
	fmt.Println(t, ok)
	var x interface{} = 90
	getTytpe(x)
}


func getTytpe(s interface{}) {
	switch s.(type){
	case int:
		fmt.Println("We have an int")
	case string:
		fmt.Println("We have a string")
	case float64:
		fmt.Println("We have a float")
	default:
		fmt.Printf("%T", s)
	}
}