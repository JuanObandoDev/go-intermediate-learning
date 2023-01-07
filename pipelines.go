package main

import "fmt"

// reader channel <-chan int
// writer channel chan<- int

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(cin <-chan int, cout chan<- int) {
	for value := range cin {
		cout <- value * 2
	}
	close(cout)
}

func Printer(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	double := make(chan int)

	go Generator(generator)
	go Double(generator, double)
	Printer(double)
}
