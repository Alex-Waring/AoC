package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func stringify(rock_map map[int]map[int]string) string {
	string_key := ""

	for y := 0; y < len(rock_map); y++ {
		for x := 0; x < len(rock_map[0]); x++ {
			string_key += rock_map[y][x]
		}
	}
	return string_key
}

func tilt(rock_map map[int]map[int]string, y_tilt int, x_tilt int) map[int]map[int]string {
	if y_tilt != 0 {
		for x := 0; x < len(rock_map[0]); x++ {
			var y int
			if y_tilt == 1 {
				y = 0
			} else {
				y = len(rock_map) - 1
			}

			for y < len(rock_map) && y >= 0 {
				if y == 0 && y_tilt == 1 {
					y += y_tilt
				} else if rock_map[y][x] == "O" && rock_map[y-y_tilt][x] == "." {
					rock_map[y][x] = "."
					rock_map[y-y_tilt][x] = "O"
					y -= y_tilt
				} else {
					y += y_tilt
				}
			}
		}
	} else {
		for y := 0; y < len(rock_map); y++ {
			var x int
			if x_tilt == 1 {
				x = 0
			} else {
				x = len(rock_map[0]) - 1
			}

			for x < len(rock_map[0]) && x >= 0 {
				if x == 0 && x_tilt == 1 {
					x += x_tilt
				} else if rock_map[y][x] == "O" && rock_map[y][x-x_tilt] == "." {
					rock_map[y][x] = "."
					rock_map[y][x-x_tilt] = "O"
					x -= x_tilt
				} else {
					x += x_tilt
				}
			}
		}
	}
	return rock_map
}

func part1(lines []string) {
	defer utils.Timer("part1")()

	rock_map := map[int]map[int]string{}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if _, ok := rock_map[y]; !ok {
				rock_map[y] = make(map[int]string)
			}
			rock_map[y][x] = string(lines[y][x])
		}
	}

	rock_map = tilt(rock_map, 1, 0)

	total := 0
	for y := 0; y < len(rock_map); y++ {
		for x := 0; x < len(rock_map[0]); x++ {
			if rock_map[y][x] == "O" {
				total += len(rock_map) - y
			}
		}
	}
	fmt.Println(total)
}

func part2(lines []string) {
	defer utils.Timer("part2")()

	rock_map := map[int]map[int]string{}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if _, ok := rock_map[y]; !ok {
				rock_map[y] = make(map[int]string)
			}
			rock_map[y][x] = string(lines[y][x])
		}
	}

	memory := []string{stringify(rock_map)}
	cycle_found := false
	loop_cycle := 0
	for cycle := 0; cycle < 1000000000 && !cycle_found; cycle++ {
		rock_map = tilt(rock_map, 1, 0)
		rock_map = tilt(rock_map, 0, 1)
		rock_map = tilt(rock_map, -1, 0)
		rock_map = tilt(rock_map, 0, -1)

		rock_map_string := stringify(rock_map)
		if utils.StringInSlice(rock_map_string, memory) {
			fmt.Print("Cycle found at ")
			fmt.Println(cycle)
			loop_cycle = cycle
			cycle_found = true
		} else {
			memory = append(memory, rock_map_string)
		}

		if cycle%1000 == 0 {
			fmt.Println(cycle)
		}
	}

	fmt.Println(loop_cycle)
	cycles_remaining := 1000000000 % loop_cycle
	for cycle := 1000000000 - cycles_remaining + 1; cycle <= 1000000000 && !cycle_found; cycle++ {
		rock_map = tilt(rock_map, 1, 0)
		rock_map = tilt(rock_map, 0, 1)
		rock_map = tilt(rock_map, -1, 0)
		rock_map = tilt(rock_map, 0, -1)
	}

	total := 0
	for y := 0; y < len(rock_map); y++ {
		for x := 0; x < len(rock_map[0]); x++ {
			if rock_map[y][x] == "O" {
				total += len(rock_map) - y
			}
		}
	}
	fmt.Println(total)
}

func main() {
	lines := utils.ReadInput("input.txt")

	part1(lines)
	part2(lines)
}
