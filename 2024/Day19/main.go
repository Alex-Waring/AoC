package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	patterns := strings.Fields(input[0])
	parsed_patterns := []string{}

	for _, pattern := range patterns {
		parsed_patterns = append(parsed_patterns, strings.TrimSuffix(pattern, ","))
	}

	designs := input[2:]

	cache := map[string]int{}

	var solutions func(string) int
	solutions = func(design string) int {
		var n int
		if n, ok := cache[design]; ok {
			return n
		}

		if design == "" {
			return 1
		}
		for _, pattern := range parsed_patterns {
			if strings.HasPrefix(design, pattern) {
				n += solutions(design[len(pattern):])
			}
		}
		cache[design] = n
		return n
	}

	possible_patterns := 0

	for _, design := range designs {
		ways_to_make := solutions(design)
		if ways_to_make >= 1 {
			possible_patterns++
		}
	}
	fmt.Println(possible_patterns)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	patterns := strings.Fields(input[0])
	parsed_patterns := []string{}

	for _, pattern := range patterns {
		parsed_patterns = append(parsed_patterns, strings.TrimSuffix(pattern, ","))
	}

	designs := input[2:]

	cache := map[string]int{}

	var solutions func(string) int
	solutions = func(design string) int {
		var n int
		if n, ok := cache[design]; ok {
			return n
		}

		if design == "" {
			return 1
		}
		for _, s := range parsed_patterns {
			if strings.HasPrefix(design, s) {
				n += solutions(design[len(s):])
			}
		}
		cache[design] = n
		return n
	}

	total_possible_patterns := 0

	for _, design := range designs {
		ways_to_make := solutions(design)
		if ways_to_make >= 1 {
			total_possible_patterns += ways_to_make
		}
	}
	fmt.Println(total_possible_patterns)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
