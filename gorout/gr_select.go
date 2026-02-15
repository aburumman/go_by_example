package main 

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		time.Sleep(1 * time.Second)
		ch1 <- "Here comes the sun"
	}()

	go func(){
		time.Sleep(1 * time.Second)
		ch2 <- "Here goes the sun"
	}()

	for i := 0; i <= 2; i++{
		select{
		case msg := <- ch1:
			fmt.Println(msg)
		case msg2 := <- ch2:
			fmt.Println(msg2)
		// default:
			fmt.Println(nil)
		case <- time.After(time.Second * 3):
			fmt.Println("Timeout waiting for process")
		}
	}
}