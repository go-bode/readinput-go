package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
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
		fmt.Println("Invalid input key")
	}

	switch input {
	case "h":
		color.Cyan(os.Hostname())
	case "e":
		os.Exit(2)
	case "t":
		fmt.Println("The OS Temp dir is --> ")
		color.Cyan(os.TempDir())
	default:
		fmt.Println("padrão")
	}
}
