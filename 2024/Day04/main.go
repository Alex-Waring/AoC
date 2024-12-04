package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	board := utils.NewBoard()

	for r, row := range input {
		for c, char := range row {
			board[utils.Position{Row: r, Col: c}] = string(char)
		}
	}

	found := 0

	// Check every position to see if it starts an XMAS
	for pos, letter := range board {
		if letter != "X" {
			continue
		}

		dirs := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right, utils.UpRight, utils.UpLeft, utils.DownLeft, utils.DownRight}

		for _, dir := range dirs {
			word := letter + board[pos.Move(dir, 1)] + board[pos.Move(dir, 2)] + board[pos.Move(dir, 3)]
			if word == "XMAS" {
				found++
			}
		}
	}
	fmt.Println(found)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	board := utils.NewBoard()

	for r, row := range input {
		for c, char := range row {
			board[utils.Position{Row: r, Col: c}] = string(char)
		}
	}

	found := 0

	// This time check id we're an A, and then check for pairs in the corners
	for pos, letter := range board {
		if letter != "A" {
			continue
		}
		up_slash := board[pos.Move(utils.DownLeft, 1)] + letter + board[pos.Move(utils.UpRight, 1)]
		if !(up_slash == "MAS" || up_slash == "SAM") {
			continue
		}
		down_slash := board[pos.Move(utils.DownRight, 1)] + letter + board[pos.Move(utils.UpLeft, 1)]
		if !(down_slash == "MAS" || down_slash == "SAM") {
			continue
		}
		found++
	}
	fmt.Println(found)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
