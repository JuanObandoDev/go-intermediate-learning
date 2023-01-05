package main

import "fmt"

// this pretends to be a father class in other languages
type Person struct {
	name string
	age  int
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetAge() int {
	return p.age
}

// this is equivalent to a class in other languages
type Employee struct {
	id       int
	vacation bool
}

// this is equivalent to a method in other languages
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) SetVacation(vacation bool) {
	e.vacation = vacation
}

func (e *Employee) GetVacation() bool {
	return e.vacation
}

// creating a constructor function
func NewEmployee(id int, vacation bool) *Employee {
	return &Employee{id, vacation}
}

// go doesn't use inheritance, but it can be simulated using composition
type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	// firts way to declare a new instance of Employee
	// e := Employee{1, "Juan", false}
	e := Employee{1, false}
	fmt.Printf("%+v\n", e)
	e.SetId(2)
	// e.SetName("Pedro")
	// fmt.Printf("id: %d, name: %s, vacation: %t\n", e.GetId(), e.GetName(), e.GetVacation())
	fmt.Printf("id: %d, vacation: %t\n", e.GetId(), e.GetVacation())

	// second way to declare a new instance of Employee
	e2 := new(Employee) // this returns a pointer to the new instance
	e2.SetId(3)
	// e2.SetName("Maria")
	e2.SetVacation(true)
	// fmt.Printf("id: %d, name: %s, vacation: %t\n", e2.GetId(), e2.GetName(), e2.GetVacation())
	fmt.Printf("id: %d, vacation: %t\n", e2.GetId(), e2.GetVacation())

	// third way to declare a new instance of Employee
	// e3 := NewEmployee(4, "Jose", false)
	e3 := NewEmployee(4, false)
	fmt.Printf("%+v\n", *e3)
	// fmt.Printf("id: %d, name: %s, vacation: %t\n", e3.GetId(), e3.GetName(), e3.GetVacation())
	fmt.Printf("id: %d, vacation: %t\n", e3.GetId(), e3.GetVacation())

	// creating a new instance of FullTimeEmployee
	fte := FullTimeEmployee{}
	fte.SetId(5)
	fte.SetName("Luis")
	fte.SetAge(30)
	fte.SetVacation(false)
	fmt.Printf("%+v\n", fte)
}
