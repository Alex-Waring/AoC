package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Board map[utils.Position]Elf

type Elf struct {
	proposed_postion utils.Position
	move             bool
}

func part1(board Board) {
	defer utils.Timer("part1")()

	// Create a list to loop through the possible positions
	direction_evalutions := []utils.Direction{
		utils.Up, utils.Down, utils.Left, utils.Right,
	}
	// printBoard(board)

	for round := 0; round < 10; round++ {
		// Suggest possible positions
		suggestedPositions := map[utils.Position]utils.Position{}
		for pos := range board {
			// If no elves nearby, don't move
			if !elfInBorder(pos, board) {
				continue
			}
			for i := 0; i < 4; i++ {
				// Get the dir index by adding round and modding len list - 1
				dir := direction_evalutions[(round+i)%4]
				// If no elves in that pos then suggest that direction and move on
				// Except first check for dupes
				if !elfInDir(dir, pos, board) {
					new_pos := pos.Move(dir, 1)
					if other_elf_pos, ok := suggestedPositions[new_pos]; ok {
						// We have a colision, set the other Elf to nil
						// don't remove from suggested positions in case three want to move there
						board[other_elf_pos] = Elf{}
					} else {
						board[pos] = Elf{
							proposed_postion: new_pos,
							move:             true,
						}
						suggestedPositions[new_pos] = pos
					}
					break
				}
			}
		}

		// Move all positions
		for pos, elf := range board {
			if elf.move {
				delete(board, pos)
				board[elf.proposed_postion] = Elf{}
			}
		}
		// printBoard(board)
	}

	// Calculate space by find the min and max of x and y and grabbing the rectange - size
	var max_col int
	var max_row int
	var min_col int
	var min_row int
	for pos := range board {
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
	rectangle_size := (max_col - min_col + 1) * (max_row - min_row + 1)
	fmt.Println(rectangle_size - len(board))
}

func part2(board Board) {
	defer utils.Timer("part2")()

	// Create a list to loop through the possible positions
	direction_evalutions := []utils.Direction{
		utils.Up, utils.Down, utils.Left, utils.Right,
	}
	// printBoard(board)

	// Loop through many many, hope we break
	for round := 0; round < 1000000; round++ {
		// Suggest possible positions
		suggestedPositions := map[utils.Position]utils.Position{}
		for pos := range board {
			// If no elves nearby, don't move
			if !elfInBorder(pos, board) {
				continue
			}
			for i := 0; i < 4; i++ {
				// Get the dir index by adding round and modding len list - 1
				dir := direction_evalutions[(round+i)%4]
				// If no elves in that pos then suggest that direction and move on
				// Except first check for dupes
				if !elfInDir(dir, pos, board) {
					new_pos := pos.Move(dir, 1)
					if other_elf_pos, ok := suggestedPositions[new_pos]; ok {
						// We have a colision, set the other Elf to nil
						// don't remove from suggested positions in case three want to move there
						board[other_elf_pos] = Elf{}
					} else {
						board[pos] = Elf{
							proposed_postion: new_pos,
							move:             true,
						}
						suggestedPositions[new_pos] = pos
					}
					break
				}
			}
		}

		total_moved := 0

		// Move all positions
		for pos, elf := range board {
			if elf.move {
				delete(board, pos)
				board[elf.proposed_postion] = Elf{}
				total_moved++
			}
		}

		// printBoard(board)

		if total_moved == 0 {
			fmt.Println(round + 1)
			return
		}

	}

}

func parseInput(file []string) Board {
	return_board := Board{}

	for y, row := range file {
		for x, column := range row {
			if column == '#' {
				return_board[utils.Position{Row: y, Col: x}] = Elf{}
			}
		}
	}

	return return_board
}

// Take a position and a dir, check if any elves are in that direction in a cone
// and reuturn true if so
func elfInDir(dir utils.Direction, pos utils.Position, board Board) bool {
	switch dir {
	case utils.Up:
		if _, ok := board[pos.Move(utils.UpLeft, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.Up, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.UpRight, 1)]; ok {
			return true
		}
	case utils.Down:
		if _, ok := board[pos.Move(utils.DownLeft, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.Down, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.DownRight, 1)]; ok {
			return true
		}
	case utils.Right:
		if _, ok := board[pos.Move(utils.UpRight, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.Right, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.DownRight, 1)]; ok {
			return true
		}
	case utils.Left:
		if _, ok := board[pos.Move(utils.UpLeft, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.Left, 1)]; ok {
			return true
		} else if _, ok := board[pos.Move(utils.DownLeft, 1)]; ok {
			return true
		}
	}
	return false
}

// Check every direction and if there is an elf there return true
func elfInBorder(pos utils.Position, board Board) bool {
	dirs := []utils.Direction{utils.UpLeft, utils.Up, utils.UpRight, utils.Left, utils.Right, utils.DownLeft, utils.Down, utils.DownRight}

	for _, dir := range dirs {
		if _, ok := board[pos.Move(dir, 1)]; ok {
			return true
		}
	}
	return false
}

func printBoard(board Board) {
	var max_col int
	var max_row int
	var min_col int
	var min_row int
	for pos := range board {
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

	for row := min_row; row <= max_row; row++ {
		for col := min_col; col <= max_col; col++ {
			if _, ok := board[utils.Position{Row: row, Col: col}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
}

func main() {
	input := parseInput(utils.ReadInput("input.txt"))

	part1(input)

	input2 := parseInput(utils.ReadInput("input.txt"))
	part2(input2)
}
