package main

import (
	"fmt"
)

func main(){
	someChan()
}

func someChan(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		ch1 <- "Some data"
	}()

	go func(){
		ch2 <- "Another data"
	}()

	select {
	case msg1 := <- ch1:
		fmt.Println(msg1)
	case msg2 := <- ch2:
		fmt.Println(msg2)
	}
}