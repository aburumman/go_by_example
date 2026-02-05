package main 

import (
	"fmt"
	"time"
)

func main(){
	buffered()
}

func buffered(){
	buf := make(chan string, 2)

	buf <- "Some string here"
	buf <- "Next string here inshaAllah"

	buffed := <- buf 
	//buffed = <- buf 
	time.Sleep(2000)
	// fmt.Println(<- buf)
	// fmt.Println(<- buf)
	fmt.Println(buffed)

}