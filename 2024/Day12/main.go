package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	board := make(map[utils.Position]rune)
	for row, line := range input {
		for col, c := range line {
			board[utils.Position{Row: row, Col: col}] = c
		}
	}

	dirs := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	seen := map[utils.Position]bool{}
	total := 0

	for pos := range board {
		if seen[pos] {
			continue
		}
		seen[pos] = true

		area := 1
		perimeter := 0

		q := utils.Queue[utils.Position]{}
		q.Push(pos)

		for !q.IsEmpty() {
			loc := q.Pop()
			for _, dir := range dirs {
				new_loc := loc.Move(dir, 1)
				if board[new_loc] != board[loc] {
					perimeter++
				} else if !seen[new_loc] {
					seen[new_loc] = true
					q.Push(new_loc)
					area++
				}
			}
		}
		total += area * perimeter
	}
	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	board := make(map[utils.Position]rune)
	for row, line := range input {
		for col, c := range line {
			board[utils.Position{Row: row, Col: col}] = c
		}
	}

	dirs := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	seen := map[utils.Position]bool{}
	total := 0

	for pos := range board {
		if seen[pos] {
			continue
		}
		seen[pos] = true

		area := 1
		sides := 0

		q := utils.Queue[utils.Position]{}
		q.Push(pos)

		for !q.IsEmpty() {
			loc := q.Pop()
			for _, dir := range dirs {
				new_loc := loc.Move(dir, 1)
				if board[new_loc] != board[loc] {
					// Find two possible corners:
					// 1. 90 degrees anti-clockwise from dir and dir are not the same as the current loc
					// This is a corner around P

					// .....
					// ...D.
					// ..1P.
					// .....
					check1 := loc.Move(rotate(dir), 1)

					// 2. 45 degrees anti-clockwise from dir is the same but dir is not
					// this is a corner around D
					// .....
					// ..2D.
					// ...P.
					// .....

					check2 := check1.Move(dir, 1)

					// If check1 is not the same as the current loc, or check2 is the same as the current loc
					// then we are at a corner, and can add a side
					if board[check1] != board[loc] || board[check2] == board[loc] {
						sides++
					}
				} else if !seen[new_loc] {
					seen[new_loc] = true
					q.Push(new_loc)
					area++
				}
			}
		}
		total += area * sides
	}
	fmt.Println(total)
}

func rotate(dir utils.Direction) utils.Direction {
	switch dir {
	case utils.Up:
		return utils.Left
	case utils.Down:
		return utils.Right
	case utils.Left:
		return utils.Down
	case utils.Right:
		return utils.Up
	}
	panic("Invalid direction")
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
