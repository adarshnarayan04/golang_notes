package main

import (
	"fmt"
	"sync"
)

//use pointer , so that we can modify the same WaitGroup instance
func task(id int, w *sync.WaitGroup) {
	//defer schedules a function call to run at the end of the current function, just before it returns
	defer w.Done() //w.done() decrements the WaitGroup counter by 1
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // wait for all goroutines to finish before exiting main
	// without this , main may exit before goroutines complete
	// when value of wg counter becomes 0 , then only Wait() unblocks
}
