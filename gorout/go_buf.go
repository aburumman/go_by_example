package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
	"math"
	"sort"
)

func main(){
	
	//part1()
	//part2()
	//part3()
	part6()
	part4()
	part5()
}

func part4(){
	var val float64
	val = 456.8904
	fmt.Println(math.Round(val *100)/ 100) 
}

func part6(){
	states := make(map[string]string)

	states["LOS"] = "Lagos"
	states["OGN"] = "OGUN"
	states["OSN"] = "OSUN"
	states["TRB"] = "TARABA"
	states["ABV"] = "ABUJA"

	states_order := make([]string, len(states))
	i := 0

	for k := range states{
		states_order[i] = k
		i++
	}

	fmt.Println(states_order, states)
	fmt.Println("Sorted states")
	sort.Strings(states_order)
	fmt.Println(states_order, states)
	
}

func part5(){
	n := time.Now()

	ndate := time.Date(2025, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(n, ndate)
	dt := "Tuesday Nov 10 23:00:00 2026"
	ft, err := time.Parse(time.ANSIC, dt)

	check(err)

	fmt.Println(ft)
}


func part1(){

start := time.Now()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Po")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	fmt.Println(time.Now(), input, time.Since(start))
	// Enter a number
	fmt.Println("Enter a number")
	num := bufio.NewReader(os.Stdin)
	num_in, err := num.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	int_num, err := strconv.ParseInt(strings.TrimSpace(num_in), 10, 64)
	fmt.Println("Your number is: ", int_num)

}

func part2(){
	fmt.Println("Provide an iinput")
	var num int
	inp, err := fmt.Scanln(&num)

	check(err)

	fmt.Println(inp)


	}

func part3(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(input)
	sinput, err := strconv.ParseInt(input, 10, 64)
	fmt.Println(sinput, err)
}


func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}