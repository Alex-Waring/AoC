package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) map[utils.Position]bool {
	defer utils.Timer("part1")()

	board := utils.NewBoard()
	location := utils.Location{}

	for r, row := range input {
		for c, char := range row {
			if char == '^' {
				location = utils.Location{
					Pos: utils.NewPosition(r, c),
					Dir: utils.Up,
				}
				board[utils.Position{Row: r, Col: c}] = "."
			} else {
				board[utils.Position{Row: r, Col: c}] = string(char)
			}
		}
	}

	places_been := map[utils.Position]bool{}
	places_been[location.Pos] = true

	for true {
		// Move the location in the direction it's facing
		new_pos := location.Pos.Move(location.Dir, 1)
		// If the new position is off the board, stop
		if _, ok := board[new_pos]; !ok {
			break
		}
		// If the new position is a wall, turn right
		if board[new_pos] == "#" {
			location.Dir = location.Dir.Turn(utils.Right)
			new_pos = location.Pos.Move(location.Dir, 1)
		}
		// If the new position is open space, move there
		location.Pos = new_pos
		places_been[location.Pos] = true
	}

	fmt.Println(len(places_been))
	return places_been
}

func part2(input []string) {
	defer utils.Timer("part2")()

	board := utils.NewBoard()
	start_location := utils.Location{}

	for r, row := range input {
		for c, char := range row {
			if char == '^' {
				start_location = utils.Location{
					Pos: utils.NewPosition(r, c),
					Dir: utils.Up,
				}
				board[utils.Position{Row: r, Col: c}] = "."
			} else {
				board[utils.Position{Row: r, Col: c}] = string(char)
			}
		}
	}

	// Do part 1 to find the places we've been
	places_to_check := part1(input)

	// Add a new obstruction to the board
	found_loops := 0
	tried := 0
	for obstruction, _ := range places_to_check {
		tried++
		if tried%1000 == 0 {
			fmt.Println(tried)
		}
		if board[obstruction] == "#" || board[obstruction] == "^" {
			continue
		}
		new_board := utils.NewBoard()
		for k, v := range board {
			if k == obstruction {
				new_board[k] = "#"
			} else {
				new_board[k] = v
			}
		}

		location := utils.Location{
			Pos: utils.Position{
				Row: start_location.Pos.Row,
				Col: start_location.Pos.Col,
			},
			Dir: start_location.Dir,
		}

		places_been := map[utils.Position]bool{}
		places_been[location.Pos] = true
		loop_tracker := map[utils.Location]bool{}

		for true {
			// Move the location in the direction it's facing
			new_pos := location.Pos.Move(location.Dir, 1)
			// If the new position is off the board, stop
			if _, ok := new_board[new_pos]; !ok {
				break
			}
			// If the new position is a wall, turn right, loop until open
			for true {
				if new_board[new_pos] == "#" {
					location.Dir = location.Dir.Turn(utils.Right)
					new_pos = location.Pos.Move(location.Dir, 1)
				}
				// If the new position is open space, move there
				if new_board[new_pos] == "." {
					location.Pos = new_pos
					break
				}
			}

			// If we've been here before, we've looped
			if _, ok := loop_tracker[location]; ok {
				found_loops++
				break
			} else {
				loop_tracker[location] = true
			}
			places_been[location.Pos] = true
		}
	}
	fmt.Println(found_loops)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
