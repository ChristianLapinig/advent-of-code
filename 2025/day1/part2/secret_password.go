package part2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func GetPassword(r io.Reader) (int, error) {
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

		// rotations > 100 are considered a full rotation
		// meaning we have passed 0 rotation / 100 times
		if rotation >= 100 {
			zeroCount += rotation / 100
			rotation %= 100 // get the number clicks after 0 has been passed
		}

		var rotationVal int
		if direction == "L" {
			rotationVal = currentDialVal - rotation
		} else { // direction == "R"
			rotationVal = currentDialVal + rotation
		}

		// 0 has been passed while turning the dial
		if (rotationVal < 0 || rotationVal > 100) && currentDialVal != 0 {
			zeroCount++
		}

		if rotationVal < 0 {
			currentDialVal = 100 + rotationVal
		} else if rotationVal > 99 {
			currentDialVal = rotationVal - 100
		} else {
			currentDialVal = rotationVal
		}

		// dial has landed on 0 or has passed 0
		if currentDialVal == 0 {
			zeroCount++
		}
	}
	return zeroCount, nil
}
