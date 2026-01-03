package main

import (
	"fmt"

	"github.com/stephen-condon/advent-of-code-2025/utilities"
)

func main() {
	fmt.Println("Part One:")
	executePartOne("example.txt")
	executePartOne("input.txt")

	fmt.Println("\nPart Two:")
	executePartTwo("example.txt")
	executePartTwo("input.txt")
}

func executePartOne(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	accessibleCount := countAccessibleRolls(input)
	fmt.Printf("%s: %d\n", filename, accessibleCount)
}

func executePartTwo(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	totalRemoved := removeAccessibleRolls(input)
	fmt.Printf("%s: %d\n", filename, totalRemoved)
}

func removeAccessibleRolls(grid []string) int {
	// Create a mutable copy of the grid
	mutableGrid := make([][]rune, len(grid))
	for i, line := range grid {
		mutableGrid[i] = []rune(line)
	}

	totalRemoved := 0

	for {
		accessiblePositions := findAccessiblePositions(mutableGrid)

		if len(accessiblePositions) == 0 {
			break
		}

		for _, pos := range accessiblePositions {
			mutableGrid[pos[0]][pos[1]] = '.'
		}

		totalRemoved += len(accessiblePositions)
	}

	return totalRemoved
}

func findAccessiblePositions(grid [][]rune) [][2]int {
	var positions [][2]int

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' && isAccessibleMutable(grid, row, col) {
				positions = append(positions, [2]int{row, col})
			}
		}
	}

	return positions
}

func isAccessibleMutable(grid [][]rune, row, col int) bool {
	directions := [][]int{
		{-1, 0}, {-1, 1}, {0, 1}, {1, 1},
		{1, 0}, {1, -1}, {0, -1}, {-1, -1},
	}

	rollCount := 0

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[newRow]) {
			if grid[newRow][newCol] == '@' {
				rollCount++
			}
		}
	}

	return rollCount <= 3
}

func countAccessibleRolls(grid []string) int {
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' && isAccessible(grid, row, col) {
				count++
			}
		}
	}

	return count
}

func isAccessible(grid []string, row, col int) bool {
	directions := [][]int{
		{-1, 0},  // North
		{-1, 1},  // NorthEast
		{0, 1},   // East
		{1, 1},   // SouthEast
		{1, 0},   // South
		{1, -1},  // SouthWest
		{0, -1},  // West
		{-1, -1}, // NorthWest
	}

	rollCount := 0

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[newRow]) {
			if grid[newRow][newCol] == '@' {
				rollCount++
			}
		}
	}

	return rollCount <= 3
}
