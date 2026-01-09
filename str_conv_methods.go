package main

import (
	"fmt"
	"strconv"
)

func main(){
	myInt := "234"
	myBool := "true"
	i, err := strconv.Atoi(myInt)
	b, err := strconv.ParseFloat(myInt, 64)
	x, err := strconv.ParseBool(myBool)
	u, err := strconv.ParseUint(myInt, 10, 64)
	q, err := strconv.ParseInt(myInt, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i, b, x, u, q)
}