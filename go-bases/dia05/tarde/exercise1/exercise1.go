package main

import (
	"fmt"
	"os"
)

func readFile(filename string) (string, error) {
	_, err := os.Open(filename)

	defer func() {
		fmt.Println("execução concluída")
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	return "File opened", nil
}

func main() {
	readFile("customer.txt")
}
