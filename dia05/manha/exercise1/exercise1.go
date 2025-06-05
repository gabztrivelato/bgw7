package main

import (
	"fmt"
)

type SalaryError struct {
}

func (e *SalaryError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum."
}

func main() {
	var salary int
	if salary < 150000 {
		err := SalaryError{}
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Must pay tax")
}
