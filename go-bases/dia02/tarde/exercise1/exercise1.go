package main

import "fmt"

func main() {
	fmt.Print("Digite uma palavra: ")
	var userWord string
	fmt.Scanln(&userWord)
	for i := 0; i < len(userWord); i++ {
		fmt.Printf("%c\n ", userWord[i])
	}
}
