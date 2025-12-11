package part2

import (
	"os"
	"testing"
)

func TestFindInValidIDs_From_File(t *testing.T) {
	content, err := os.ReadFile("../product_id_ranges_text.txt")
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	actual, err := FindInvalidIDs(content)
	if err != nil {
		t.Fatalf("error while getting invalid ID sum: %v", err)
	}

	expected := 4174379265
	if actual != expected {
		t.Errorf("FAILED: got %d, expected %d", actual, expected)
	}
}

func TestIsInvalidID(t *testing.T) {
	for _, test := range []struct {
		Input    string
		Expected bool
	}{
		{
			Input:    "12341234",
			Expected: true,
		},
		{
			Input:    "1212121212",
			Expected: true,
		},
		{
			Input:    "1111111",
			Expected: true,
		},
		{
			Input:    "22",
			Expected: true,
		},
		{
			Input:    "1188511885",
			Expected: true,
		},
		{
			Input:    "446446",
			Expected: true,
		},
		{
			Input:    "38593859",
			Expected: true,
		},
		{
			Input:    "565656",
			Expected: true,
		},
		{
			Input:    "824824824",
			Expected: true,
		},
		{
			Input:    "2121212121",
			Expected: true,
		},
		{
			Input:    "123456",
			Expected: false,
		},
		{
			Input:    "1",
			Expected: false,
		},
		{
			Input:    "212121212",
			Expected: false,
		},
		{
			Input:    "",
			Expected: false,
		},
	} {
		t.Run(test.Input, func(t *testing.T) {
			actual := isInvalidID(test.Input)
			if actual != test.Expected {
				t.Errorf("FAILED: got %v, expected %v for %s", actual, test.Expected, test.Input)
			}
		})
	}
}
