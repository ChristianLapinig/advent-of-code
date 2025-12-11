package part2

import (
	"strconv"
	"strings"
)

func FindInvalidIDs(content []byte) (int, error) {
	idSum := 0
	line := strings.ReplaceAll(string(content), "\n", "")
	idRangeLines := strings.Split(line, ",")
	var idRanges [][]int

	// Generate ID pairs
	for _, idLine := range idRangeLines {
		pair := strings.Split(idLine, "-")
		low, err := strconv.Atoi(pair[0])
		if err != nil {
			return -1, err
		}
		high, err := strconv.Atoi(pair[1])
		if err != nil {
			return -1, err
		}
		idRanges = append(idRanges, []int{low, high})
	}

	// TODO: Optimize this if possible
	for _, idRange := range idRanges {
		for i := idRange[0]; i <= idRange[1]; i++ {
			id := strconv.Itoa(i)
			if isInvalidID(id) {
				idSum += i
			}
		}
	}

	return idSum, nil
}

func isInvalidID(s string) bool {
	if s == "" || len(s) < 2 {
		return false
	}

	n := len(s)

	for length := 1; length <= n/2; length++ {
		// search chunks that divide evenly
		if n%length == 0 {
			pattern := s[:length]

			// check if pattern repeats
			matches := true
			for i := 0; i < n; i += length {
				if s[i:i+length] != pattern {
					matches = false
					break
				}
			}

			if matches {
				return true
			}
		}
	}
	return false
}
