package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	if filePath == "" {
		fmt.Println("File name must be passed as an argument. Exiting.")
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error opening file %s. Exiting.\n", filePath)
		log.Fatalf("%v\n", err)
	}
	defer file.Close()
	password, err := getPassword(file)
	if err != nil {
		fmt.Println("Error retrieving password. Exiting.")
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("The password is %d\n", password)
}
