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

	totalSum := 0

	for _, line := range input {
		maxNumber := findLargestTwoDigitNumber(line)
		totalSum += maxNumber
	}

	fmt.Printf("%s: %d\n", filename, totalSum)
}

func executePartTwo(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	totalSum := 0

	for _, line := range input {
		maxNumber := findLargestTwelveDigitNumber(line)
		totalSum += maxNumber
	}

	fmt.Printf("%s: %d\n", filename, totalSum)
}

func findLargestTwelveDigitNumber(line string) int {
	var digits []int
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			digits = append(digits, int(line[i]-'0'))
		}
	}

	if len(digits) < 12 {
		return 0
	}

	result := make([]int, 12)
	startPos := 0

	for pos := 0; pos < 12; pos++ {
		digitsNeeded := 12 - pos
		maxDigit := -1
		maxDigitPos := -1

		searchEnd := len(digits) - digitsNeeded + 1
		for i := startPos; i < searchEnd; i++ {
			if digits[i] > maxDigit {
				maxDigit = digits[i]
				maxDigitPos = i
			}

			if maxDigit == 9 {
				break
			}
		}

		result[pos] = maxDigit
		startPos = maxDigitPos + 1
	}
	number := 0
	for i := 0; i < 12; i++ {
		number = number*10 + result[i]
	}

	return number
}

func findLargestTwoDigitNumber(line string) int {
	maxNumber := 0

	for i := 0; i < len(line); i++ {
		if line[i] < '0' || line[i] > '9' {
			continue
		}

		for j := i + 1; j < len(line); j++ {
			if line[j] < '0' || line[j] > '9' {
				continue
			}

			firstDigit := int(line[i] - '0')
			secondDigit := int(line[j] - '0')
			number := firstDigit*10 + secondDigit

			if number > maxNumber {
				maxNumber = number
			}
		}
	}

	return maxNumber
}
