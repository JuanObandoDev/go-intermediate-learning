package main

import (
	"fmt"
	"time"
)

// this pretends to be a father class in other languages
type Person struct {
	dni  string
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

func (p *Person) SetDni(dni string) {
	p.dni = dni
}

func (p *Person) GetDni() string {
	return p.dni
}

// this is equivalent to a class in other languages
type Employee struct {
	id       int
	vacation bool
	position string
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

func (e *Employee) SetPosition(position string) {
	e.position = position
}

func (e *Employee) GetPosition() string {
	return e.position
}

// creating a constructor function
func NewEmployee(id int, vacation bool) *Employee {
	return &Employee{id, vacation, ""}
}

// go doesn't use inheritance, but it can be simulated using composition
type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

// this is equivalent to a method in other languages
func (fte FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

// go doesn't use inheritance, but it can be simulated using composition
type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

// this is equivalent to a method in other languages
func (te TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

// this is equivalent to an interface in other languages
type PrintInfo interface {
	getMessage() string
}

// applying polymorphism
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

var GetFullTimeEmployeeById = func(id int, dni string) (FullTimeEmployee, error) {
	var fte FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return fte, err
	}

	fte.Employee = e

	p, err := GetPersonByDni(dni)
	if err != nil {
		return fte, err
	}

	fte.Person = p

	return fte, nil
}

var GetPersonByDni = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func main() {
	// firts way to declare a new instance of Employee
	// e := Employee{1, "Juan", false}
	e := Employee{1, false, ""}
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

	// creating a new instance of TemporaryEmployee
	te := TemporaryEmployee{}
	getMessage(te)
	getMessage(fte)
}
