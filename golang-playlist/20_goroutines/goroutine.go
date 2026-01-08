package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {

	for i := 0; i <= 10; i++ {
		go task(i)
		// it is asynchronous call, so the loop will not wait for task to complete
		//all will be scheduled by go runtime scheduler
		//on different cores if available
	}
	// now all the 11 goroutines are started ( they go to the scheduler )
	//but the main function is now complete , so it will exit
	// because of that all other goroutines will also stop immediately
	// to avoid that we can add sleep in main goroutine

	time.Sleep(time.Second*2) // if not , then output will be empty
}
