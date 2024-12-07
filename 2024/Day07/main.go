package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	total := 0

	for _, line := range input {
		test_result := utils.IntegerOf(strings.Split(line, ":")[0])
		inputs := utils.StringToIntList(strings.Fields(strings.Split(line, ":")[1]))

		if recursive_check(inputs, test_result, false) {
			total += test_result
		}
	}
	fmt.Println(total)
}

func recursive_check(inputs []int, target int, check_concat bool) bool {
	if len(inputs) == 1 {
		return inputs[0] == target
	}
	if recursive_check(inputs[:len(inputs)-1], target-inputs[len(inputs)-1], check_concat) {
		return true
	}
	if utils.IsDivisible(target, inputs[len(inputs)-1]) {
		if recursive_check(inputs[:len(inputs)-1], target/inputs[len(inputs)-1], check_concat) {
			return true
		}
	}
	if check_concat && utils.EndsWith(target, inputs[len(inputs)-1]) {
		// Avoid an edge case where we are trying to remove the target from the inputs
		// but we still have more numbers to check
		if inputs[len(inputs)-1] == target {
			return false
		}
		if recursive_check(inputs[:len(inputs)-1], utils.RemoveSuffix(target, inputs[len(inputs)-1]), true) {
			return true
		}

	}
	return false
}

func part2(input []string) {
	defer utils.Timer("part2")()

	total := 0

	for _, line := range input {
		test_result := utils.IntegerOf(strings.Split(line, ":")[0])
		inputs := utils.StringToIntList(strings.Fields(strings.Split(line, ":")[1]))

		if recursive_check(inputs, test_result, true) {
			total += test_result
		}
	}
	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
