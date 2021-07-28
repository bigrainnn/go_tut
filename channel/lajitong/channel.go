//package main_1
//
//import (
//	"fmt"
//	"sync"
//)
//
//func doworker(id int, c chan int, wg *sync.WaitGroup) {
//	//for {
//	//	// ok 表示是否还有值
//	//	n, ok := <-c
//	//	if !ok {
//	//		break
//	//	}
//	//	fmt.Printf("Worker %d received %d\n", id, n)
//	//}
//	// 或者
//	for n := range c {
//		fmt.Printf("Worker %d received %c\n", id, n)
//		wg.Done()
//	}
//}
//
////type worker struct {
////	in chan  int
////	done chan bool
////}
//
//type worker struct {
//	in chan  int
//	wg *sync.WaitGroup
//}
//
//
//func ctrateWorker(id int, wg *sync.WaitGroup) worker {
//	w := worker{
//		in: make(chan int),
//		wg: wg,
//	}
//	go doworker(id, w.in, wg)
//	return w
//}
//
////func chanDemo () {
////	var workers [10]worker
////	for i := 0; i < 10; i++ {
////		workers[i] = ctrateWorker(i)
////	}
////
////	for i := 0; i< 10; i++ {
////		workers[i].in <- 'a' + i
////	}
////
////	for _, worker := range workers {
////		<-worker.done
////	}
////
////	for i := 0; i< 10; i++ {
////		workers[i].in <- 'A' + i
////	}
////
////	for _, worker := range workers {
////		<-worker.done
////	}
////
////	// wait for all of them
////	//for _,worker := range workers {
////	//	<-worker.done
////	//	<-worker.done
////	//}
////	//time.Sleep(time.Millisecond)
////}
//
//
//func chanDemo_v2 () {
//	var wg sync.WaitGroup
//
//	var workers [10]worker
//	for i := 0; i < 10; i++ {
//		workers[i] = ctrateWorker(i, &wg)
//	}
//
//	for i := 0; i< 10; i++ {
//		workers[i].in <- 'a' + i
//		wg.Add(1)
//	}
//
//	for i := 0; i< 10; i++ {
//		workers[i].in <- 'A' + i
//		wg.Add(1)
//	}
//
//	wg.Wait()
//
//}
//
//func main()  {
//	chanDemo_v2()
//	// bufferedChannel()
//	// channelClose()
//}