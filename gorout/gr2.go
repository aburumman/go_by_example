package main 

import ( 
	"time"
	"fmt")

func main(){
	f()
	go f()

	go func (ms string){
		for x := range ms{
			fmt.Println(x)
		}
	}("new")

	time.Sleep(time.Second)
}

func f(){
	for x := 0; x <= 3; x++ {
		fmt.Println(x)
	}
}