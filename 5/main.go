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

	ranges, ids := parseInput(input)
	freshCount := countFreshIngredients(ranges, ids)
	fmt.Printf("%s: %d\n", filename, freshCount)
}

func executePartTwo(filename string) {
	input := utilities.LoadInput(filename)
	if len(input) == 0 {
		fmt.Println("No input found")
		return
	}

	ranges, _ := parseInput(input)
	totalFresh := countTotalFreshIDs(ranges)
	fmt.Printf("%s: %d\n", filename, totalFresh)
}

func countTotalFreshIDs(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start position
	sortedRanges := make([]Range, len(ranges))
	copy(sortedRanges, ranges)

	// Simple bubble sort (good enough for small inputs)
	for i := 0; i < len(sortedRanges); i++ {
		for j := i + 1; j < len(sortedRanges); j++ {
			if sortedRanges[j].start < sortedRanges[i].start {
				sortedRanges[i], sortedRanges[j] = sortedRanges[j], sortedRanges[i]
			}
		}
	}

	// Merge overlapping ranges and count total IDs
	totalCount := 0
	currentStart := sortedRanges[0].start
	currentEnd := sortedRanges[0].end

	for i := 1; i < len(sortedRanges); i++ {
		if sortedRanges[i].start <= currentEnd+1 {
			// Ranges overlap or are adjacent, merge them
			if sortedRanges[i].end > currentEnd {
				currentEnd = sortedRanges[i].end
			}
		} else {
			// No overlap, count the current range and start a new one
			totalCount += currentEnd - currentStart + 1
			currentStart = sortedRanges[i].start
			currentEnd = sortedRanges[i].end
		}
	}

	// Add the last range
	totalCount += currentEnd - currentStart + 1

	return totalCount
}

type Range struct {
	start int
	end   int
}

func parseInput(lines []string) ([]Range, []int) {
	var ranges []Range
	var ids []int
	parsingRanges := true

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				start, _ := strconv.Atoi(parts[0])
				end, _ := strconv.Atoi(parts[1])
				ranges = append(ranges, Range{start: start, end: end})
			}
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	return ranges, ids
}

func countFreshIngredients(ranges []Range, ids []int) int {
	count := 0

	for _, id := range ids {
		if isFresh(id, ranges) {
			count++
		}
	}

	return count
}

func isFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}
