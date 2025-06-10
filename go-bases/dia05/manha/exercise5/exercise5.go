package main

import "fmt"

type SalaryError struct {
}

func (e *SalaryError) Error() string {
	return "Error: the worker cannot have worked less than 80 hours per month"
}

func calcSalary(hoursWorked int, hourlyRate float64) (float64, error) {
	if hoursWorked < 80 || hoursWorked < 0 {
		return 0, &SalaryError{}
	}
	salary := float64(hoursWorked) * hourlyRate
	if salary >= 150000 {
		salary = salary * 0.9
	}
	return salary, nil
}

func main() {
	salary, err := calcSalary(80, 100.0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Salary: ", salary)
	}
}
