package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Location struct {
	land     string
	visited  bool
	finished bool
}

func set_grid(grid map[utils.Position]Location, input []string) {
	for r, row := range input {
		for c, column := range row {
			pos := utils.NewPosition(r, c)
			loc := Location{land: string(column), visited: false}
			grid[pos] = loc
		}
	}
}

func part1(grid map[utils.Position]Location, start_pos utils.Position, rows int, cols int, loops int) int {
	directions := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	type Entry struct {
		steps int
		pos   utils.Position
	}

	visited := make(map[utils.Position][]int)

	q := utils.Queue[Entry]{}
	q.Push(Entry{
		steps: 0,
		pos:   start_pos,
	})

	for !q.IsEmpty() {
		entry := q.Pop()

		if steps, exists := visited[entry.pos]; exists {
			if utils.IntInSlice(entry.steps, steps) {
				continue
			} else {
				steps = append(steps, entry.steps)
				visited[entry.pos] = steps
			}
		} else {
			visited[entry.pos] = []int{entry.steps}
		}

		if entry.steps >= loops {
			new_loc := grid[entry.pos]
			new_loc.finished = true
			grid[entry.pos] = new_loc
			continue
		}

		for _, direction := range directions {
			new_pos := entry.pos.Move(direction, 1)

			if 0 <= new_pos.Col && new_pos.Col <= cols && 0 <= new_pos.Row && new_pos.Row <= rows {
				if grid[new_pos].land != "#" {
					q.Push(Entry{
						steps: entry.steps + 1,
						pos:   new_pos,
					})
				}
			}
		}
	}

	visited_count := 0

	for row := 0; row <= rows; row++ {
		for col := 0; col <= cols; col++ {
			pos := utils.NewPosition(row, col)
			loc := grid[pos]

			if loc.finished {
				visited_count++
				fmt.Print("O")
			} else {
				fmt.Print(loc.land)
			}
		}
		fmt.Println()
	}

	return visited_count
}

func main() {
	input := utils.ReadInput("input.txt")
	grid := make(map[utils.Position]Location)
	var start_pos utils.Position

	rows := 0
	cols := 0

	for r, row := range input {
		for c, column := range row {
			pos := utils.NewPosition(r, c)
			loc := Location{land: string(column), visited: false}
			grid[pos] = loc
			if column == 'S' {
				start_pos = pos
			}
			cols = max(cols, r)
		}
		rows = max(rows, r)
	}
	fmt.Println(part1(grid, start_pos, rows, cols, 64))

	// Step 2, find values in a 5.5 grid

	needed_steps := 26501365
	grid_length := 131
	mod := needed_steps % grid_length

	for r, row := range input {
		for c, column := range row {
			for x := 1; x <= 5; x++ {
				for y := 1; y <= 5; y++ {
					pos := utils.NewPosition(r*x, c*y)
					loc := Location{land: string(column), visited: false}
					grid[pos] = loc
				}
			}
			cols = max(cols, r)
		}
		rows = max(rows, r)
	}
	new_rows := rows * 5
	steps55 := grid_length*2 + mod
	start_pos = utils.NewPosition(grid_length*2+mod, grid_length*2+mod)
	total55 := part1(grid, start_pos, new_rows, new_rows, grid_length*2+mod)

	// find values in a 7.7 grid

	for r, row := range input {
		for c, column := range row {
			for x := 1; x <= 7; x++ {
				for y := 1; y <= 7; y++ {
					pos := utils.NewPosition(r*x, c*y)
					loc := Location{land: string(column), visited: false}
					grid[pos] = loc
				}
			}
			cols = max(cols, r)
		}
		rows = max(rows, r)
	}
	new_rows = rows * 7
	steps77 := grid_length*3 + mod
	start_pos = utils.NewPosition(grid_length*3+mod, grid_length*3+mod)
	total77 := part1(grid, start_pos, new_rows, new_rows, grid_length*3+mod)

	// find values in a 9.9 grid

	for r, row := range input {
		for c, column := range row {
			for x := 1; x <= 9; x++ {
				for y := 1; y <= 9; y++ {
					pos := utils.NewPosition(r*x, c*y)
					loc := Location{land: string(column), visited: false}
					grid[pos] = loc
				}
			}
			cols = max(cols, r)
		}
		rows = max(rows, r)
	}
	new_rows = rows * 9
	steps99 := grid_length*4 + mod
	start_pos = utils.NewPosition(grid_length*4+mod, grid_length*4+mod)
	total99 := part1(grid, start_pos, new_rows, new_rows, grid_length*4+mod)

	// f(n) = total55
	// f(n+2) = total77
	// f(n+4) = totall99

	// x is the number of steps, y is the result
	// gives data points
	// steps55, total55...

	// Gives follow for L(needed_steps)
	L0 := ((needed_steps - steps77) * (needed_steps - steps99)) / ((steps55 - steps77) * (steps55 - steps99))
	L1 := ((needed_steps - steps55) * (needed_steps - steps99)) / ((steps77 - steps55) * (steps77 - steps99))
	L2 := ((needed_steps - steps55) * (needed_steps - steps77)) / ((steps99 - steps55) * (steps99 - steps77))
	fmt.Println(L0)
	fmt.Println(total55)
	fmt.Println(L1)
	fmt.Println(total77)
	fmt.Println(L2)
	fmt.Println(total99)
	result := L0*total55 + L1*total77 + L2*total99

	fmt.Println(result)
}
