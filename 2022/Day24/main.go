package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Ground struct {
	blizzards []utils.Direction
}

type Board map[utils.Position]Ground

func (b Board) blizzardSize() (int, int) {
	var max_col int
	var max_row int
	var min_col int
	var min_row int
	for pos := range b {
		if pos.Col > max_col {
			max_col = pos.Col
		}
		if pos.Col < min_col {
			min_col = pos.Col
		}
		if pos.Row > max_row {
			max_row = pos.Row
		}
		if pos.Row < min_row {
			min_row = pos.Row
		}
	}

	return max_col, max_row - 1
}

// Wrap returns the position when stepping outside of the board.
// Pass in the position before wrapping, and the direction, it takes 1 step
func (b Board) Wrap(d utils.Direction, p utils.Position) utils.Position {
	// If the position doesn't exist, return the position
	if _, exists := b[p]; !exists {
		return p
	}

	// If moving is still within the board return that pos
	if _, valid := b[p.Move(d, 1)]; valid {
		return p.Move(d, 1)
	}

	// Move the other way as far as possible
	for {
		new_pos := p.Move(d.Flip(), 1)
		if _, exists := b[new_pos]; !exists {
			return p
		}
		p = new_pos
	}
}

func (b Board) makeEmptyCopy() Board {
	return_board := Board{}

	for pos := range b {
		return_board[pos] = Ground{}
	}
	return return_board
}

// Cache for avoiding calculating the board again
// Maps board state to round number
type boardCache map[int]Board

func part1(input Board) {
	defer utils.Timer("part1")()

	// We loop every LCM of blizzard size so we can precompute all possible boards
	loopMod := utils.LCM(input.blizzardSize())
	boards := computeBoardStates(loopMod, input)

	// Find the goal position
	blizzardCols, blizzardRows := input.blizzardSize()
	goalPos := utils.NewPosition(blizzardRows+1, blizzardCols)

	fmt.Println(findDistance(
		utils.NewPosition(0, 1), goalPos, boards, 0, loopMod,
	))
}

func part2(input Board) {
	defer utils.Timer("part2")()

	// We loop every LCM of blizzard size so we can precompute all possible boards
	loopMod := utils.LCM(input.blizzardSize())
	boards := computeBoardStates(loopMod, input)

	// Find the goal position
	blizzardCols, blizzardRows := input.blizzardSize()
	goalPos := utils.NewPosition(blizzardRows+1, blizzardCols)

	firstLeg := findDistance(
		utils.NewPosition(0, 1), goalPos, boards, 0, loopMod,
	)
	secondLeg := findDistance(
		goalPos, utils.NewPosition(0, 1), boards, firstLeg, loopMod,
	)
	thirdLeg := findDistance(
		utils.NewPosition(0, 1), goalPos, boards, secondLeg, loopMod,
	)

	fmt.Println(thirdLeg)
}

func findDistance(startPos utils.Position, endPos utils.Position, boards boardCache, startTime int, loopMod int) int {
	elves := map[utils.Position]bool{
		startPos: true,
	}
	dirs := []utils.Direction{
		utils.Up, utils.Down, utils.Left, utils.Right, utils.DontMove,
	}

	// Keep looping, stepping forward in time, tracking all possible positions of elves
	for i := startTime; i >= 0; i++ {
		board := boards[i%loopMod]

		// Loop over every elf, and add possible positions to the set of elves
		newElves := map[utils.Position]bool{}
		for pos := range elves {
			for _, dir := range dirs {
				newPos := pos.Move(dir, 1)
				// If reached goal, return
				if newPos == endPos {
					return i
				}
				// If pos not in board skip
				if _, exists := board[newPos]; !exists {
					continue
				}
				// If in board and not goal, add to list
				if len(board[newPos].blizzards) == 0 {
					newElves[newPos] = true
				}
			}
		}
		elves = newElves
	}
	panic("Didn't find a solution")
}

func computeBoardStates(lcm int, board Board) boardCache {
	bc := boardCache{}
	bc[0] = board

	for i := 1; i <= lcm; i++ {
		bc[i] = advanceBoard(bc[i-1])
	}
	return bc
}

func advanceBoard(board Board) Board {
	// Make an empty copy to shift into
	return_board := board.makeEmptyCopy()

	for pos, ground := range board {
		for _, blizzard := range ground.blizzards {
			// Wrap handles logic for wrapping, find next pos of blizzard
			new_pos := board.Wrap(blizzard, pos)

			// Append it to the new board
			return_board[new_pos] = Ground{
				blizzards: append(return_board[new_pos].blizzards, blizzard),
			}
		}
	}
	return return_board
}

func parseInput(file []string) Board {
	return_board := Board{}

	for y, row := range file {
		for x, column := range row {
			switch column {
			case '#':
				// Skipping the walls makes wrapping easier,
				continue
			case '.':
				return_board[utils.Position{Row: y, Col: x}] = Ground{}
			case '>':
				return_board[utils.Position{Row: y, Col: x}] = Ground{
					blizzards: []utils.Direction{utils.Right},
				}
			case '<':
				return_board[utils.Position{Row: y, Col: x}] = Ground{
					blizzards: []utils.Direction{utils.Left},
				}
			case '^':
				return_board[utils.Position{Row: y, Col: x}] = Ground{
					blizzards: []utils.Direction{utils.Up},
				}
			case 'v':
				return_board[utils.Position{Row: y, Col: x}] = Ground{
					blizzards: []utils.Direction{utils.Down},
				}
			}
		}
	}

	return return_board
}

func main() {
	input := parseInput(utils.ReadInput("input.txt"))
	part1(input)
	input = parseInput(utils.ReadInput("input.txt"))
	part2(input)
}
