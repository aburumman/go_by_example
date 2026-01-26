package main

import (
	"fmt"
	"runtime"
)

func main() {
	some()
}

func some() {
	var mylist []uint16
	for i := 0; i < 100000; i++ {
		mylist = append(mylist, 100)
	}

	var mymem runtime.MemStats
	runtime.ReadMemStats(&mymem)

	fmt.Printf("TotalAlloc: %v MiB", mymem.TotalAlloc/1024/1024)
}
