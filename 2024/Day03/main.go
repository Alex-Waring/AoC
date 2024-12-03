package main

import (
	"fmt"
	"regexp"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	total := 0

	for _, line := range input {
		r := regexp.MustCompile(`mul\(\d*,\d*\)`)
		matches := r.FindAllString(line, -1)

		for _, match := range matches {
			num_r := regexp.MustCompile(`\d+`)
			nums := num_r.FindAllString(match, -1)
			total += utils.IntegerOf(nums[0]) * utils.IntegerOf(nums[1])
		}
	}

	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	total := 0
	matches := []string{}
	r := regexp.MustCompile(`mul\(\d*,\d*\)|do\(\)|don't\(\)`)

	for _, line := range input {
		matches = append(matches, r.FindAllString(line, -1)...)
	}

	enabled := true

	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		num_r := regexp.MustCompile(`\d+`)
		nums := num_r.FindAllString(match, -1)
		total += utils.IntegerOf(nums[0]) * utils.IntegerOf(nums[1])
	}

	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
