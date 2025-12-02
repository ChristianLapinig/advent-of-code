package part1

import (
	"os"
	"strings"
	"testing"
)

func TestGetPassword(t *testing.T) {
	t.Run("testing from file", func(t *testing.T) {
		data, err := os.ReadFile("rotations_test.txt")
		if err != nil {
			t.Fatalf("FAILED - Error loading file: %v", err)
		}

		reader := strings.NewReader(string(data))
		actual, err := GetPassword(reader)
		if err != nil {
			t.Fatalf("FAILED - Error while getting password: %v", err)
		}

		if actual != 3 {
			t.Errorf("FAILED - expected 3, got %d", actual)
		}
	})

	t.Run("custom inputs", func(t *testing.T) {
		for _, test := range []struct {
			Name     string
			Content  []byte
			Expected int
		}{
			{
				Name:     "Empty content should return 0",
				Content:  []byte(""),
				Expected: 0,
			},
			{
				Name:     "L50 R225 L15 L10 R20 L20 L25 should return 33",
				Content:  []byte("L50\nR225\nL15\nL10\nR20\nL20\nL25"),
				Expected: 3,
			},
		} {
			t.Run(test.Name, func(t *testing.T) {
				reader := strings.NewReader(string(test.Content))
				actual, err := GetPassword(reader)
				if err != nil {
					t.Fatalf("FAILED - Error while getting password: %v", err)
				}

				if actual != test.Expected {
					t.Errorf("FAILED - expected %d, got %d", test.Expected, actual)
				}
			})
		}
	})
}
