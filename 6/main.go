package main

import (
	"fmt"
	"strconv"
	"strings"

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

	grandTotal := calculateGrandTotal(input)
	fmt.Printf("%s: %d\n", filename, grandTotal)
}

func executePartTwo(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	operationRow := input[len(input)-1]
	boundaries := findProblemBoundaries(operationRow)

	grandTotal := 0

	for i := len(boundaries) - 1; i >= 0; i-- {
		startCol := boundaries[i]
		var endCol int
		if i < len(boundaries)-1 {
			endCol = boundaries[i+1]
		} else {
			endCol = len(operationRow)
		}

		var problemLines []string
		for j := 0; j < len(input)-1; j++ { // Exclude operation row
			if startCol < len(input[j]) {
				line := input[j][startCol:min(endCol, len(input[j]))]
				line = strings.TrimRight(line, " ")
				if line != "" {
					problemLines = append(problemLines, line)
				}
			}
		}

		operation := string(operationRow[startCol])
		result := solveSingleProblemVertical(problemLines, operation)
		grandTotal += result
	}

	fmt.Printf("%s: %d\n", filename, grandTotal)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findProblemBoundaries(operationRow string) []int {
	var boundaries []int
	inProblem := false

	for i := 0; i < len(operationRow); i++ {
		if operationRow[i] != ' ' {
			if !inProblem {
				boundaries = append(boundaries, i)
				inProblem = true
			}
		} else {
			inProblem = false
		}
	}

	return boundaries
}

func solveSingleProblemVertical(problemLines []string, operation string) int {
	maxLen := 0
	for _, line := range problemLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	var numbers []int

	for col := maxLen - 1; col >= 0; col-- {
		numStr := ""

		for _, line := range problemLines {
			if col < len(line) && line[col] != ' ' {
				numStr += string(line[col])
			}
		}

		if numStr != "" {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
	}

	return evaluateProblem(numbers, operation)
}

func calculateGrandTotal(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	var tokenLines [][]string
	for _, line := range lines {
		tokens := strings.Fields(line)
		tokenLines = append(tokenLines, tokens)
	}

	if len(tokenLines) == 0 || len(tokenLines[0]) == 0 {
		return 0
	}

	numProblems := len(tokenLines[0])
	grandTotal := 0

	for col := 0; col < numProblems; col++ {
		var numbers []int
		var operation string

		for row := 0; row < len(tokenLines); row++ {
			if col >= len(tokenLines[row]) {
				continue
			}

			token := tokenLines[row][col]

			if token == "+" || token == "*" {
				operation = token
			} else {
				num, err := strconv.Atoi(token)
				if err == nil {
					numbers = append(numbers, num)
				}
			}
		}

		if len(numbers) > 0 && operation != "" {
			result := evaluateProblem(numbers, operation)
			grandTotal += result
		}
	}

	return grandTotal
}

func evaluateProblem(numbers []int, operation string) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if operation == "+" {
			result += numbers[i]
		} else if operation == "*" {
			result *= numbers[i]
		}
	}

	return result
}
