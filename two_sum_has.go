package main

import (
	"fmt"
)

func main() {
	my_list := []int{34, 5, 6, 8, 9, 4, 3, 10}
	target := 10
	fmt.Println(two_sum(my_list, target))
}

func two_sum(mylist []int, target int) (int, int) {
	my_map := make(map[int]int, len(mylist))
	//index_a := 0

	for x, i := range mylist {
		//fmt.Println("x is %d, and i is %d", x, i)
		y := target - i
		//fmt.Println("y is: ", y)
		if _, ok := my_map[y]; ok {
			return my_map[y], x
		}
		my_map[i] = x
	}
	return 0, -1
}
