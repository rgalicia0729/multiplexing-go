package main

import (
	"fmt"
	"time"
)

func doSomething(t time.Duration, c chan<- int, param int) {
	time.Sleep(t)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	t1 := 4 * time.Second
	t2 := 2 * time.Second

	go doSomething(t1, c1, 1)
	go doSomething(t2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}
