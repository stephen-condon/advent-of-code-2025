package main

import (
	"fmt"
	"strconv"
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
	input := LoadInput(filename)
	position := 50
	numZeroes := 0

	for _, line := range input {
		command := parseCommand(line)

		switch command.Direction {
		case "L":
			position = (position - command.Steps + 100) % 100
		case "R":
			position = (position + command.Steps) % 100
		}

		if position == 0 {
			numZeroes++
		}
	}

	fmt.Println(numZeroes)
}

func executePartTwo(filename string) {
	input := LoadInput(filename)
	position := 50
	numZeroes := 0

	for _, line := range input {
		command := parseCommand(line)

		switch command.Direction {
		case "L":
			for i := 0; i < command.Steps; i++ {
				position = (position - 1 + 100) % 100
				if position == 0 {
					numZeroes++
				}
			}

		case "R":
			for i := 0; i < command.Steps; i++ {
				position = (position + 1) % 100
				if position == 0 {
					numZeroes++
				}
			}
		}
	}

	fmt.Println(numZeroes)
}

type Command struct {
	Direction string
	Steps     int
}

func parseCommand(command string) *Command {
	steps, _ := strconv.Atoi(command[1:])
	return &Command{
		Direction: string(command[0]),
		Steps:     steps,
	}
}
