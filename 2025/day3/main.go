package main

import (
	"fmt"
	"log"
	"os"

	"day3/part1"
	"day3/part2"
)

func main() {
	fmt.Println("Advent of Code - 2025 - Day 3: Lobby")
	if len(os.Args) < 2 {
		fmt.Println("Path to input file required.")
		fmt.Println("Usage: go run main.go <input-file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error reading file %v", err)
	}

	result1, err := part1.GetMaxVoltage(content)
	if err != nil {
		log.Fatalf("error calculating result for part 1: %v", err)
	}
	fmt.Println("Part 1 answer: ", result1)

	result2, err := part2.GetMaxVoltage(content)
	if err != nil {
		log.Fatalf("error calculating result for part 2: %v", err)
	}
	fmt.Println("Part 2 answer: ", result2)
}
