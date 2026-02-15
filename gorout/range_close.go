package main 

import (
	"fmt"
)

func fib(c chan int, val int){
	x, y := 0, 1
	for q := 1; q <= val; q++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main(){
	c := make(chan int, 6)

	fib(c, cap(c))
	//for x := range c{
		if v, ok := <- c ; ok{
		//fmt.Println(x)
		fmt.Println(v)
		}
	//}
}