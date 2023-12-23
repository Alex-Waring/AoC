package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Entry struct {
	pos     utils.Position
	history []utils.Position
}

func getPossibleMoves(grid map[utils.Position]string, pos utils.Position, max_row int, max_col int) []utils.Position {
	directions := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}
	return_list := []utils.Position{}

	switch grid[pos] {
	case ">":
		new_pos := pos.Move(utils.Right, 1)
		if char, exists := grid[new_pos]; exists && char != "#" {
			return_list = append(return_list, new_pos)
		}
	case "v":
		new_pos := pos.Move(utils.Down, 1)
		if char, exists := grid[new_pos]; exists && char != "#" {
			return_list = append(return_list, new_pos)
		}
	case ".":
		for _, direction := range directions {
			new_pos := pos.Move(direction, 1)
			if char, exists := grid[new_pos]; exists && char != "#" {
				return_list = append(return_list, new_pos)
			}
		}
	}
	return return_list
}

func getPossibleMovesPart2(grid map[utils.Position]string, pos utils.Position, max_row int, max_col int) []utils.Position {
	directions := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}
	return_list := []utils.Position{}

	for _, direction := range directions {
		new_pos := pos.Move(direction, 1)
		if char, exists := grid[new_pos]; exists && char != "#" {
			return_list = append(return_list, new_pos)
		}
	}
	return return_list
}

func drawMap(grid map[utils.Position]string, history []utils.Position, max_row int, max_col int) {
	for row := 0; row <= max_row; row++ {
		for col := 0; col <= max_col; col++ {
			pos := utils.NewPosition(row, col)
			loc := grid[pos]

			if utils.PosInSlice(pos, history) {
				fmt.Print("O")
			} else {
				fmt.Print(loc)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(grid map[utils.Position]string, finish utils.Position, max_row int, max_col int) {
	defer utils.Timer("part1")()
	finished := []int{}

	q := utils.Queue[Entry]{}
	q.Push(Entry{
		pos:     utils.Position{Row: 0, Col: 1},
		history: []utils.Position{},
	})

	for !q.IsEmpty() {
		entry := q.Pop()

		// If we're at the finish, add the number of steps to finished
		if entry.pos == finish {
			finished = append(finished, len(entry.history))
			continue
		}

		// Get the possible moves
		next_moves := getPossibleMoves(grid, entry.pos, max_row, max_col)

		// For each possible move, if we've been there in the history discard
		// If we havn't, create a new entry with it in
		for _, move := range next_moves {
			if !utils.PosInSlice(move, entry.history) {
				q.Push(Entry{
					pos:     move,
					history: append(entry.history, entry.pos),
				})
				// drawMap(grid, append(entry.history, entry.pos), max_row, max_col)
			}
		}
	}
	fmt.Println(utils.FindMax(finished))
}

func findNextJunction(grid map[utils.Position]string, entry Entry, max_row int, max_col int, finish utils.Position) []Entry {
	return_list := []Entry{}

	if entry.pos == finish {
		return_list = append(return_list, entry)
		return return_list
	}

	// Get the possible moves
	next_moves := getPossibleMovesPart2(grid, entry.pos, max_row, max_col)

	// If we have two possible moves, add the one that isn't in the history to the
	// entry and continue
	if len(next_moves) == 2 {
		for _, move := range next_moves {
			if !utils.PosInSlice(move, entry.history) {
				new_entry := Entry{
					history: append(entry.history, entry.pos),
					pos:     move,
				}
				return findNextJunction(grid, new_entry, max_row, max_col, finish)
			}
		}
	} else {
		// Otherwise do the normal loop and try and solve it
		for _, move := range next_moves {
			if !utils.PosInSlice(move, entry.history) {
				return_list = append(return_list, Entry{
					pos:     move,
					history: append(entry.history, entry.pos),
				})
			}
		}
	}

	return return_list
}

func part2(grid map[utils.Position]string, finish utils.Position, max_row int, max_col int) {
	defer utils.Timer("part2")()
	finished := []int{}

	q := utils.Queue[Entry]{}
	q.Push(Entry{
		pos:     utils.Position{Row: 0, Col: 1},
		history: []utils.Position{},
	})

	for !q.IsEmpty() {
		entry := q.Pop()

		// If we're at the finish, add the number of steps to finished
		if entry.pos == finish {
			finished = append(finished, len(entry.history))
			continue
		}

		next_moves := findNextJunction(grid, entry, max_row, max_col, finish)
		for _, next := range next_moves {
			q.Push(next)
		}
	}
	fmt.Println(utils.FindMax(finished))
}

func main() {
	lines := utils.ReadInput("input.txt")
	grid := map[utils.Position]string{}

	max_row := len(lines) - 1
	max_col := len(lines[0]) - 1

	for row, line := range lines {
		for col, char := range line {
			new_pos := utils.NewPosition(row, col)
			grid[new_pos] = string(char)
		}
	}

	finish := utils.NewPosition(max_row, max_col-1)

	part1(grid, finish, max_row, max_col)
	part2(grid, finish, max_row, max_col)
}
