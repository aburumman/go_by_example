package main

import (
	"fmt"
)

func main(){
	var a, b, res uint32
	fmt.Println("Enter 2 value separaated by space")
	fmt.Scanln(&a, &b)
	res = a + b
	fmt.Println(res)
}