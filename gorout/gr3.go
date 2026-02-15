package main 

import (
	"fmt"
)

func main (){
	msg := make(chan string)

	go func(){msg <- "Starting"}()

	message := <- msg
	fmt.Println(message)
	buf()
}

func buf(){
	msg := make(chan string, 2)

	msg <- "Starting"
	msg <- "Next"
	close(msg)

	for i := range msg{
		fmt.Println(i)
	}
}
