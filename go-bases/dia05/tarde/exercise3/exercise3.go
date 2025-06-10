package main

import (
	"fmt"
	"os"
	"strings"
)

type validationError struct {
	name string
}

func (e *validationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.name)
}

func addClient(file string, name string, id int, phoneNumber string, address string) string {
	matriz := readFile(file)

	defer func() {
		fmt.Println("End of execution")
	}()

	defer func() {
		fmt.Println("Several erros were detected at runtime")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	clientes := make(map[string]bool)
	for i, linha := range matriz {
		id := linha[0]
		if clientes[id] {
			panic("Error: client already exists")
		}
		i++
		clientes[id] = true
	}

	if err := validateData(name, id, phoneNumber, address); err != nil {
		panic(err)
	}

	return "Client added successfully"
}

func validateData(name string, id int, phoneNumber string, address string) error {
	err := &validationError{}
	if name == "" {
		err = &validationError{name: "name cannot be empty"}
		return err
	}
	if id <= 0 {
		err = &validationError{name: "id must be a positive integer"}
		return err
	}
	if phoneNumber == "" {
		err = &validationError{name: "phone number cannot be empty"}
		return err
	}
	if address == "" {
		err = &validationError{name: "address cannot be empty"}
		return err
	}
	return nil
}

func readFile(filename string) [][]string {
	file, err := os.Open(filename)

	defer func() {
		fmt.Println("execução concluída")
		file.Close()
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	linhas := strings.Split(string(data), "\n")

	var matriz [][]string
	for _, linha := range linhas {
		linha = strings.TrimSpace(linha)
		if linha == "" {
			continue
		}
		colunas := strings.Split(linha, ",")
		matriz = append(matriz, colunas)
	}
	return matriz
}

func main() {
	fmt.Println(addClient(("customer.txt"), "John Doe", 123, "555-1234", "123 Elm St"))
}
