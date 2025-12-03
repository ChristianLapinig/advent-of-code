package part2

import (
	"os"
	"strings"
	"testing"
)

func TestGetPassword(t *testing.T) {
	t.Run("testing from file", func(t *testing.T) {
		data, err := os.ReadFile("../rotations_test.txt")
		if err != nil {
			t.Fatalf("FAILED - Error loading file: %v", err)
		}

		reader := strings.NewReader(string(data))
		actual, err := GetPassword(reader)
		if err != nil {
			t.Fatalf("FAILED - Error while getting password: %v", err)
		}

		if actual != 6 {
			t.Errorf("FAILED - expected 6, got %d", actual)
		}
	})
}
