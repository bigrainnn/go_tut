package main

import (
	"fmt"
	"sync"
)

func doworker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan  int
	done func()
}


func ctrateWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doworker(id, w)
	return w
}

func chanDemo_v2 () {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = ctrateWorker(i, &wg)
	}

	//wg.Add(20)
	for i := 0; i< 10; i++ {
		//wg.Add(1)
		workers[i].in <- 'a' + i
		wg.Add(1)
	}

	for i := 0; i< 10; i++ {
		//wg.Add(1)
		workers[i].in <- 'A' + i
		wg.Add(1)
	}

	// wg.Add(1)  这个东西一定要放在前面。

	wg.Wait()

}

//


func main()  {
	chanDemo_v2()
	// bufferedChannel()
	// channelClose()
}