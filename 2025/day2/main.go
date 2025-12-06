package main

import (
	"day2/part1"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code - Day 2")
	if len(os.Args) < 2 {
		fmt.Println("Exiting. Path to file must be passed.")
		fmt.Println("Usage: go run main.go <file-path>")
		os.Exit(1)
	}
	filePath := os.Args[1]

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file. Exiting.")
		os.Exit(1)
	}

	idSum, err := part1.FindInvalidIDs(content)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Invalid product ID sum:", idSum)
}
