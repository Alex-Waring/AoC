package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Mapper map[string]int

type InverseMapper map[int]string

func part1(input []string) {
	defer utils.Timer("part1")()
	total := 0
	for _, line := range input {
		outputs := strings.Split(strings.Split(line, "|")[1], " ")
		for _, o := range outputs {
			if utils.IntInSlice(len(o), []int{2, 3, 4, 7}) {
				total++
			}
		}
	}
	fmt.Println("Total: ", total)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	results := []string{}
	for _, line := range input {
		outputs := strings.Fields(strings.Split(line, "|")[1])
		inputs := strings.Fields(strings.Split(line, "|")[0])

		result := ""

		mapper := solve(inputs)
		for _, o := range outputs {
			o = utils.SortStringAlphabetically(o)
			result += fmt.Sprint(mapper[o])
		}
		results = append(results, result)
	}

	total := 0
	for _, r := range results {
		total += utils.IntegerOf(r)
	}
	fmt.Println("Total: ", total)
}

func solve(inputs []string) Mapper {
	mapper := Mapper{}
	inverseMapper := InverseMapper{}
	// Grab the easy ones, 2 digits is a 1, 3 digits is a 7, 4 digits is a 4, 7 digits is an 8
	for _, input := range inputs {
		if len(input) == 2 {
			mapper[utils.SortStringAlphabetically(input)] = 1
			inverseMapper[1] = input
		} else if len(input) == 3 {
			mapper[utils.SortStringAlphabetically(input)] = 7
			inverseMapper[7] = input
		} else if len(input) == 4 {
			mapper[utils.SortStringAlphabetically(input)] = 4
			inverseMapper[4] = input
		} else if len(input) == 7 {
			mapper[utils.SortStringAlphabetically(input)] = 8
			inverseMapper[8] = input
		}
	}

	// 3 is a 5 digit long input with 1 totally inside
	for _, input := range inputs {
		if len(input) == 5 && numberInString(input, inverseMapper[1]) {
			mapper[utils.SortStringAlphabetically(input)] = 3
			inverseMapper[3] = input
		}
	}

	// 6 digits are 0, 6, 9.
	for _, input := range inputs {
		if len(input) == 6 {
			// if 7 is in it then it's a 9 or 0
			if numberInString(input, inverseMapper[7]) {
				// If 4 inside then it;s a 9, otherwise it's a 0
				if numberInString(input, inverseMapper[4]) {
					mapper[utils.SortStringAlphabetically(input)] = 9
					inverseMapper[9] = input
				} else {
					mapper[utils.SortStringAlphabetically(input)] = 0
					inverseMapper[0] = input
				}
			} else {
				// Otherwise it's a 6
				mapper[utils.SortStringAlphabetically(input)] = 6
				inverseMapper[6] = input
			}
		}
	}

	// We are left with 2 and 5, both of which have 5 digits
	// Compare them to a 6, if one off it's a 5, if two off it's a 2
	for _, input := range inputs {
		_, exists := mapper[utils.SortStringAlphabetically(input)]
		if exists {
			continue
		} else {
			// Don't strictly need to check for 5 digits, but it's a good sanity check
			if len(input) == 5 {
				if oneOffNumber(input, inverseMapper[6]) {
					mapper[utils.SortStringAlphabetically(input)] = 5
					inverseMapper[5] = input
				} else {
					mapper[utils.SortStringAlphabetically(input)] = 2
					inverseMapper[2] = input
				}
			}
		}
	}

	return mapper
}

func numberInString(s string, n string) bool {
	for _, rune := range n {
		if !strings.Contains(s, string(rune)) {
			return false
		}
	}
	return true
}

func oneOffNumber(s string, n string) bool {
	diff := 0
	for _, rune := range n {
		if !strings.Contains(s, string(rune)) {
			diff++
		}
	}
	return diff == 1
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
