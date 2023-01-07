package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	cint := make(chan int, 2) // this is the stoplight, 2 is the number of goroutines allowed to execute at the same time
	for i := 0; i < 10; i++ {
		cint <- 1
		wg.Add(1)
		go DoSomething(&wg, cint)
	}
	wg.Wait()
}

func DoSomething(wg *sync.WaitGroup, cint chan int) {
	defer wg.Done()
	fmt.Println("DoSomething start")
	time.Sleep(4 * time.Second)
	fmt.Println("DoSomething end")
	<-cint
}
