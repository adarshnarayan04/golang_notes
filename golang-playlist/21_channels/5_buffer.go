package main

import (
	"fmt"
)

func main() {

	emailChan := make(chan string, 100)//buffered channel with capacity of 100
	// means we can send 100 values without blocking
	//if the buffer is full, then the sender will block until some space is available in the buffer

	//emailChan := make(chan string, 1) //it will cause deadlock as we are exceeding the buffer capacity

	emailChan <- "1@example.com" //send without blocking
	emailChan <- "2@example.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	//it will not cause deadlock as we are not exceeding the buffer capacity





}

