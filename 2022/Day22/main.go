package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(commands string, board utils.Board, pos utils.Position) {
	defer utils.Timer("part1")()
	instructions := parseCommand(commands)
	var final_dir utils.Direction

	for _, i := range instructions {
		for d, s := range i {
			pos = moveForward(board, pos, d, s)
			final_dir = d
		}
	}
	fmt.Println(1000*(pos.Row+1) + 4*(pos.Col+1) + facingToNum(final_dir))
}

func part2(commands string, board utils.Board, pos utils.Position) {
	defer utils.Timer("part2")()
	instructions := parseCommand(commands)
	var final_dir utils.Direction

	for _, i := range instructions {
		for d, s := range i {
			pos = moveCubeForward(board, pos, d, s)
			final_dir = d
		}
	}
	fmt.Println(1000*(pos.Row+1) + 4*(pos.Col+1) + facingToNum(final_dir))
}

func parseCommand(command string) []map[utils.Direction]int {
	commands := []map[utils.Direction]int{}
	r := regexp.MustCompile("[0-9]*[RL]")
	facing := utils.Right

	for _, c := range r.FindAllString(command, -1) {
		rotation := c[len(c)-1:]
		steps := utils.IntegerOf(c[:len(c)-1])

		commands = append(commands, map[utils.Direction]int{facing: steps})

		if rotation == "R" {
			facing = facing.Turn(utils.Right)
		} else {
			facing = facing.Turn(utils.Left)
		}
	}

	// Find the last command
	r = regexp.MustCompile("[0-9]*$")
	steps := utils.IntegerOf(r.FindString(command))
	commands = append(commands, map[utils.Direction]int{facing: steps})
	return commands
}

func moveForward(board utils.Board, start_pos utils.Position, direction utils.Direction, steps int) utils.Position {
	for step := 0; step < steps; step++ {
		new_pos := start_pos.Move(direction, 1)

		if _, exists := board[new_pos]; !exists {
			new_pos = board.Wrap(direction, start_pos)
		}

		if board[new_pos] == "#" {
			return start_pos
		}
		start_pos = new_pos
	}
	return start_pos
}

func moveCubeForward(board utils.Board, start_pos utils.Position, direction utils.Direction, steps int) utils.Position {
	for step := 0; step < steps; step++ {
		new_pos := start_pos.Move(direction, 1)

		if _, exists := board[new_pos]; !exists {
			new_pos = cubeWrap(direction, start_pos, board)
		}

		if board[new_pos] == "#" {
			return start_pos
		}
		start_pos = new_pos
	}
	return start_pos
}

func facingToNum(facing utils.Direction) int {
	switch facing {
	case utils.Right:
		return 0
	case utils.Down:
		return 1
	case utils.Left:
		return 2
	case utils.Up:
		return 3
	}
	return 0
}

// This function wraps based on the map being folded into a cube.
func cubeWrap(d utils.Direction, p utils.Position, b utils.Board) utils.Position {
	face := findFace(b, p)
	face_length := int(math.Sqrt(float64(len(b)) / 6))

	// Damn I hate this, but it works I think
	switch face {
	case 1:
		if d == utils.Up {
			new_pos := utils.Position{
				Row: (p.Col - face_length) + 3*face_length,
				Col: 0,
			}
			return new_pos
		} else {
			new_pos := utils.Position{
				Row: p.Row - face_length*2,
				Col: 0,
			}
			return new_pos
		}
	case 2:
		if d == utils.Up {
			new_pos := utils.Position{
				Row: b.Bottom(),
				Col: (p.Col - 2*face_length),
			}
			return new_pos
		} else if d == utils.Down {
			new_pos := utils.Position{
				Row: p.Col,
				Col: p.Row,
			}
			return new_pos
		} else {
			new_pos := utils.Position{
				Row: p.Row + 2*face_length,
				Col: p.Col - face_length,
			}
			return new_pos
		}
	case 3:
		if d == utils.Left {
			new_pos := utils.Position{
				Row: 2 * face_length,
				Col: p.Row - face_length,
			}
			return new_pos
		} else {
			new_pos := utils.Position{
				Row: p.Col,
				Col: p.Row,
			}
			return new_pos
		}
	case 4:
		if d == utils.Left {
			new_pos := utils.Position{
				Row: p.Row - 2*face_length,
				Col: face_length,
			}
			return new_pos
		} else {
			new_pos := utils.Position{
				Row: p.Col + face_length,
				Col: face_length,
			}
			return new_pos
		}
	case 5:
		if d == utils.Right {
			new_pos := utils.Position{
				Row: face_length - (p.Row - 2*face_length),
				Col: 3 * face_length,
			}
			return new_pos
		} else {
			new_pos := utils.Position{
				Row: p.Col + 2*face_length,
				Col: face_length - 1,
			}
			return new_pos
		}
	case 6:
		if d == utils.Right {
			new_pos := utils.Position{
				Row: 3 * face_length,
				Col: p.Row - 2*face_length,
			}
			return new_pos
		} else if d == utils.Down {
			new_pos := utils.Position{
				Row: 0,
				Col: 2*face_length + p.Col,
			}
			return new_pos
		} else {
			new_pos := utils.Position{
				Row: 0,
				Col: p.Row - 2*face_length,
			}
			return new_pos
		}
	}
	panic("not handled")
}

// Return an index for the face of the cube, following
//
//	1122
//	1122
//	33
//	33
//
// 4455
// 4455
// 66
// 66
func findFace(b utils.Board, p utils.Position) int {
	face_length := int(math.Sqrt(float64(len(b)) / 6))

	// Real Input
	if p.Row < face_length {
		if p.Col < 3*face_length {
			return 4
		} else {
			return 5
		}
	} else if p.Row >= 2*face_length {
		return 2
	} else if p.Col > face_length {
		return 1
	} else if p.Col > 2*face_length {
		return 3
	} else {
		return 5
	}
}

func main() {
	input := utils.ReadInput("input.txt")

	commands := input[len(input)-1]
	b := input[:len(input)-2]
	board := utils.NewBoard()

	start_x := 100

	for y, row := range b {
		for x, c := range row {
			if c != ' ' {
				if x < start_x && y == 0 {
					start_x = x
				}
				board[utils.Position{Row: y, Col: x}] = string(c)
			}
		}
	}
	start_pos := utils.Position{Row: 0, Col: start_x}
	findFace(board, start_pos)
	part1(commands, board, start_pos)
	part2(commands, board, start_pos)
}
