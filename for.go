package main

import (
"fmt"

)

func main(){
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

