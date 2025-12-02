package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	inputStr := input[0]
	invalidTotal := 0
	for _, line := range strings.Split(inputStr, ",") {
		start, end := utils.IntegerOf(strings.Split(line, "-")[0]), utils.IntegerOf(strings.Split(line, "-")[1])

		for i := start; i <= end; i++ {
			str := fmt.Sprintf("%d", i)
			if len(str)%2 != 0 {
				continue
			}
			mid := len(str) / 2
			if str[:mid] == str[mid:] {
				invalidTotal += i
			}
		}
	}
	println(invalidTotal)
}

func sliceAllEmpty(s []string) bool {
	for _, str := range s {
		if str != "" {
			return false
		}
	}
	return true
}

func part2(input []string) {
	defer utils.Timer("part2")()
	inputStr := input[0]
	invalidTotal := 0
	for _, line := range strings.Split(inputStr, ",") {
		start, end := utils.IntegerOf(strings.Split(line, "-")[0]), utils.IntegerOf(strings.Split(line, "-")[1])

	pattern_loop:
		for i := start; i <= end; i++ {
			str := fmt.Sprintf("%d", i)
			pattern := ""
			for _, char := range str {
				pattern += string(char)
				// Split it by the pattern, if we've found a match then the len will be > 2
				// and all the parts will be ""
				patternSplit := strings.Split(str, pattern)
				if len(patternSplit) > 2 && sliceAllEmpty(patternSplit) {
					fmt.Println(i)
					invalidTotal += i
					continue pattern_loop
				}
			}
		}
	}
	println(invalidTotal)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
