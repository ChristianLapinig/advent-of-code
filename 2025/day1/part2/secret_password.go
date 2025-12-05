package part2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"day1/constants"
)

func GetPassword(r io.Reader) (int, error) {
	current := 50
	zeroCount := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) < 2 {
			return -1, fmt.Errorf("invalid line: %q", line)
		}

		direction := line[0]
		rotation, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return -1, err
		}

		// count full rotations
		if rotation < 0 {
			zeroCount += (-rotation) / constants.DialSize
		} else {
			zeroCount += rotation / constants.DialSize
		}

		rotation = rotation % constants.DialSize
		if direction == constants.DirectionLeft {
			rotation = -rotation
		}

		next := current + rotation
		wrapped := (next < 0 || next >= constants.DialSize) && current != 0
		current = (next%constants.DialSize + constants.DialSize) % constants.DialSize
		if current == 0 || wrapped {
			zeroCount++
		}
	}
	return zeroCount, nil
}
