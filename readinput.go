package main

import "fmt"
import "os"

func main() {
	fmt.Println("Press some key...")
	for {
		readInput()
	}
}

func readInput() {
	var input int
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("Invalid input")
	}

	switch input {
	case 1:
		fmt.Println("hi")
	case 2:
		os.Exit(2)
	default:
		fmt.Println("padrão")
	}
}
