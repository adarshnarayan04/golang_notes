package main

import (
	"fmt"
)


// goroutine synchronizer
func task(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("processing...")
}
func main() {

	done := make(chan bool)
	go task(done)

	<-done // block

	// this act as a wait group
	// as last <-done will block until some value is received from done channel
	// we receive value from done channel only when task function is complete ( as defer func() { done <- true }() is called at the end of task function )
	// so main goroutine will wait here until task goroutine is complete
	//so this is a way to synchronize goroutines ( like wait groups )

	//use this pattern when you have only one goroutine to wait for
	//else use sync.WaitGroup for multiple goroutines

}

