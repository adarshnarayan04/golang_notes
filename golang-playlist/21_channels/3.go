package main

import (
	"fmt"
)


// receive ( as recieve from channel)
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult //blocking
}
func main() {

	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result // blocking

	fmt.Println(res)

}

