package main

import "fmt"

func main() {
	//channels are used to communicate between goroutines
	// they are like pipes to send and receive data

	messageChan := make(chan string)

	messageChan <- "ping" // blocking (channel send operation)
	// it will wait here until another goroutine receives the data

	//in this case main goroutie is also recieving the data , but that code is after this line
	// so it will result in deadlock

	msg := <-messageChan //blocking (channel receive operation)
	//this is also blocking until data is available in channel

	fmt.Println(msg)
}

//this code will result in deadlock because main goroutine is trying to send data to channel
// but there is no other goroutine to receive that data
// so both goroutines are waiting for each other indefinitely
// to avoid that we can use another goroutine to send or receive data