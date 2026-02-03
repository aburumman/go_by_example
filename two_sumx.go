package main

import (
	"fmt"
)

func main(){


}
func sum2(nums []int, target int) []int{
	for i, left := range(nums){
		for j, right := range(nums){
			if left + right == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{0,0}
}