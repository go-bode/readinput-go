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
	var input string
	fmt.Println("Press a option above")
	fmt.Println("h - show hostname")
	fmt.Println("e - exit program")
	fmt.Println("t - show user temp dir")
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("Invalid input")
	}

	switch input {
	case "h":
		fmt.Println(os.Hostname())
	case "e":
		os.Exit(2)
	case "t":
		fmt.Println(os.TempDir())
	default:
		fmt.Println("padrão")
	}
}
