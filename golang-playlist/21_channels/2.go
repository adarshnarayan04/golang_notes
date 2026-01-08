package main

import (
	"fmt"
	"time"
)

//pass by reference ( as channels are reference types)
func processNum(numChan chan int) {
	// infinite loop to keep receiving data from channel
	for {
		num := <-numChan // receive value from channel ( this will block until value is available)
		//after this main go routine go to next line as the data is processed
		//also this go routine also become unblocked as it has received data
		fmt.Println("processing num", num)
	}
}
func main() {
	numChan := make(chan int)

	go processNum(numChan)
	
	numChan <- 10
	numChan <- 20
	numChan <- 30 // this will not be print as when num:= <- numChan ,(now it will not block) the main goroutine go next line and exit
	// so print num is never executted

	time.Sleep(time.Second*2) // added sleep to allow goroutine to process all numbers

}

