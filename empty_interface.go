package main

import (
	"fmt"
)

func desc(i interface{}){
	fmt.Printf("Type %T, \n Value: %v \n", i, i)
}



func main(){
	var some interface{}
	some = 456
	fmt.Println(some.(int))
	s := "stringsome"
	i := 456
	desc(s)
	desc(i)
	strc := struct {
		name string
	}{
		name: "string",
	}
	desc(strc)
	desc(some)
}