package main

import "fmt"

func main() {
	var age int
	var yearsAtCompany int
	var salary float64

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&age)
	fmt.Print("Digite quantos anos você trabalha na empresa: ")
	fmt.Scanln(&yearsAtCompany)
	fmt.Print("Digite seu salário: ")
	fmt.Scanln(&salary)

	if age > 22 && yearsAtCompany > 1 {
		if salary < 100000.0 {
			fmt.Println("Você é elegível para o empréstimo com juros.")
		} else {
			fmt.Println("Você é elegível para o empréstimo sem juros.")
		}
	} else {
		fmt.Println("Você não é elegível para o empréstimo.")
	}
}
