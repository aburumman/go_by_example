package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//var t chan int
	t := make(chan int)
	go makeRand(t)
	//close(t)
	for val := range t {
		fmt.Printf("%d \n", val)
	}

}

func makeRand(v chan int) {
	for i := 0; i <= 10; i++ {
		v <- rand.Intn(1000)
		if len(v) >= 10 {
			return
		}
	}
	close(v)
}
