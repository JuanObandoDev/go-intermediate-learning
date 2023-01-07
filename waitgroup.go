package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go DoSomething(&wg)
	}

	wg.Wait()
}

func DoSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
}
