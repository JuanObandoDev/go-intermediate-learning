package main

import "testing"

func TestGetEmployeeById(t *testing.T) {
	tables := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			1,
			"1",
			func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id:       1,
						position: "CEO",
					}, nil
				}

				GetPersonByDni = func(dni string) (Person, error) {
					return Person{
						name: "John Doe",
						age:  35,
						dni:  "1",
					}, nil
				}
			},
			FullTimeEmployee{
				Person: Person{
					age:  35,
					dni:  "1",
					name: "John Doe",
				},
				Employee: Employee{
					id:       1,
					position: "CEO",
				},
			},
		},
	}

	OriginalGetEmployeeById := GetEmployeeById
	OriginalGetPersonByDni := GetPersonByDni

	for _, item := range tables {
		item.mockFunc()
		fte, err := GetFullTimeEmployeeById(item.id, item.dni)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if fte.age != item.expectedEmployee.age {
			t.Errorf("Error, got: %d, want: %d.", fte.age, item.expectedEmployee.age)
		}

		_, err = GetPersonByDni(item.dni)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	}

	GetEmployeeById = OriginalGetEmployeeById
	GetPersonByDni = OriginalGetPersonByDni
}
