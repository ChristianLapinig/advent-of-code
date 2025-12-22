package main

import (
	"fmt"
	"log"
	"os"

	"day4/part1"
)

func main() {
	fmt.Println("Advent of Code 2025 - Day 4: Printing Department")
	if len(os.Args) < 2 {
		fmt.Println("Path to input file required.")
		fmt.Println("Usage: go run main.go <input-file>")
		os.Exit(1)
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("error reading file %v", err)
	}

	fmt.Println("Part 1:", part1.CountAccessibleRolls(content))
}
