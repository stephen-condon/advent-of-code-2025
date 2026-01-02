package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stephen-condon/advent-of-code-2025/utilities"
)

// Example input: 11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
// The ranges are separated by commas (,); each range gives its first ID and last ID separated by a dash (-).
// you can find the invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice. So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.
// None of the numbers have leading zeroes; 0101 isn't an ID at all. (101 is a valid ID that you would ignore.)

/*
	Example results:

	11-22 has two invalid IDs, 11 and 22.
	95-115 has one invalid ID, 99.
	998-1012 has one invalid ID, 1010.
	1188511880-1188511890 has one invalid ID, 1188511885.
	222220-222224 has one invalid ID, 222222.
	1698522-1698528 contains no invalid IDs.
	446443-446449 has one invalid ID, 446446.
	38593856-38593862 has one invalid ID, 38593859.

	The rest of the ranges contain no invalid IDs.

	Adding up all the invalid IDs in this example produces 1227775554.
*/

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

	// Parse the ranges from the first line
	ranges := parseRanges(input[0])

	totalInvalidSum := 0

	for _, r := range ranges {
		invalidIDs := findInvalidIDsInRange(r.start, r.end)
		for _, id := range invalidIDs {
			totalInvalidSum += id
		}
	}

	fmt.Printf("%s: %d\n", filename, totalInvalidSum)
}

// implement a executePartTwo method, with the following changes to assumptions (do not modify any code used for part one)
/*
	Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice. So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times), and 1111111 (1 seven times) are all invalid IDs.

	From the same example as before:

	11-22 still has two invalid IDs, 11 and 22.
	95-115 now has two invalid IDs, 99 and 111.
	998-1012 now has two invalid IDs, 999 and 1010.
	1188511880-1188511890 still has one invalid ID, 1188511885.
	222220-222224 still has one invalid ID, 222222.
	1698522-1698528 still contains no invalid IDs.
	446443-446449 still has one invalid ID, 446446.
	38593856-38593862 still has one invalid ID, 38593859.
	565653-565659 now has one invalid ID, 565656.
	824824821-824824827 now has one invalid ID, 824824824.
	2121212118-2121212124 now has one invalid ID, 2121212121.
	Adding up all the invalid IDs in this example produces 4174379265.
*/

func executePartTwo(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	// Parse the ranges from the first line
	ranges := parseRanges(input[0])

	totalInvalidSum := 0

	for _, r := range ranges {
		invalidIDs := findInvalidIDsInRangePartTwo(r.start, r.end)
		for _, id := range invalidIDs {
			totalInvalidSum += id
		}
	}

	fmt.Printf("%s: %d\n", filename, totalInvalidSum)
}

func findInvalidIDsInRangePartTwo(start, end int) []int {
	var invalidIDs []int

	for id := start; id <= end; id++ {
		if isInvalidIDPartTwo(id) {
			invalidIDs = append(invalidIDs, id)
		}
	}

	return invalidIDs
}

// isInvalidIDPartTwo checks if a number is made of a sequence repeated at least twice
// Examples: 11 (1 two times), 111 (1 three times), 12341234 (1234 two times), 123123123 (123 three times)
func isInvalidIDPartTwo(id int) bool {
	str := strconv.Itoa(id)
	length := len(str)

	// Try all possible pattern lengths from 1 to length/2
	// The pattern must repeat at least twice, so max pattern length is length/2
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		// Check if the string length is divisible by the pattern length
		if length%patternLen == 0 {
			pattern := str[:patternLen]
			numRepeats := length / patternLen

			// Check if the entire string is made of this pattern repeated
			if numRepeats >= 2 {
				isValid := true
				for i := 0; i < numRepeats; i++ {
					start := i * patternLen
					end := start + patternLen
					if str[start:end] != pattern {
						isValid = false
						break
					}
				}
				if isValid {
					return true
				}
			}
		}
	}

	return false
}

type Range struct {
	start int
	end   int
}

func parseRanges(line string) []Range {
	var ranges []Range
	parts := strings.Split(line, ",")

	for _, part := range parts {
		rangeParts := strings.Split(part, "-")
		if len(rangeParts) == 2 {
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			ranges = append(ranges, Range{start: start, end: end})
		}
	}

	return ranges
}

func findInvalidIDsInRange(start, end int) []int {
	var invalidIDs []int

	for id := start; id <= end; id++ {
		if isInvalidID(id) {
			invalidIDs = append(invalidIDs, id)
		}
	}

	return invalidIDs
}

// isInvalidID checks if a number is made of a sequence repeated twice
// Examples: 11 (1 repeated), 6464 (64 repeated), 123123 (123 repeated)
func isInvalidID(id int) bool {
	str := strconv.Itoa(id)
	length := len(str)

	// The string must have even length to be repeated twice
	if length%2 != 0 {
		return false
	}

	halfLength := length / 2
	firstHalf := str[:halfLength]
	secondHalf := str[halfLength:]

	return firstHalf == secondHalf
}
