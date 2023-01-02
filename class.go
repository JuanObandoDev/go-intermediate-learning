package main

import "fmt"

// this is equivalent to a class in other languages
type Employee struct {
	id       int
	name     string
	vacation bool
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

func (e *Employee) SetVacation(vacation bool) {
	e.vacation = vacation
}

func (e *Employee) GetVacation() bool {
	return e.vacation
}

// creating a constructor function
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{id, name, vacation}
}

func main() {
	// firts way to declare a new instance of Employee
	e := Employee{1, "Juan", false}
	fmt.Printf("%+v\n", e)
	e.SetId(2)
	e.SetName("Pedro")
	fmt.Printf("id: %d, name: %s, vacation: %t\n", e.GetId(), e.GetName(), e.GetVacation())

	// second way to declare a new instance of Employee
	e2 := new(Employee) // this returns a pointer to the new instance
	e2.SetId(3)
	e2.SetName("Maria")
	e2.SetVacation(true)
	fmt.Printf("id: %d, name: %s, vacation: %t\n", e2.GetId(), e2.GetName(), e2.GetVacation())

	// third way to declare a new instance of Employee
	e3 := NewEmployee(4, "Jose", false)
	fmt.Printf("%+v\n", *e3)
	fmt.Printf("id: %d, name: %s, vacation: %t\n", e3.GetId(), e3.GetName(), e3.GetVacation())
}
