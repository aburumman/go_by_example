package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter a text:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("You Entered: ", input)
		b := make([]byte, len(input))
		n := strings.NewReader(input)
		_, err = n.Read(b)
		if err != nil {
			fmt.Println("Issue occured: ", err)
		}
		fmt.Println("Using string reader, ", string(b))
}