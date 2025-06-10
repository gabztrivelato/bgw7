package main

import "fmt"

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Id       int
	Position string
	Person   Person
}

func (e *Employee) PrintEmployee() {
	fmt.Println("Id:", e.Id, "Position:", e.Position, "Name:", e.Person.Name, "Date of Birth:", e.Person.DateOfBirth)
}

func main() {

	person := Person{
		Id:          2,
		Name:        "Jane Smith",
		DateOfBirth: "1992-02-02",
	}

	employee := Employee{
		Id:       1,
		Position: "Software Engineer",
		Person:   person,
	}

	employee.PrintEmployee()

}
