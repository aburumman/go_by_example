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
	xbool := false
	//some_int := int64(678)
	some_int := int64(678)
	xb := strconv.FormatBool(xbool)
	xy := strconv.FormatInt(some_int, 10)
	fmt.Println(i, b, x, u, q)
	fmt.Println(xb, xy)
}