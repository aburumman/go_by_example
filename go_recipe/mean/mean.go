package main

import (
	"fmt"
	"time"
)

func main(){
	var list []float64 
	list = []float64{3, 5,6, 90, 34, 12, 6 }
	fmt.Println(mean(list))
	wer := []string{"you", "you", "are", "very", "good", "good", "very"}
	fmt.Println(word_freq(wer))
	fmt.Println()
	mytime()
	timeParse()
	chtz()
}

func chtz(){
	ch, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return 
	}
	ch_zone := time.Date(2020, time.February, 9, 19, 30, 0, 0, ch)
	fmt.Println("Chicago: ", ch_zone)


	ts, err := time.LoadLocation("America/NewYork")
	if err != nil {
		fmt.Println(err)
		return 
	} 
	t := time.Date(2023, time.April, 9, 18, 30, 0, 0, ts)
	fmt.Println("New York", t)

	nytime := ch_zone.In(ts)
	ch_time := t.In(ch)
	fmt.Println(ch_time)
	fmt.Println(nytime)
}

func check(err error){
	if err != nil {
		fmt.Println("Error occured : ", err)
		return
	}
}

func conv_time(ts, from, to string) (string, error){
	//from_tz := time.Now().Local().UTC().Location()
	to_tz, err := time.LoadLocation(to)
	check(err)
	from_tz, err := time.LoadLocation(from)
	check(err)
	tsi, err := time.Parse(ts, "08 June, 2020")
	check(err)
	from_date := time.Date(tsi.Year(), tsi.Month(), tsi.Day(), tsi.Hour(), tsi.Minute(), tsi.Second(), 0, from_tz)
	const format = "2006-01-02T15:04"
	from_datex, err := time.ParseInLocation(format, ts, from_tz)
	check(err)
	toTime := from_datex.In(to_tz)
	fmt.Println(toTime)
	conv_time := from_date.In(to_tz)
	return conv_time.String(), nil
}

func mean(list []float64) (mean_val float64){

	var x int
	var total float64
	for x < len(list) {
		total += list[x]
		x++;
	}

	mean_val = total / float64(len(list))

	return
}

func word_freq(list []string) map[string]int {
	//var word_map map[string]int
	word_map := make(map[string]int)

	for _, word:= range list {
		word_map[word]++
	}
	return word_map
}

func mytime(){
	ln := time.Date(1940, time.October, 9, 18, 30, 0, 0, time.UTC)
	fmt.Println(ln)

	fmt.Println(ln.Format("2006-01-02"))
	fmt.Println(ln.Format("Mon, Jan 02"))
	fmt.Println(ln.Format(time.RFC3339Nano))

	d := 3500 * time.Millisecond
	fmt.Println(d)
}

func timeParse(){
	ts := "June 18, 2022"
	t , err := time.Parse("January 02, 2006", ts)
	if err != nil {
		fmt.Println("Parse err: ", err)
	} else {
		fmt.Println(t)
	}

	dur := "1h"
	d, err := time.ParseDuration(dur)
	if err != nil {
		fmt.Println("Parse Duration err: ", err)
	} else {
		fmt.Println(d)
	}

}