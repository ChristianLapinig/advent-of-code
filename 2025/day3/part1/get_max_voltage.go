package part1

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	InvalidBank = "invalid bank. bank must be more than 2 characters long"
)

func GetMaxVoltage(content []byte) (int, error) {
	banks := strings.Split(string(content), "\n")
	voltageSum := 0
	for _, bank := range banks {
		if len(bank) < 2 {
			continue
		}
		maxVoltage, err := getMaxVoltageFromBank(bank)
		if err != nil {
			return -1, fmt.Errorf("error getting max voltage from bank %s: %v", bank, err)
		}
		voltageSum += maxVoltage
	}
	return voltageSum, nil
}

func getMaxVoltageFromBank(bank string) (int, error) {
	if len(bank) < 2 {
		return -1, fmt.Errorf("%s: %s", InvalidBank, bank)
	}

	left := 0
	startingVoltage := fmt.Sprintf("%s%s", string(bank[left]), string(bank[left+1]))
	maxVoltage, err := strconv.Atoi(startingVoltage)
	if err != nil {
		return -1, nil
	}
	for right := left + 1; right < len(bank); right++ {
		voltageStr := fmt.Sprintf("%s%s", string(bank[left]), string(bank[right]))
		voltage, err := strconv.Atoi(voltageStr)
		if err != nil {
			return -1, err
		}
		maxVoltage = max(maxVoltage, voltage)
		if bank[right] > bank[left] {
			left = right
		}
	}
	return maxVoltage, nil
}
