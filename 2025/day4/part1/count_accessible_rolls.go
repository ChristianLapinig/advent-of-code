package part1

import (
	"strings"
)

type Direction struct {
	row, col int
}

func CountAccessibleRolls(content []byte) int {
	count := 0
	grid := strings.Split(string(content), "\n")
	grid = grid[:len(grid)-1] // Remove last newline

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				if isAccessibleRoll(grid, i, j) {
					count++
				}
			}
		}
	}

	return count
}

func isAccessibleRoll(grid []string, row, col int) bool {
	directions := []Direction{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for _, direction := range directions {
		r := row + direction.row
		c := col + direction.col
		// point must be within the bounds of the grid
		if r >= 0 && r < rows && c >= 0 && c < cols {
			if grid[r][c] == '@' {
				count++
			}
		}
	}

	return count < 4
}
