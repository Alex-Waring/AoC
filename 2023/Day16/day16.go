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

func part1(grid map[utils.Position]Location, start_y int, start_x int, start_dy int, start_dx int) int {
	q := utils.Queue[[]int]{}
	q.Push([]int{start_y, start_x, start_dy, start_dx})

	seen := []string{}

	for !q.IsEmpty() {
		beam := q.Pop()
		y, x, dy, dx := beam[0], beam[1], beam[2], beam[3]

		xx, yy := x+dx, y+dy

		key := fmt.Sprintf("%d,%d,%d,%d", yy, xx, dy, dx)

		// if we've seen it continue
		if utils.StringInSlice(key, seen) {
			continue
		}

		// if invalid coords continue
		position := utils.NewPosition(yy, xx)
		if _, exists := grid[position]; !exists {
			continue
		}

		seen = append(seen, key)

		switch tile := grid[position].mirror; tile {
		case "/":
			dx, dy = -dy, -dx
		case `\`:
			dx, dy = dy, dx
		case "|":
			if dx == 1 || dx == -1 {
				dx, dy = 0, 1
				q.Push([]int{yy, xx, -1, 0})
			}
		case "-":
			if dy == 1 || dy == -1 {
				dx, dy = 1, 0
				q.Push([]int{yy, xx, 0, -1})
			}
		}

		q.Push([]int{yy, xx, dy, dx})
	}

	for _, value := range seen {
		x, y := strings.Split(value, ",")[1], strings.Split(value, ",")[0]

		position := grid[utils.NewPosition(utils.IntegerOf(y), utils.IntegerOf(x))]
		position.energized = true

		grid[utils.NewPosition(utils.IntegerOf(y), utils.IntegerOf(x))] = position
	}

	total := 0

	for _, position := range grid {
		if position.energized {
			total++
		}
	}
	return total
}

func main() {
	input := utils.ReadInput("input.txt")

	grid := map[utils.Position]Location{}

	for y, line := range input {
		for x, charecter := range line {
			poition := utils.NewPosition(y, x)
			grid[poition] = Location{mirror: string(charecter)}
		}
	}

	fmt.Println(part1(grid, 0, -1, 0, 1))

	part2_values := []int{}

	no_of_rows := len(input)
	no_of_cols := len(input[0])

	for y := 0; y <= no_of_rows; y++ {
		part2_values = append(part2_values, part1(grid, y, -1, 0, 1))
		part2_values = append(part2_values, part1(grid, y, no_of_cols, 0, -1))
	}
	for x := 0; x <= no_of_cols; x++ {
		part2_values = append(part2_values, part1(grid, -1, x, 1, 0))
		part2_values = append(part2_values, part1(grid, no_of_rows, x, -1, 0))
	}

	fmt.Println(slices.Max(part2_values))
}
