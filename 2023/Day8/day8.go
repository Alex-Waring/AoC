package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type instruction struct {
	left  string
	right string
}

func part1(directions string, instructions map[string]instruction, start_location string) int {
	defer utils.Timer("part1")
	loops := 0
	for start_location != "ZZZ" {
		if directions[loops%len(directions)] == 'R' {
			start_location = instructions[start_location].right
		} else {
			start_location = instructions[start_location].left
		}
		loops++
	}
	return loops
}

func endsWith(location string, letter string) bool {
	regex := "[A-Z][A-Z]" + letter
	found, _ := regexp.MatchString(regex, location)
	return found
}

func main() {
	defer utils.Timer("main")
	lines := utils.ReadInput("input.txt")
	directions := lines[0]
	raw_instructions := utils.RemoveSliceSpaces(lines[1:])

	instructions := map[string]instruction{}

	for _, line := range raw_instructions {
		key := strings.Fields(line)[0]
		left := strings.TrimPrefix(strings.TrimSuffix(strings.Fields(line)[2], ","), "(")
		right := strings.TrimSuffix(strings.Fields(line)[3], ")")

		instructions[key] = instruction{left: left, right: right}
	}
	current_location := "AAA"
	loops := part1(directions, instructions, current_location)
	fmt.Println(loops)

	ghost_starts := []string{}
	steps := []int{}
	for key, _ := range instructions {
		if endsWith(key, "A") {
			ghost_starts = append(ghost_starts, key)
		}
	}

	for _, start := range ghost_starts {
		loops = 0
		for !endsWith(start, "Z") {
			if directions[loops%len(directions)] == 'R' {
				start = instructions[start].right
			} else {
				start = instructions[start].left
			}
			loops++
		}
		steps = append(steps, loops)
	}
	part2 := utils.LCM(steps[0], steps[1], steps[2:]...)
	fmt.Println(part2)
}
