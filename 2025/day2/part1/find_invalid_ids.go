package part1

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

	for _, idRange := range idRanges {
		for i := idRange[0]; i <= idRange[1]; i++ {
			idStr := strconv.Itoa(i)
			mid := len(idStr) / 2
			left, right := idStr[:mid], idStr[mid:]

			if left == right {
				idSum += i
			}
		}
	}

	return idSum, nil
}
