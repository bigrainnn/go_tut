package main

import (
	"fmt"
	"math/rand"
	"time"
)


// 主程序退出了，自己跳出来
func msgGen(service string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			case <- done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleaning done")
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}


//// 这里有坑
//func fanIn(chs ...chan string) chan string {
//	c := make(chan string)
//	for _, ch := range chs {
//		chCopy := ch   //!!!!!!!!!!
//		go func() {
//			for {
//				c <- <- chCopy
//			}
//		}()
//	}
//	return c
//}

//  优化 这里有坑
func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <- in
			}
		}(ch)
	}
	return c
}


//// 到底知不知道 chan 的数量
//func fanIn(c1, c2 chan string) chan string {
//	c := make(chan string)
//	go func() {
//		for {
//			c <- <-c1
//		}
//	}()
//	go func() {
//		for {
//			c <- <-c2
//		}
//	}()
//	return c
//}


//通过select实现FANIN
func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <- c1:
				c <- m
			case m := <- c2:
				c <- m
			}
		}
	}()
	return c
}


// 非阻塞
func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m,true
	default:
		return "", false
	}
}

func timeoutWalt(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m,true
	case <-time.After(timeout):
		return "", false
	}
}

//func main() {
//	m := msgGen("service1")
//	m1 := msgGen("service2")
//	// m2 := fanIn(m, m1)
//	m2 := fanInBySelect(m,m1)
//	// 这里是先等待第一个，再等待第二个
//	//for {
//	//	fmt.Println(<-m)
//	//	fmt.Println(<-m1)
//	//}
//	for {
//		fmt.Println(<-m2)
//	}
//	// 同时等待
//
//}

func main() {
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWalt(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}
	done <- struct{}{}
	<-done
}
