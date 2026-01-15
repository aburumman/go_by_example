package main

import (
	"fmt"
)

func main() {
	//name := "This is the name"
	uber := []rune("Sir_King_Über")
	uberx := "Sir_King_Über"
	fmt.Println(len(uber))
	fmt.Println(len([]rune(uber)))
	fmt.Println([]rune(uber)[:5])
	fmt.Println(len([]rune(uber)))
	fmt.Println(len([]byte(uberx)))
	// for i := 0; i < len(name); i++ {
	// 	fmt.Print(string(name[i]), " ")
	// }

	// for i := 0; i < len(uber); i++ {
	// 	fmt.Print(string(uber[i]), " ")
	// }
	// bname := []byte(name)
	// for i := 0; i < len(bname); i++ {
	// 	fmt.Print(bname[i], " ")
	// }

	// logLevel := "デバッグ"

	// for index, val := range logLevel{
	// 	fmt.Println(index, val)
	// }
}
