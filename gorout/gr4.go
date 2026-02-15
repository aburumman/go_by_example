package main 

import (
	"fmt"
	//"time"
	"io"
	//"bytes"
	"strings"
)

func main(){

	letter := strings.NewReader("Some new words here")
	fmt.Println(letter.Len())
	//fmt.Print(countAlp(letter))
}

func countAlphabet(c io.Reader) (int, error){
	count := 0
	byt := make([]byte, 1024)

	for {
		n, err := c.Read(byt)
		for _, l := range byt[:n]{
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z'){
				count++
			}

		}
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, err
		}
	}
}
func countAlp(r io.Reader) (int, error){
	count := 0
	byt := make([]byte, 1024)

	

	for {
		n, err := r.Read(byt)
		for _, val := range byt[:n] {
			if (val >= 'a' && val <= 'z' || val >= 'A' && val <= 'Z'){
				count++
			}
		}
		if err == io.EOF {
			return count, nil
		}

		if err != nil {
			return 0, nil
		}
		
	}
	

	
}
