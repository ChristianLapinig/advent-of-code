package part1

import (
	"os"
	"testing"
)

func TestFindInvalidProductIDs_From_File(t *testing.T) {
	content, err := os.ReadFile("../product_id_ranges_text.txt")
	if err != nil {
		t.Fatalf("error reading file: %q", err)
	}

	actual, err := FindInvalidIDs(content)
	expected := 1227775554
	if err != nil {
		t.Fatalf("error finding invalid product ids: %q", err)
	}

	if actual != expected {
		t.Errorf("FAILED: expected %d, got %d", expected, actual)
	}
}
