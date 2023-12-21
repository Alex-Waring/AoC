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
			}
		}
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

	// The next block centre is 131 steps away, if we can reach it then we can do the max number
	// of steps in our current block, and we can skip to the next one

	// Fan out from there, if we have odd left, different max to even

	// When we hit edge blocks, we can calculate manually

	needed_steps := 26501365
	next_block := 131

	total := 0

	set_grid(grid, input)
	max_even := part1(grid, start_pos, rows, cols, next_block+1)
	set_grid(grid, input)
	max_odd := part1(grid, start_pos, rows, cols, next_block)

	full_blocks := (needed_steps - (needed_steps % next_block)) / next_block
	fmt.Println(full_blocks)

	type Entry struct {
		block utils.Position
		steps int
	}

	q := utils.Queue[Entry]{}
	seen := map[Entry]int{}
	directions := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	q.Push(Entry{
		block: utils.NewPosition(0, 0),
		steps: 26501365,
	})

	cache := map[int]int{}

	for !q.IsEmpty() {
		entry := q.Pop()
		fmt.Println(q.Len())

		if entry.steps > next_block {
			if entry.block.ManhattanZero()%2 == 0 {
				total += max_even
			} else {
				total += max_odd
			}

			for _, direction := range directions {
				new_pos := entry.block.Move(direction, 1)
				if new_pos.ManhattanZero()%2 == 0 {
					total += max_even
					seen[entry] = max_even
				} else {
					total += max_odd
					seen[entry] = max_odd
				}
				q.Push(Entry{
					block: new_pos,
					steps: entry.steps - next_block,
				})
			}
		} else {
			var result int
			if cache_entry, exists := cache[entry.steps]; exists {
				result = cache_entry
			} else {
				set_grid(grid, input)
				result := part1(grid, start_pos, rows, cols, entry.steps)
				seen[entry] = result
				cache[entry.steps] = result
			}
			total += result

		}
	}
	fmt.Println(total)

}
