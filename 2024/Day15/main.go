package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	grid := utils.NewBoard()
	var start utils.Position

	for row, line := range input {
		if line == "" {
			break
		}
		for col, char := range line {
			grid[utils.NewPosition(row, col)] = string(char)
			if char == '@' {
				start = utils.NewPosition(row, col)
			}
		}
	}

	instructions := []string{}

	for _, line := range input {
		if strings.Contains(line, "#") {
			continue
		}
		for _, char := range line {
			instructions = append(instructions, string(char))
		}
	}

	for _, instruction := range instructions {
		dir := directionOf(instruction)
		squaresMoving := checkDirection(grid, []utils.Position{start}, dir, start)

		if len(squaresMoving) == 1 {
			continue
		}

		// We now have a list of locations to move, the first of which is where the robot was
		// and the second is where the robot is moving to. We need to update the grid to reflect
		grid[squaresMoving[0]] = "."
		grid[squaresMoving[1]] = "@"
		start = squaresMoving[1]
		if len(squaresMoving) > 2 {
			for i := 2; i < len(squaresMoving); i++ {
				grid[squaresMoving[i]] = "O"
			}
		}
	}

	total := 0

	for loc, val := range grid {
		if val == "O" {
			total += 100*loc.Row + loc.Col
		}
	}
	fmt.Println(total)
}

func checkDirection(grid utils.Board, pos []utils.Position, dir utils.Direction, start utils.Position) []utils.Position {
	lastPos := pos[len(pos)-1]
	newLoc := lastPos.Move(dir, 1)
	if grid[newLoc] == "." {
		return append(pos, newLoc)
	}
	if grid[newLoc] == "#" {
		return []utils.Position{start}
	}
	if grid[newLoc] == "O" {
		return checkDirection(grid, append(pos, newLoc), dir, start)
	}
	fmt.Println(grid[newLoc])
	panic("Invalid grid value")
}

func directionOf(arrow string) utils.Direction {
	switch arrow {
	case "^":
		return utils.Up
	case "v":
		return utils.Down
	case "<":
		return utils.Left
	case ">":
		return utils.Right
	default:
		panic("Invalid arrow")
	}
}

func part2(input []string) {
	defer utils.Timer("part2")()
	grid := utils.NewBoard()
	var start utils.Position

	for row, line := range input {
		if line == "" {
			break
		}
		for col, char := range line {
			if char == '#' {
				grid[utils.NewPosition(row, col*2)] = "#"
				grid[utils.NewPosition(row, col*2+1)] = "#"
			} else if char == '@' {
				grid[utils.NewPosition(row, col*2)] = "@"
				grid[utils.NewPosition(row, col*2+1)] = "."
				start = utils.NewPosition(row, col*2)
			} else if char == '.' {
				grid[utils.NewPosition(row, col*2)] = "."
				grid[utils.NewPosition(row, col*2+1)] = "."
			} else {
				grid[utils.NewPosition(row, col*2)] = "["
				grid[utils.NewPosition(row, col*2+1)] = "]"
			}
		}
	}

	instructions := []string{}

	for _, line := range input {
		if strings.Contains(line, "#") {
			continue
		}
		for _, char := range line {
			instructions = append(instructions, string(char))
		}
	}

	for _, instruction := range instructions {
		dir := directionOf(instruction)
		newLoc := start.Move(dir, 1)
		if grid[newLoc] == "#" {
			continue
		}
		if grid[newLoc] == "." {
			grid[start] = "."
			grid[newLoc] = "@"
			start = newLoc
			continue
		}
		var boxPos [2]utils.Position
		if grid[newLoc] == "[" {
			boxPos = [2]utils.Position{newLoc, newLoc.Move(utils.Right, 1)}
		} else {
			boxPos = [2]utils.Position{newLoc, newLoc.Move(utils.Left, 1)}
		}
		boxesToMove := findMovedBoxes(grid, boxPos, dir)
		canMove := true
		for _, box := range boxesToMove {
			if !boxCanMove(grid, box, dir) {
				canMove = false
				break
			}
		}
		if !canMove {
			continue
		}
		grid = moveBoxes(grid, boxesToMove, dir)
		grid[start] = "."
		grid[newLoc] = "@"
		start = newLoc
	}

	total := 0

	for loc, val := range grid {
		if val == "[" {
			total += 100*loc.Row + loc.Col
		}
	}
	fmt.Println(total)
}

func moveBoxes(grid utils.Board, boxes [][2]utils.Position, dir utils.Direction) utils.Board {
	currentLeft := map[utils.Position]bool{}
	currentRight := map[utils.Position]bool{}

	for _, box := range boxes {
		box = orderBoxPos(box)
		currentLeft[box[0]] = true
		currentRight[box[1]] = true
	}

	finalLeft := map[utils.Position]bool{}
	finalRight := map[utils.Position]bool{}

	for loc := range currentLeft {
		newLoc := loc.Move(dir, 1)
		finalLeft[newLoc] = true
	}
	for loc := range currentRight {
		newLoc := loc.Move(dir, 1)
		finalRight[newLoc] = true
	}

	// Locations in the first set and not the second must become .
	for loc := range currentLeft {
		if _, ok := finalLeft[loc]; !ok {
			grid[loc] = "."
		}
	}
	for loc := range currentRight {
		if _, ok := finalRight[loc]; !ok {
			grid[loc] = "."
		}
	}

	// Locations in the second set and not the first must become a box
	for loc := range finalLeft {
		if _, ok := currentLeft[loc]; !ok {
			grid[loc] = "["
		}
	}
	for loc := range finalRight {
		if _, ok := currentRight[loc]; !ok {
			grid[loc] = "]"
		}
	}

	return grid
}

func boxCanMove(grid utils.Board, pos [2]utils.Position, dir utils.Direction) bool {
	if grid[pos[1].Move(dir, 1)] == "#" || grid[pos[0].Move(dir, 1)] == "#" {
		return false
	}
	return true
}

func findMovedBoxes(grid utils.Board, pos [2]utils.Position, dir utils.Direction) [][2]utils.Position {
	newLoc1 := pos[0].Move(dir, 1)
	newLoc2 := pos[1].Move(dir, 1)

	list := [][2]utils.Position{}
	pos = orderBoxPos(pos)
	list = append(list, pos)

	// index 0 is on the left, index 1 is on the right
	newLoc := orderBoxPos([2]utils.Position{newLoc1, newLoc2})

	switch dir {
	case utils.Right:
		// Only check if the right side of the box can move
		if grid[newLoc[1]] == "[" {
			moved := [2]utils.Position{newLoc[1], newLoc[1].Move(dir, 1)}
			list = append(list, findMovedBoxes(grid, moved, dir)...)
		}
	case utils.Left:
		// Only check if the left side of the box can move
		if grid[newLoc[0]] == "]" {
			moved := [2]utils.Position{newLoc[0], newLoc[0].Move(dir, 1)}
			list = append(list, findMovedBoxes(grid, moved, dir)...)
		}
	case utils.Up, utils.Down:
		// Check if both sides of the box can move
		// First case is when alligning the box vertically
		if grid[newLoc[0]] == "[" {
			moved := [2]utils.Position{newLoc[0], newLoc[0].Move(utils.Right, 1)}
			list = append(list, findMovedBoxes(grid, moved, dir)...)
		}
		// Now check for overhanging
		if grid[newLoc[0]] == "]" {
			moved := [2]utils.Position{newLoc[0], newLoc[0].Move(utils.Left, 1)}
			list = append(list, findMovedBoxes(grid, moved, dir)...)
		}
		if grid[newLoc[1]] == "[" {
			moved := [2]utils.Position{newLoc[1], newLoc[1].Move(utils.Right, 1)}
			list = append(list, findMovedBoxes(grid, moved, dir)...)
		}
	}
	return list
}

func orderBoxPos(pos [2]utils.Position) [2]utils.Position {
	if pos[0].Col < pos[1].Col {
		return pos
	}
	return [2]utils.Position{pos[1], pos[0]}
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
