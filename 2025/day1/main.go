package main

import (
	"day1/part1"
	"day1/part2"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Exiting. Path to file must be passed.")
		fmt.Println("Usage: go run main.go <file-path>")
		os.Exit(1)
	}
	filePath := os.Args[1]

	reader1, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error opening file %s. Exiting.\n", filePath)
		log.Fatalf("%v\n", err)
	}
	defer reader1.Close()
	password1, err := part1.GetPassword(reader1)
	if err != nil {
		fmt.Println("Error retrieving password. Exiting.")
		log.Fatalf("%v\n", err)
	}

	reader2, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error opening file %s. Exiting.\n", filePath)
		log.Fatalf("%v\n", err)
	}
	defer reader2.Close()
	password2, err := part2.GetPassword(reader2)
	if err != nil {
		fmt.Println("Error retrieving password. Exiting.")
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("The password is %d\n", password1)
	fmt.Printf("Wait. Hold on. The correct password is %d\n", password2)
}
