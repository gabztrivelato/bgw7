package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
}

func main() {
	var salary int
	if salary <= 10000 {
		newErr := errors.New("error: salary is less than 10000")
		fmt.Println(newErr.Error())
	}
}
