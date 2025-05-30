package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Idade Benjamin: %d\n", employees["Benjamin"])

	var employeesWithAgeGreaterThan21 int
	for _, age := range employees {
		if age < 21 {
			employeesWithAgeGreaterThan21++
		}
	}
	fmt.Printf("Quantidade de funcionários com idade maior que 21: %d\n", employeesWithAgeGreaterThan21)
	employees["Federico"] = 25
	delete(employees, "Pedro")
}
