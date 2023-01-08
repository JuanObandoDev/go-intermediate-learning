package main

import (
	"fmt"
	"time"
)

func main() {
	cint1 := make(chan int)
	cint2 := make(chan int)
	d1 := time.Second * 4
	d2 := time.Second * 2

	go DoSomething(d1, cint1, 1)
	go DoSomething(d2, cint2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-cint1:
			fmt.Println("Channel 1: ", channelMsg1)

		case channelMsg2 := <-cint2:
			fmt.Println("Channel 2: ", channelMsg2)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
