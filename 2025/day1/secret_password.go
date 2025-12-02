package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func getPassword(r io.Reader) (int, error) {
	currentDialVal := 50
	zeroCount := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error converting string to number. Exiting.")
			return -1, err
		}

		// rotations > 100 should be considered a full rotation
		if rotation > 100 {
			rotation %= 100
		}

		var rotationValue int
		if "L" == direction {
			rotationValue = currentDialVal - rotation
		} else { // direction == "R"
			rotationValue = currentDialVal + rotation
		}

		if rotationValue < 0 {
			currentDialVal = 100 + rotationValue
		} else if rotationValue > 99 {
			currentDialVal = rotationValue - 100
		} else {
			currentDialVal = rotationValue
		}

		if 0 == currentDialVal {
			zeroCount++
		}
	}
	return zeroCount, nil
}
