package main 

import (
	"fmt"
)

func main(){
	readChan()
}

func readChan(){
	achan := make(chan int)

	go func(){
		for i := range 5 {
			achan <- i
		}
		defer close(achan)
	}()
	
	field, ok := <- achan
	if ok {
	fmt.Println(field)
	}
	sec, ok := <- achan
	if ok {
		fmt.Println(sec)
	}
}