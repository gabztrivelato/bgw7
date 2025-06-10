package main

import (
	"fmt"
	"os"
)

func readFile(filename string) (string, error) {
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

	return string(data), nil
}

func main() {
	data, err := readFile("customer.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(data)
}
