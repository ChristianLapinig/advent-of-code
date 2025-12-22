package part1

import (
	"os"
	"testing"
)

func TestCountAccessibleRolls_From_File(t *testing.T) {
	content, err := os.ReadFile("../test_input.txt")
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	actual := CountAccessibleRolls(content)
	expected := 13
	if actual != expected {
		t.Errorf("FAILED: got %d, expected %d", actual, expected)
	}
}
