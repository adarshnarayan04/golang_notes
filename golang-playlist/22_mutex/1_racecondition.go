package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
}

func (p *post) inc() {
	//Each p.views++ is not atomic. It’s a read–modify–write sequence:
	// load views
	// add 1
	// store back

	// Two goroutines can overlap like this:

	// G1 loads 0
	// G2 loads 0
	// G1 stores 1
	// G2 stores 1

	//so causing race condition
	// to avoid that we can use Mutex

	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Go(func(){
			myPost.inc()
		})
		// go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.views)
}
