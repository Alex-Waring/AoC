package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	// Store the rules in a map that says the key must be after the value
	after := map[int][]int{}
	to_print := [][]int{}

	after, to_print = parseinput(input)

	ok_rules := [][]int{}

out:
	for _, print := range to_print {
		// for every number, check if there are any values after it in the before map,
		// which means it's wrong
		for i := 0; i < len(print); i++ {
			if _, ok := after[print[i]]; !ok {
				continue
			}
			for j := i + 1; j < len(print); j++ {
				numbers_to_check := after[print[i]]
				for _, num := range numbers_to_check {
					if print[j] == num {
						// A value is after that should be before
						continue out
					}
				}
			}
		}
		ok_rules = append(ok_rules, print)
	}

	middle := 0
	for _, rule := range ok_rules {
		middle_indles := (len(rule) - 1) / 2
		middle += rule[middle_indles]
	}

	fmt.Println(middle)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	// Store the rules in a map that says the key must be after the value
	after := map[int][]int{}
	to_print := [][]int{}

	after, to_print = parseinput(input)

	bad_rules := [][]int{}
out:
	for _, print := range to_print {
		// for every number, check if there are any values after it in the before map,
		// which means it's wrong
		for i := 0; i < len(print); i++ {
			if _, ok := after[print[i]]; !ok {
				continue
			}
			for j := i + 1; j < len(print); j++ {
				numbers_to_check := after[print[i]]
				for _, num := range numbers_to_check {
					if print[j] == num {
						// A value is after that should be before
						bad_rules = append(bad_rules, print)
						continue out
					}
				}
			}
		}
	}

	middle := 0
	for _, rule := range bad_rules {
		sort.Slice(rule, func(i, j int) bool {
			first := rule[i]
			second := rule[j]
			if _, ok := after[second]; !ok {
				return false
			}
			number_that_must_come_after := after[second]
			if !(utils.IntListContains(number_that_must_come_after, first)) {
				return false
			}
			return true
		})
		middle_indles := (len(rule) - 1) / 2
		middle += rule[middle_indles]
	}

	fmt.Println(middle)
}

func parseinput(input []string) (map[int][]int, [][]int) {
	after := map[int][]int{}
	to_print := [][]int{}

	for _, line := range input {
		if strings.Contains(line, "|") {
			numbers := strings.Split(line, "|")
			after[utils.IntegerOf(numbers[1])] = append(after[utils.IntegerOf(numbers[1])], utils.IntegerOf(numbers[0]))
		} else if strings.Contains(line, ",") {
			numbers := strings.Split(line, ",")
			print := []int{}
			for _, num := range numbers {
				print = append(print, utils.IntegerOf(num))
			}
			to_print = append(to_print, print)
		}
	}
	return after, to_print
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
