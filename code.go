package main

import (
	"fmt"
	"strconv"
)

func main(){

	do_some()
}

func do_some(){
	a := "900"
	b := int(90)
	c := strconv.Itoa(b)
	fmt.Printf("%T", c )
	fmt.Println("\n")
	fmt.Println(strconv.Atoi(a))
}