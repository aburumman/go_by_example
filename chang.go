package main

import (
	"fmt"
)

func main() {}

func learn_chan() {
	mychan := make(chan string)

	go func() {
		mychan <- "Whats up"
	}()

	msg := <-mychan
	fmt.Println(msg)
}
