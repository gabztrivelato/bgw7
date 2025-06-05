package main

import (
	"fmt"
)

type SalaryError struct {
}

func main() {
	var salary int
	if salary <= 10000 {
		newErr := fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		fmt.Println(newErr.Error())
	}
}
