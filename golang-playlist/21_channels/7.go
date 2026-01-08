package main

import (
	"fmt"
)



func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	//Each iteration, select blocks until one channel has data
	//Select is used for Channels
	//Switch is used for Variables/Values (switch checks multiple values)

	for i := 0; i < 2; i++ {
		// in each iteration , one of the case will be selected randomly if multiple are ready
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}





}

