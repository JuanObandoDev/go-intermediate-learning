package main

import "fmt"

// variadic functions can be called with any number of trailing arguments.
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// returns with name
func getValues(x int) (double int, triple int, quadruple int) {
	double = x * 2
	triple = x * 3
	quadruple = x * 4
	return
}

func main() {
	fmt.Println(sum(16, 20))
	fmt.Println(sum(20, 24, 32, 16))
	fmt.Println(getValues(2))
}
