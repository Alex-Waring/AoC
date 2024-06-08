package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Grid map[utils.Position]int

func part1(input []string) {
	defer utils.Timer("part1")()
	grid := Grid{}

	for _, line := range input {
		start := strings.Split(line, " ")[0]
		end := strings.Split(line, " ")[2]

		x1 := utils.IntegerOf(strings.Split(start, ",")[0])
		x2 := utils.IntegerOf(strings.Split(end, ",")[0])
		y1 := utils.IntegerOf(strings.Split(start, ",")[1])
		y2 := utils.IntegerOf(strings.Split(end, ",")[1])

		if isLine(x1, x2, y1, y2) {
			for _, y := range utils.MakeRange(y1, y2) {
				for _, x := range utils.MakeRange(x1, x2) {
					grid[utils.Position{Row: y, Col: x}]++
				}
			}
		}
	}

	count := 0
	for _, v := range grid {
		if v > 1 {
			count++
		}
	}
	fmt.Println("Overlapping squares: ", count)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	grid := Grid{}

	for _, line := range input {
		start := strings.Split(line, " ")[0]
		end := strings.Split(line, " ")[2]

		x1 := utils.IntegerOf(strings.Split(start, ",")[0])
		x2 := utils.IntegerOf(strings.Split(end, ",")[0])
		y1 := utils.IntegerOf(strings.Split(start, ",")[1])
		y2 := utils.IntegerOf(strings.Split(end, ",")[1])

		if isLine(x1, x2, y1, y2) {
			for _, y := range utils.MakeRange(y1, y2) {
				for _, x := range utils.MakeRange(x1, x2) {
					grid[utils.Position{Row: y, Col: x}]++
				}
			}
		} else {
			x_range := makeOrderedRange(x1, x2)
			y_range := makeOrderedRange(y1, y2)

			if len(x_range) != len(y_range) {
				panic("line not at 45 degrees")
			}

			for i := 0; i < len(x_range); i++ {
				grid[utils.Position{Row: y_range[i], Col: x_range[i]}]++
			}
		}
	}

	count := 0
	for _, v := range grid {
		if v > 1 {
			count++
		}
	}
	fmt.Println("Overlapping squares: ", count)
}

func makeOrderedRange(a int, b int) []int {
	if a > b {
		return utils.MakeRange(a, b)
	}
	r := utils.MakeRange(b, a)
	slices.Reverse(r)
	return r

}

func isLine(x1 int, x2 int, y1 int, y2 int) bool {
	return x1 == x2 || y1 == y2
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
