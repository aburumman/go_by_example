package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	zm := []string{"<< Zum zum zum >>"}
	fmt.Println(lienlen(zm))
}

func lienlen(words []string) int{
	total := 0

	for _, t := range words{
		total += utf8.RuneCountInString(t)
	}

	numspace := len(words) -1

	return total+numspace
}



