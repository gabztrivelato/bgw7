package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
}

func (e *SalaryError) Error() string {
	return "Error: salary is less than 10000"
}

func main() {
	var salary int
	if salary <= 10000 {
		err := &SalaryError{}
		fmt.Println(errors.Is(err, &SalaryError{}))
	}
}
