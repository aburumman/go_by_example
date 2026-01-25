package main 

import (
	"fmt"
)

type Queue struct{
	slice []int
}

func (q *Queue) Enque(i int){
// Add to que from back
 q.slice = append(q.slice, i)
}

func (q *Queue) Deque() (int){
// Remove from que
// Remove first item
ret := q.slice[0]
q.slice = q.slice[1:]
//q.slice = q.slice[1:len(q.slice)]
return ret
}

func (q *Queue) String() (string){
	return fmt.Sprintf("%v", q.slice)
}

func main(){
	var myque *Queue
	myque = new(Queue)

	myque.Enque(123)
	fmt.Println("Adding to que: ", myque)
	myque.Enque(124)
	fmt.Println("Adding to que: ", myque)
	myque.Enque(12)
	fmt.Println("Adding to que: ", myque)
	myque.Enque(129)
	fmt.Println("Adding to que: ", myque)
	myque.Enque(123)
	fmt.Println("Adding to que: ", myque)
	myque.Enque(123)
	fmt.Println("Adding to que: ", myque)
	myque.Enque(124)
	fmt.Println("Adding to que: ", myque)
	fmt.Println("Removing from que: ",myque)
	myque.Deque()
	fmt.Println("Removing from que: ", myque)
	myque.Deque()
	fmt.Println("Removing from que: ",myque)
	myque.Deque()
	fmt.Println("Removing from que: ",myque)
}