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
