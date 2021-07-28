package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//for {
	//	// ok 表示是否还有值
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %d\n", id, n)
	//}
	// 或者
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func ctrateWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo () {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = ctrateWorker(i)
	}

	for i := 0; i< 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

// 发数据就必须有人收
// 有缓存的channel
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main()  {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}