package main

import (
	"errors"
	"fmt"
	"log"
	"os"
    "io"
    "time"
    "bytes"
    "strings"
)

func main(){

   // PrintOne()
   file := "name.txt"
   withbuf()
   WriteTo(file)
   defer Readxfile(file)
   Wread()
}

func Wread(){
    m := strings.NewReader("Some reader")
    data, _ := io.ReadAll(m)
    fmt.Println(string(data))
}

func withbuf(){
    var buf bytes.Buffer
    buf.WriteString("Something New \n")
    io.Copy(os.Stdout, &buf)
    //buf.ReadString([]byte("Something new"))
}

func Readxfile(file string){
    data, err := os.ReadFile(file)
    check(err)
    fmt.Println(string(data))
}

func sum(){
    sum := 1

    for sum < 1000 {
        fmt.Println(sum)
        sum += sum
        if sum > 200{
            return
        }
    }
}

func WriteTo(name string) (string, error){
    file, err := os.Create(name)
    check(err)
    //defer close(file)
    file.WriteString("This is what i want \n")
    io.WriteString(file, "The next thing \n")

    defer file.Close()

    return "Done", nil
}

func check(err error){
    if err != nil{
        log.Fatalln(err)
    }
}

func PrintOne(){
    if len(os.Args) < 2{
        panic( errors.New("Please providean argument"))
    }
    inp := os.Args[1]

    byt := make([]byte, 64)

    fil, err := os.Open(inp)
    check(err)
    defer fil.Close()
    for {
    _, err := fil.Read(byt)
    
    if err == io.EOF {
        return
    }
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(byt))
    time.Sleep(1 * time.Second)
    
    }
}