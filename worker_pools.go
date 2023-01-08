package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := len(tasks)
	jobs := make(chan int, len(tasks))
	res := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, res)
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-res
	}
}

func Worker(id int, jobs <-chan int, res chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started job %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, finished job %d and fib %d\n", id, job, fib)
		res <- fib
	}
}

func Fibonacci(a int) int {
	if a <= 1 {
		return a
	}
	return Fibonacci(a-1) + Fibonacci(a-2)
}
