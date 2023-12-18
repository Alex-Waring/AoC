package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string, lagoon map[utils.Position]bool) {
	position := utils.NewPosition(0, 0)
	lagoon[position] = true

	rows := 0
	columns := 0
	rows_start := 0
	columns_start := 0

	for _, instruction := range input {
		direction := strings.Split(instruction, " ")[0]
		distance := strings.Split(instruction, " ")[1]
		// colour := strings.Split(instruction, " ")[2]

		var standard_direction utils.Direction

		switch direction {
		case "U":
			standard_direction = utils.Up
		case "D":
			standard_direction = utils.Down
		case "L":
			standard_direction = utils.Left
		case "R":
			standard_direction = utils.Right
		}

		for i := 0; i < utils.IntegerOf(distance); i++ {
			position = position.Move(standard_direction, 1)
			rows = max(rows, position.Row)
			columns = max(columns, position.Col)
			columns_start = min(columns_start, position.Col)
			rows_start = min(rows_start, position.Row)
			lagoon[position] = true
		}
	}

	// At no point is an edge next to another unrelated edge, so we can ignore direction
	scanned_map := make(map[utils.Position]bool)
	for y := rows_start; y <= rows; y++ {
		crossed := 0
		for x := columns_start; x <= columns; x++ {
			draw_position := utils.NewPosition(y, x)
			previous_poition := utils.NewPosition(y, x-1)
			next_position := utils.NewPosition(y, x+1)
			above := utils.NewPosition(y-1, x)
			below := utils.NewPosition(y+1, x)

			// Crosses when .###.##..#...#
			// gives values 00011122233334

			if lagoon[draw_position] && lagoon[above] && lagoon[below] {
				crossed++
			} else if lagoon[previous_poition] && lagoon[draw_position] && lagoon[below] && !lagoon[next_position] {
				crossed++
			} else if !lagoon[previous_poition] && lagoon[draw_position] && lagoon[below] && lagoon[next_position] {
				crossed++
			} else if !lagoon[draw_position] && crossed%2 == 1 {
				scanned_map[draw_position] = true
			}
			fmt.Print(crossed)
		}
		fmt.Println()
	}

	total := 0
	for y := rows_start; y <= rows; y++ {
		for x := columns_start; x <= columns; x++ {
			draw_position := utils.NewPosition(y, x)
			if scanned_map[draw_position] || lagoon[draw_position] {
				fmt.Print("#")
				total++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println(total)
}

func part2(input []string, lagoon map[utils.Position]bool) {
	position := utils.NewPosition(0, 0)
	lagoon[position] = true

	rows := 0
	columns := 0
	rows_start := 0
	columns_start := 0

	for _, instruction := range input {
		colour := strings.Split(instruction, " ")[2]
		colour = strings.Trim(colour, "(")
		colour = strings.Trim(colour, ")")
		direction := colour[len(colour)-1:]

		var standard_direction utils.Direction

		switch direction {
		case "3":
			standard_direction = utils.Up
		case "1":
			standard_direction = utils.Down
		case "2":
			standard_direction = utils.Left
		case "0":
			standard_direction = utils.Right
		}

		distance, _ := strconv.ParseInt(colour[1:6], 16, 64)
		int_distance := int(distance)

		for i := 0; i < int_distance; i++ {
			position = position.Move(standard_direction, 1)
			rows = max(rows, position.Row)
			columns = max(columns, position.Col)
			columns_start = min(columns_start, position.Col)
			rows_start = min(rows_start, position.Row)
			lagoon[position] = true
		}
	}

	// At no point is an edge next to another unrelated edge, so we can ignore direction
	scanned_map := make(map[utils.Position]bool)
	for y := rows_start; y <= rows; y++ {
		crossed := 0
		for x := columns_start; x <= columns; x++ {
			draw_position := utils.NewPosition(y, x)
			previous_poition := utils.NewPosition(y, x-1)
			next_position := utils.NewPosition(y, x+1)
			above := utils.NewPosition(y-1, x)
			below := utils.NewPosition(y+1, x)

			// Crosses when .###.##..#...#
			// gives values 00011122233334

			if lagoon[draw_position] && lagoon[above] && lagoon[below] {
				crossed++
			} else if lagoon[previous_poition] && lagoon[draw_position] && lagoon[below] && !lagoon[next_position] {
				crossed++
			} else if !lagoon[previous_poition] && lagoon[draw_position] && lagoon[below] && lagoon[next_position] {
				crossed++
			} else if !lagoon[draw_position] && crossed%2 == 1 {
				scanned_map[draw_position] = true
			}
		}
		fmt.Println()
	}

	total := 0
	for y := rows_start; y <= rows; y++ {
		for x := columns_start; x <= columns; x++ {
			draw_position := utils.NewPosition(y, x)
			if scanned_map[draw_position] || lagoon[draw_position] {
				total++
			}
		}
	}
	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")

	lagoon := make(map[utils.Position]bool)

	part2(input, lagoon)
}
