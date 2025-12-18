package part1

import (
	"fmt"
	"os"
	"testing"
)

func TestGetMaxVoltageFromBank(t *testing.T) {
	for _, test := range []struct {
		Input    string
		Expected int
	}{
		{
			Input:    "987654321111111",
			Expected: 98,
		},
		{
			Input:    "811111111111119",
			Expected: 89,
		},
		{
			Input:    "234234234234278",
			Expected: 78,
		},
		{
			Input:    "818181911112111",
			Expected: 92,
		},
		{
			Input:    "2",
			Expected: -1,
		},
		{
			Input:    "",
			Expected: -1,
		},
	} {
		t.Run(test.Input, func(t *testing.T) {
			actual, err := getMaxVoltageFromBank(test.Input)
			if err != nil && err.Error() != fmt.Sprintf("%s: %s", InvalidBank, test.Input) {
				t.Fatalf("an error occurred: %v", err)
			}
			if actual != test.Expected {
				t.Errorf("FAILED: got %d, expected %d", actual, test.Expected)
			}
		})
	}
}

func TestGetMaxVoltage_From_File(t *testing.T) {
	content, err := os.ReadFile("../test_input.txt")
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	actual, err := GetMaxVoltage(content)
	if err != nil {
		t.Fatalf("error calculating max joltage: %v", err)
	}
	expected := 357
	if actual != expected {
		t.Errorf("FAILED: got %d, expected %d", actual, expected)
	}
}
