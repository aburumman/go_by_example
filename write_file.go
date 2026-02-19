package main

import (
	"fmt"
	"log"
	"os"
	//"io"
)

func main(){
	write_file()

}

// func read_file(name){
// 	if file, err := os.Stat(name); err != nil {
// 		if os.IsNotExist(err)
// 	}
//}
func write_file(){
	fmt.Println("Enter the name of the file you want to create: ")
	var file_name string 
	fmt.Scanln(&file_name)
	if _, err := os.Stat(file_name); !os.IsNotExist(err){
		fmt.Println("The file already exist")
		return
	}
	
	fmt.Println("YOu want to create file %s", file_name)
	file, err := os.Create(file_name)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	
	msg := "The is the content"

	//n ,err := os.
	fmt.Fprintln(file, msg)
	x := os.WriteFile(file_name, []byte("This is a byte"), 0644)
	fmt.Println(x)
}