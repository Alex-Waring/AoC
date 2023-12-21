package main

import (
	"fmt"
	"math"

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

func find_visited(grid map[utils.Position]Location, visited map[utils.Position]int, start_pos utils.Position, rows int, cols int) map[utils.Position]int {
	directions := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	type Entry struct {
		steps int
		pos   utils.Position
	}

	q := utils.Queue[Entry]{}
	q.Push(Entry{
		steps: 0,
		pos:   start_pos,
	})

	for !q.IsEmpty() {
		entry := q.Pop()

		if _, exists := visited[entry.pos]; !exists {
			visited[entry.pos] = entry.steps
		} else if entry.steps < visited[entry.pos] {
			visited[entry.pos] = entry.steps
		} else {
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
	return visited
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

	// Step 2

	needed_steps := 26501365
	grid_length := 131

	n := math.Floor(float64(needed_steps / grid_length))
	fmt.Println(n)

	even := n * n
	odd := (n + 1) * (n + 1)
	even_corners := n
	odd_corners := (n - 1)

	visited := map[utils.Position]int{}
	visited = find_visited(grid, visited, start_pos, rows, cols)

	var even_corners_val int
	var odd_corners_val int
	var even_full_val int
	var odd_full_val int
	for _, steps := range visited {
		if steps%2 == 0 {
			even_full_val++
			if steps > 65 {
				even_corners_val++
			}
		} else {
			odd_full_val++
			if steps > 65 {
				odd_corners_val++
			}
		}
	}
	result := even_corners*float64(even_corners_val) - odd_corners*float64(odd_corners_val) + even*float64(even_full_val) + odd*float64(odd_full_val)
	fmt.Println(int(result))
}
