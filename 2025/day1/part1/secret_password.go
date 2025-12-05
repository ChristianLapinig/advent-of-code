package part1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"day1/constants"
)

func GetPassword(r io.Reader) (int, error) {
	current := constants.StartingDialVal
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

		// ignore full spins
		if rotation >= constants.DialSize {
			rotation %= constants.DialSize
		}

		if direction == constants.DirectionLeft {
			rotation = -rotation
		}

		current = ((current+rotation)%constants.DialSize + constants.DialSize) % constants.DialSize

		if current == 0 {
			zeroCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return zeroCount, nil
}
