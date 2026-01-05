package main

import (
	"fmt"

	"github.com/stephen-condon/advent-of-code-2025/utilities"
)

type Beam struct {
	row, col int
}

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

	splitCount := simulateBeams(input)
	fmt.Printf("%s: %d\n", filename, splitCount)
}

func executePartTwo(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	pathCount := countAllPaths(input)
	fmt.Printf("%s: %d\n", filename, pathCount)
}

func countAllPaths(grid []string) int {
	startCol := -1
	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == 'S' {
			startCol = col
			break
		}
	}

	if startCol == -1 {
		return 0
	}

	// Use memoization to cache path counts from each state
	// Key: "row,col,direction" where direction is the choice made at last splitter
	memo := make(map[string]int)

	var countPaths func(row, col int, fromDirection string) int
	countPaths = func(row, col int, fromDirection string) int {
		row++

		if row >= len(grid) {
			return 1
		}

		if col < 0 || col >= len(grid[row]) {
			return 1
		}

		key := fmt.Sprintf("%d,%d,%s", row, col, fromDirection)
		if count, exists := memo[key]; exists {
			return count
		}

		cell := grid[row][col]
		totalPaths := 0

		if cell == '^' {
			// At a splitter, we have two choices: go left or go right
			// Count paths from both choices
			leftPaths := countPaths(row, col-1, "L")
			rightPaths := countPaths(row, col+1, "R")
			totalPaths = leftPaths + rightPaths
		} else if cell == '.' {
			totalPaths = countPaths(row, col, fromDirection)
		}

		memo[key] = totalPaths
		return totalPaths
	}

	return countPaths(0, startCol, "START")
}

func simulateBeams(grid []string) int {
	startCol := -1
	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == 'S' {
			startCol = col
			break
		}
	}

	if startCol == -1 {
		return 0
	}

	beams := []Beam{{row: 0, col: startCol}}
	visited := make(map[string]bool)
	splitCount := 0

	for len(beams) > 0 {
		var nextBeams []Beam

		for _, beam := range beams {
			beam.row++

			if beam.row >= len(grid) {
				continue
			}

			if beam.col < 0 || beam.col >= len(grid[beam.row]) {
				continue
			}

			cell := grid[beam.row][beam.col]

			if cell == '^' {
				splitCount++

				leftBeam := Beam{row: beam.row, col: beam.col - 1}
				rightBeam := Beam{row: beam.row, col: beam.col + 1}

				leftKey := fmt.Sprintf("%d,%d", leftBeam.row, leftBeam.col)
				rightKey := fmt.Sprintf("%d,%d", rightBeam.row, rightBeam.col)

				if !visited[leftKey] {
					visited[leftKey] = true
					nextBeams = append(nextBeams, leftBeam)
				}
				if !visited[rightKey] {
					visited[rightKey] = true
					nextBeams = append(nextBeams, rightBeam)
				}
			} else if cell == '.' {
				beamKey := fmt.Sprintf("%d,%d", beam.row, beam.col)
				if !visited[beamKey] {
					visited[beamKey] = true
					nextBeams = append(nextBeams, beam)
				}
			}
		}

		beams = nextBeams
	}

	return splitCount
}
