package part2

import (
	"os"
	"testing"
)

func TestGetMaxVoltageFromBank(t *testing.T) {
	const Limit = 12
	for _, test := range []struct {
		Input    string
		Exepcted int
	}{
		{
			Input:    "811111111111119",
			Exepcted: 811111111119,
		},
		{
			Input:    "234234234234278",
			Exepcted: 434234234278,
		},
		{
			Input:    "818181911112111",
			Exepcted: 888911112111,
		},
		{
			Input:    "987654321111111",
			Exepcted: 987654321111,
		},
	} {
		t.Run(test.Input, func(t *testing.T) {
			actual, err := getMaxVoltageFromBank(test.Input, Limit)
			if err != nil {
				t.Fatalf("error calculating max voltage: %v", err)
			}
			if actual != test.Exepcted {
				t.Errorf("FAILED: got %d, expected %d from input %s", actual, test.Exepcted, test.Input)
			}
		})
	}
}

func TestGetMaxVoltage_From_File(t *testing.T) {
	content, err := os.ReadFile("../test_input.txt")
	if err != nil {
		t.Fatalf("error loading file: %v", err)
	}

	actual, err := GetMaxVoltage(content)
	if err != nil {
		t.Fatalf("error calculating max voltage: %v", err)
	}

	expected := 3121910778619
	if actual != expected {
		t.Errorf("FAILED: got %d, expected %d", actual, expected)
	}
}
