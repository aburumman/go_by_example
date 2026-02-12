package main

import (
	"flag"
	"os"
	"time"
	"fmt"
)

var somevar string
var wait time.Duration
var debug bool

func main(){
	name := flag.String("name", "", "The name to say hello")
	flag.Parse()

	if *name == "" {
		fmt.Println("Add name to use this package")
		flag.Usage()
		os.Exit(1)
	}
	time.Sleep(wait)
	fmt.Println("Carry on")
	flag.StringVar(&somevar, "File_Name", "Choose a file name")
	flag.Parse()
}

func init(){
	flag.BoolVar(&debug, "debug", false, "Allows you to debug the programme")
	flag.StringVar(&somevar, "File_Name", "", "Choose a file name")
	defaultWait, err := time.ParseDuration("5s")
	if err != nil {
		panic("Coulndt parse duration")
	}
	flag.DurationVar(&wait, "wait-time", defaultWait, "Time to wait before starting")
	flag.Parse()
	//flag.Parse
}