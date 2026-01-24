package main

import (
	"fmt"
)

func main() {

	// word := argErr{
	// 	42, "Work on now",
	// }

	fmt.Println(f(42))
}

type argErr struct {
	arg     int
	message string
}

var str string = "Some string"

func (a argErr) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.message)

}

func (a argErr) String() string {
	return fmt.Sprintf("%d - %s", a.arg, str)

}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, argErr{arg, "Can't work now"}
	}
	return arg + 3, nil
}
