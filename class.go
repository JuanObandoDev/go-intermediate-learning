package main

import "fmt"

// this is equivalent to a class in other languages
type Employee struct {
	id   int
	name string
}

// this is equivalent to a method in other languages
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{1, "Juan"}
	fmt.Printf("%+v\n", e)
	e.SetId(2)
	e.SetName("Pedro")
	fmt.Printf("id: %d, name: %s", e.GetId(), e.GetName())
}
