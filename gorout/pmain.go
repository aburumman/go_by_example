package main 

import (
	"fmt"
	"time"
)

func say(s string){
	for i := 0; i < 5; i ++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go say("Hello")
	say("World")

	usum()
}

func lisum(x []int, c chan int){
	v := 0
	for _, X := range x {
		v += X
	}

	c <- v
}

func usum(){
	c := make(chan int)
	some_list := []int{67, 89, -100, 56 }
	go lisum(some_list, c)
	go lisum(some_list[:len(some_list)/2], c)
	x, y := <- c, <-c
	fmt.Println(x, y, x+y)
}