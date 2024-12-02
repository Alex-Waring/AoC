package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	safe := 0

	for _, line := range input {
		numbers := strings.Fields(line)
		line_safe := lineSafe(numbers)

		if line_safe {
			safe++
		}
	}
	fmt.Println(safe)
}

func lineSafe(numbers []string) bool {
	diffs := []int{}

	line_safe := true

	for i := 1; i < len(numbers); i++ {
		a := utils.IntegerOf(numbers[i-1])
		b := utils.IntegerOf(numbers[i])
		diffs = append(diffs, a-b)
	}

	increasing := false
	if diffs[0] > 0 {
		increasing = true
	}

	for i := 0; i < len(diffs); i++ {
		if utils.Abs(diffs[i]) > 3 || diffs[i] == 0 {
			line_safe = false
			break
		}
		if diffs[i] > 0 && !increasing {
			line_safe = false
			break
		}
		if diffs[i] < 0 && increasing {
			line_safe = false
			break
		}
	}
	return line_safe
}

func part2(input []string) {
	defer utils.Timer("part2")()

	safe := 0

	for _, line := range input {
		numbers := strings.Fields(line)
		line_safe := lineSafe(numbers)

		for i := 0; i < len(numbers); i++ {
			numbersCopy := make([]string, len(numbers))
			copy(numbersCopy, numbers)
			numbersCopy = utils.Remove(numbersCopy, i)
			pop_safe := lineSafe(numbersCopy)
			if pop_safe {
				line_safe = true
				break
			}
		}

		if line_safe {
			safe++
		}
	}
	fmt.Println(safe)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
