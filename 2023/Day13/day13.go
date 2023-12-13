package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func verticalReflection(input []string, max_differences int) (bool, int) {
	input = utils.RemoveSliceSpaces(input)
	for vertical := 1; vertical < len(input[0]); vertical++ {
		differences := 0
		reflection := true

		left_side := utils.MakeRange(0, vertical-1)
		right_side := utils.MakeRange(vertical, len(input[0])-1)

		length := min(len(left_side), len(right_side))

		left_side = utils.MakeRange(vertical-length, vertical-1)
		right_side = right_side[0:length]

		for row := 0; reflection && row < len(input); row++ {
			for i := 0; i < length; i++ {
				if input[row][right_side[i]] != input[row][left_side[len(left_side)-i-1]] {
					differences++
					if differences > max_differences {
						reflection = false
					}
				}
			}
		}
		if reflection && differences == max_differences {
			return true, vertical
		}
	}
	return false, 0
}

func horizontalReflection(input []string, max_differences int) (bool, int) {
	input = utils.RemoveSliceSpaces(input)
	rotated_input := make([]string, len(input[0]))

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			rotated_input[x] += string(input[y][x])
		}
	}

	return verticalReflection(rotated_input, max_differences)
}

func part1(puzzles []string) {
	defer utils.Timer("part1")()
	total := 0
	for _, puzzle := range puzzles {
		if ok, vertical_symmetry := verticalReflection(strings.Split(puzzle, "\n"), 0); ok {
			total += vertical_symmetry
		} else {
			_, horizontal_symmetry := horizontalReflection(strings.Split(puzzle, "\n"), 0)
			total += horizontal_symmetry * 100
		}
	}
	fmt.Println(total)
}

func part2(puzzles []string) {
	defer utils.Timer("part2")()
	total := 0
	for _, puzzle := range puzzles {
		if ok, vertical_symmetry := verticalReflection(strings.Split(puzzle, "\n"), 1); ok {
			total += vertical_symmetry
		} else {
			_, horizontal_symmetry := horizontalReflection(strings.Split(puzzle, "\n"), 1)
			total += horizontal_symmetry * 100
		}
	}
	fmt.Println(total)
}

func main() {
	raw_input, err := os.ReadFile("input.txt")
	utils.Check(err)

	puzzles := strings.Split(string(raw_input), "\n\n")
	part1(puzzles)
	part2(puzzles)
}
