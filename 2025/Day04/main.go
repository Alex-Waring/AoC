package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func findStacksToRemove(board utils.Board) []utils.Position {
	stacksToRemove := []utils.Position{}

	for pos, val := range board {
		if val != "@" {
			continue
		}
		posToCheck := []utils.Position{
			pos.Slide(-1, -1), pos.Slide(-1, 0), pos.Slide(-1, 1),
			pos.Slide(0, -1), pos.Slide(0, 1),
			pos.Slide(1, -1), pos.Slide(1, 0), pos.Slide(1, 1),
		}
		count := 0
		for _, checkPos := range posToCheck {
			if board[checkPos] == "@" {
				count++
			}
		}
		if count < 4 {
			stacksToRemove = append(stacksToRemove, pos)
		}
	}
	return stacksToRemove
}

func part1(input []string) {
	defer utils.Timer("part1")()
	board := utils.NewBoard()

	for row, line := range input {
		for col, char := range line {
			board[utils.Position{Row: row, Col: col}] = string(char)
		}
	}

	validPaper := findStacksToRemove(board)

	fmt.Println(len(validPaper))
}

func part2(input []string) {
	defer utils.Timer("part2")()

	board := utils.NewBoard()

	for row, line := range input {
		for col, char := range line {
			board[utils.Position{Row: row, Col: col}] = string(char)
		}
	}

	stacksCanBeRemoved := true
	rollsRemoved := 0

	for stacksCanBeRemoved {
		stacksToRemove := findStacksToRemove(board)
		if len(stacksToRemove) == 0 {
			stacksCanBeRemoved = false
		} else {
			rollsRemoved += len(stacksToRemove)
			for _, pos := range stacksToRemove {
				board[pos] = "."
			}
		}
	}
	fmt.Println(rollsRemoved)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
