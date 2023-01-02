package main

import "fmt"

// this is equivalent to a class in other languages
type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{1, "Juan"}
	fmt.Printf("%+v\n", e)
}
