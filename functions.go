package main

import "fmt"

func main() {
	x := func(n int) int {
		return n * 2
	}(5)
	fmt.Println(x)
}
