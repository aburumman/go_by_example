package main

import (
	"fmt"
)

func main() {
	xlist := []int{45, 6, 34, 12, 67, 89, 0, 10}
	fmt.Println(mysort(xlist))
}

func mysort(slist []int) ([]int, error) {
	//lim := len(slist) / 2
	y := 0
	for x := range len(slist) - 1 {
		for x := range len(slist) {
			if slist[x] > slist[x+1] {
				slist[x], slist[x+1] = slist[x+1], slist[x]
			}
		}
		y += x
	}
	y = 0

	return slist, nil
}
