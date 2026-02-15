package main

import "fmt"

func fib(n int, cx chan int) {
	//mychan := make(chan int)
	x, y := 0, 1

	go func() {
		for i := 1; i <= n; i++ {
			cx <- x

			x, y = y, x+y
		}
		close(cx)
	}()

	for c := range cx {
		fmt.Println(c)
	}

}

func fib2 ( c, n chan int){
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x,y = y, x+y 
		case <- n:
			fmt.Println("Quit")
			return 
		}
	}
}

func main() {
	var n int = 8
	ch := make(chan int)
	go fib(n, ch)

	v := make(chan int)
q := make(chan int)

go func (){
	for x := 0; x < 10; x ++ {
		fmt.Println(<- v)
	}
	q <- 0
}()


fib2(v, q)
}
