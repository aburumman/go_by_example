package main

import (
	"fmt"
	"runtime"
)

func main(){
	//var mlist []int
	var mlist []int8
	for i := 0; i <= 10000000; i++ {
		mlist = append(mlist, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotallAlloc (Heap) %v MiB \n", m.TotalAlloc/1024/1024)
}