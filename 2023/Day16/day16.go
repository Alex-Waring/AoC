package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Location struct {
	energized bool
	mirror    string
}

func part1(grid map[int]map[int]Location, start_y int, start_x int, start_dy int, start_dx int) int {
	beams := [][]int{{start_y, start_x, start_dy, start_dx}}

	seen := []string{}

	for len(beams) > 0 {
		beam := beams[0]
		y, x, dy, dx := beam[0], beam[1], beam[2], beam[3]

		xx, yy := x+dx, y+dy

		key := fmt.Sprintf("%d,%d,%d,%d", yy, xx, dy, dx)

		// if we've seen it continue
		if utils.StringInSlice(key, seen) {
			beams = beams[1:]
			continue
		}

		// if invalid coords continue
		if !((0 <= yy && yy < len(grid)) && (0 <= xx && xx < len(grid[0]))) {
			beams = beams[1:]
			continue
		}

		seen = append(seen, key)

		switch tile := grid[yy][xx].mirror; tile {
		case "/":
			dx, dy = -dy, -dx
		case `\`:
			dx, dy = dy, dx
		case "|":
			if dx == 1 || dx == -1 {
				dx, dy = 0, 1
				beams = append(beams, []int{yy, xx, -1, 0})
			}
		case "-":
			if dy == 1 || dy == -1 {
				dx, dy = 1, 0
				beams = append(beams, []int{yy, xx, 0, -1})
			}
		}

		beams = append(beams, []int{yy, xx, dy, dx})

		beams = beams[1:]
	}

	for _, value := range seen {
		x, y := strings.Split(value, ",")[1], strings.Split(value, ",")[0]

		location := grid[utils.IntegerOf(y)][utils.IntegerOf(x)]
		location.energized = true

		grid[utils.IntegerOf(y)][utils.IntegerOf(x)] = location
	}

	total := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x].energized {
				total++
			}
		}
	}
	return total
}

func main() {
	input := utils.ReadInput("input.txt")

	grid := map[int]map[int]Location{}

	for y, line := range input {
		for x, charecter := range line {
			if _, ok := grid[y]; !ok {
				grid[y] = make(map[int]Location)
			}
			grid[y][x] = Location{mirror: string(charecter)}
		}
	}

	fmt.Println(part1(grid, 0, -1, 0, 1))

	part2_values := []int{}

	for y := 0; y < len(grid); y++ {
		part2_values = append(part2_values, part1(grid, y, -1, 0, 1))
		part2_values = append(part2_values, part1(grid, y, len(grid[0]), 0, -1))
	}
	for x := 0; x < len(grid[0]); x++ {
		part2_values = append(part2_values, part1(grid, -1, x, 1, 0))
		part2_values = append(part2_values, part1(grid, len(grid), x, -1, 0))
	}

	fmt.Println(slices.Max(part2_values))
}
