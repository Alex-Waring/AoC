package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/Alex-Waring/AoC/utils"
)

type Location struct {
	map_value        rune
	empty_row        bool
	empty_column     bool
	parsed_map_value rune
}

func generatePairs(list [][]int) [][][]int {
	return_list := [][][]int{}
	for i := 0; i < len(list); i++ {
		for x := i; x < len(list); x++ {
			if x != i {
				return_list = append(return_list, [][]int{list[i], list[x]})
			}
		}
	}
	return return_list
}

func distance(universe map[int]map[int]Location, start []int, end []int) int {
	start_x, start_y := start[1], start[0]
	end_x, end_y := end[1], end[0]

	distance := math.Abs(float64(start_x)-float64(end_x)) + math.Abs(float64(start_y)-float64(end_y))

	x_normaliser := 0
	if start_x > end_x {
		x_normaliser = -1
	} else {
		x_normaliser = 1
	}

	y_normaliser := 0
	if start_y > end_y {
		y_normaliser = -1
	} else {
		y_normaliser = 1
	}

	for x := start_x; x != end_x; x = x + x_normaliser {
		if universe[start_y][x+x_normaliser].empty_column {
			distance += 999999
		}
	}

	for y := start_y; y != end_y; y = y + y_normaliser {
		if universe[y+y_normaliser][start_x].empty_row {
			distance += 999999
		}
	}

	return int(distance)
}

func main() {
	lines := utils.ReadInput("input.txt")

	universe := map[int]map[int]Location{}
	planets := [][]int{}

	for y, line := range lines {
		for x, location := range line {
			if _, ok := universe[y]; !ok {
				universe[y] = make(map[int]Location)
			}
			universe[y][x] = Location{map_value: location, empty_row: false, empty_column: false, parsed_map_value: location}
		}
	}

	empty_row := utils.SliceFilledWithRune(len(universe), '.')
	rows := len(universe)
	for y := 0; y < rows; y++ {
		row := []rune{}
		for x := 0; x < len(universe[y]); x++ {
			row = append(row, universe[y][x].map_value)
		}
		if slices.Equal(row, empty_row) {
			for x := 0; x < len(universe[y]); x++ {
				location := universe[y][x]
				location.parsed_map_value = 'e'
				location.empty_row = true
				universe[y][x] = location
			}
		}
	}

	empty_column := utils.SliceFilledWithRune(len(universe), '.')
	columns := len(universe[0])

	for x := 0; x < columns; x++ {
		column := []rune{}
		for y := 0; y < len(universe); y++ {
			column = append(column, universe[y][x].map_value)
		}
		if slices.Equal(column, empty_column) {
			for y := 0; y < len(universe); y++ {
				location := universe[y][x]
				location.parsed_map_value = 'e'
				location.empty_column = true
				universe[y][x] = location
			}
		}
	}

	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {
			if universe[y][x].map_value == '#' {
				planets = append(planets, []int{y, x})
			}
		}
	}

	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {
			fmt.Print(string(universe[y][x].parsed_map_value))
		}
		fmt.Println()
	}

	total_distance := 0
	for _, pair := range generatePairs(planets) {
		distance := distance(universe, pair[0], pair[1])
		total_distance += int(distance)
	}
	fmt.Println(total_distance)
}
