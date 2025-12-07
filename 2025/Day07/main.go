package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	board := utils.NewBoard()
	startPos := utils.Position{}

	for row, line := range input {
		for col, char := range line {
			board[utils.Position{Row: row, Col: col}] = string(char)
			if string(char) == "S" {
				startPos = utils.Position{Row: row, Col: col}
			}
		}
	}

	bottom := board.Bottom()

	q := utils.Queue[utils.Position]{}
	visited := map[utils.Position]bool{}
	q.Push(startPos)
	split := 0
	for !q.IsEmpty() {
		current := q.Pop()
		if visited[current] {
			continue
		} else {
			visited[current] = true
		}
		if current.Row == bottom {
			continue
		}
		newPos := current.Move(utils.Down, 1)
		if _, exists := board[newPos]; exists {
			if board[newPos] == "." {
				q.Push(newPos)
			} else if board[newPos] == "^" {
				q.Push(newPos.Move(utils.Right, 1))
				q.Push(newPos.Move(utils.Left, 1))
				split++
			}
		}
	}
	fmt.Println(split)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	board := utils.NewBoard()

	for row, line := range input {
		for col, char := range line {
			board[utils.Position{Row: row, Col: col}] = string(char)
			if string(char) == "S" {
				board[utils.Position{Row: row, Col: col}] = "1"
			}
		}
	}

	for row := 0; row < board.Bottom(); row++ {
		for col := 0; col <= board.Right(); col++ {
			val := board[utils.Position{Row: row, Col: col}]
			if val == "^" || val == "." {
				continue
			}
			timelines := utils.IntegerOf(val)
			nextPos := utils.Position{Row: row + 1, Col: col}
			nextVal := board[nextPos]
			if nextVal == "." {
				board[nextPos] = fmt.Sprintf("%d", timelines)
			} else if nextVal == "^" {
				nextPosLeft := nextPos.Move(utils.Left, 1)
				nextPosRight := nextPos.Move(utils.Right, 1)
				valLeft := board[nextPosLeft]
				valRight := board[nextPosRight]
				// We can assume that no two ^ are next to each other
				if valLeft == "." {
					board[nextPosLeft] = fmt.Sprintf("%d", timelines)
				} else if valLeft != "^" {
					existing := utils.IntegerOf(valLeft)
					board[nextPosLeft] = fmt.Sprintf("%d", existing+timelines)
				}
				if valRight == "." {
					board[nextPosRight] = fmt.Sprintf("%d", timelines)
				} else if valRight != "^" {
					existing := utils.IntegerOf(valRight)
					board[nextPosRight] = fmt.Sprintf("%d", existing+timelines)
				}
			} else {
				existing := utils.IntegerOf(nextVal)
				board[nextPos] = fmt.Sprintf("%d", existing+timelines)
			}
		}
	}

	total := 0
	bottom := board.Bottom()
	for col := 0; col <= board.Right(); col++ {
		pos := utils.Position{Row: bottom, Col: col}
		val := board[pos]
		if val == "." || val == "^" {
			continue
		}
		total += utils.IntegerOf(val)
	}
	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
