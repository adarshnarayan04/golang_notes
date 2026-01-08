package main

import (
	"fmt"
	"time"
)

//emailChan <-chan string : receive only channel ( can only receive data from channel)
//done chan<- bool : send only channel ( can only send data to channel)
// test chan string : bidirectional channel ( can send and receive data)
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {

	emailChan := make(chan string, 100) //also act as queue like structure
	done := make(chan bool) //used of synchronization between goroutines

	go emailSender(emailChan, done)

	for i := 0; i < 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending.")

	// this is important ( to avoid deadlock )
	close(emailChan)
	<-done
	//as range loop in emailSender will exit when channel is closed
	//else it will keep looping forever waiting for more data
	// so defer will not execute and main goroutine will wait forever on <-done causing deadlock





}

