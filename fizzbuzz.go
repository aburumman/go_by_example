package main

import (
	"fmt"
)

/// fiizbuzz

func main() {
	fizzbuzz(50)
}

func fizzbuzz(x int) {
	for i := range x {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("buzz")
			continue
		}

		fmt.Println(i)
	}
}
