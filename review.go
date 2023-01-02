package main

import (
	"fmt"
	"strconv"
	"time"
)

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x, y)

	// parsing string to int
	value, err := strconv.ParseInt("7", 0, 64)

	// validation of parsing error
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(value)
	}

	// maps (key, value)
	m := make(map[string]int)
	m["k1"] = 7
	fmt.Println(m["k1"])

	// slices
	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Println(i, v)
	}

	// append to slice
	s = append(s, 16)
	for i, v := range s {
		fmt.Println(i, v)
	}

	// channels
	c := make(chan int)

	// goroutines
	go doSomething(c)
	<-c

	// pointers
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}
