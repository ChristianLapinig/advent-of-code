package part2

import (
	"strconv"
	"strings"
)

const (
	InvalidBank = "invalid bank. bank must be more than 2 characters long"
	DigitLimit  = 12
)

func GetMaxVoltage(content []byte) (int, error) {
	banks := strings.Split(string(content), "\n")
	voltageSum := 0

	for _, bank := range banks {
		maxVoltage, err := getMaxVoltageFromBank(bank, DigitLimit)
		if err != nil {
			return -1, err
		}
		voltageSum += maxVoltage
	}

	return voltageSum, nil
}

func getMaxVoltageFromBank(bank string, limit int) (int, error) {
	n := len(bank)

	if limit >= n {
		return toInt(bank), nil
	}

	if limit == 0 {
		return 0, nil
	}

	var result strings.Builder
	result.Grow(limit)

	start := 0 // starting position for searching
	for i := range limit {
		remaining := limit - i - 1
		end := n - remaining

		// find maximum digit within the current range
		maxDigit := bank[start]
		maxPosition := start

		for j := start; j < end; j++ {
			if bank[j] > maxDigit {
				maxDigit = bank[j]
				maxPosition = j
			}
		}

		result.WriteByte(maxDigit)
		start = maxPosition + 1
	}

	return toInt(result.String()), nil
}

func toInt(str string) int {
	if str == "" {
		return 0
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
