package main

import "fmt"

type Students struct {
	FirstName           string
	LastName            string
	Id                  int
	DateOfMatriculation string
}

func (s *Students) showStudentDetail() {
	fmt.Printf("Id: %d, Name: %s %s, Date of Matriculation: %s\n", s.Id, s.FirstName, s.LastName, s.DateOfMatriculation)
}

func main() {
	student1 := Students{
		FirstName:           "John",
		LastName:            "Doe",
		Id:                  1,
		DateOfMatriculation: "2024-09-01",
	}

	student1.showStudentDetail()
}
