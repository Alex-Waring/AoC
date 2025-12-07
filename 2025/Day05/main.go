package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	freshRanges := [][2]int{}
	freshIDs := map[int]bool{}

	for _, line := range input {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start := utils.IntegerOf(parts[0])
			end := utils.IntegerOf(parts[1])
			freshRanges = append(freshRanges, [2]int{start, end})
		} else if line != "" {
			id := utils.IntegerOf(line)
			for _, r := range freshRanges {
				if id >= r[0] && id <= r[1] {
					freshIDs[id] = true
				}
			}
		}
	}
	fmt.Println(len(freshIDs))
}

func part2(input []string) {
	defer utils.Timer("part1")()
	freshRanges := [][2]int{}

	for _, line := range input {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start := utils.IntegerOf(parts[0])
			end := utils.IntegerOf(parts[1])
			freshRanges = append(freshRanges, [2]int{start, end})
		}
	}

	// Sort the slices by start value
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	// Merge overlapping ranges
	mergedRanges := [][2]int{}
	currentRange := freshRanges[0]

	for i := 1; i < len(freshRanges); i++ {
		if freshRanges[i][0] <= currentRange[1]+1 {
			if freshRanges[i][1] > currentRange[1] {
				currentRange[1] = freshRanges[i][1]
			}
		} else {
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = freshRanges[i]
		}
	}
	mergedRanges = append(mergedRanges, currentRange)

	totalFresh := 0
	for _, r := range mergedRanges {
		totalFresh += r[1] - r[0] + 1
	}
	fmt.Println(totalFresh)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
