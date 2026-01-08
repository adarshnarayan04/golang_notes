package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc() {
	defer func() {
		p.mu.Unlock()
		//so that when function exits , lock is released
	}()


	//use Lock only for critical section
	//as it create bottleneck , as other goroutines have to wait for lock to be released
	p.mu.Lock()
	p.views += 1

	//p.mu.Unlock()  ( can also be done here but if function panics before this line then lock will not be released)
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Go( func(){
			myPost.inc()
		})

	}

	wg.Wait()
	fmt.Println(myPost.views)
}
