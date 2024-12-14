package main

import (
	"fmt"
	"math"

	"github.com/Alex-Waring/AoC/utils"
)

type Robot struct {
	pos  utils.Position
	velx int
	vely int
}

func part1(input []string) {
	defer utils.Timer("part1")()

	max_x := 101
	max_y := 103

	robots := []Robot{}

	for _, line := range input {
		var x, y, velx, vely int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &velx, &vely)
		robots = append(robots, Robot{pos: utils.Position{Row: y, Col: x}, velx: velx, vely: vely})
	}

	for i := 0; i < 100; i++ {
		for i, robot := range robots {
			new_pos := robot.pos.Slide(robot.vely, robot.velx)
			if new_pos.Row >= max_y {
				new_pos.Row -= max_y
			} else if new_pos.Row < 0 {
				new_pos.Row += max_y
			}
			if new_pos.Col >= max_x {
				new_pos.Col -= max_x
			} else if new_pos.Col < 0 {
				new_pos.Col += max_x
			}
			robots[i].pos = new_pos
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0

	middle_x := int(math.Ceil(float64(max_x)/2)) - 1
	middle_y := int(math.Ceil(float64(max_y)/2)) - 1

	for _, robot := range robots {
		if robot.pos.Row < middle_y && robot.pos.Col < middle_x {
			q1++
		} else if robot.pos.Row < middle_y && robot.pos.Col > middle_x {
			q2++
		} else if robot.pos.Row > middle_y && robot.pos.Col < middle_x {
			q3++
		} else if robot.pos.Row > middle_y && robot.pos.Col > middle_x {
			q4++

		}
	}
	fmt.Println(q1, q2, q3, q4)

	fmt.Println(q1 * q2 * q3 * q4)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	max_x := 101
	max_y := 103

	robots := []Robot{}

	for _, line := range input {
		var x, y, velx, vely int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &velx, &vely)
		robots = append(robots, Robot{pos: utils.Position{Row: y, Col: x}, velx: velx, vely: vely})
	}

	for i := 0; i < 10000; i++ {
		for i, robot := range robots {
			new_pos := robot.pos.Slide(robot.vely, robot.velx)
			if new_pos.Row >= max_y {
				new_pos.Row -= max_y
			} else if new_pos.Row < 0 {
				new_pos.Row += max_y
			}
			if new_pos.Col >= max_x {
				new_pos.Col -= max_x
			} else if new_pos.Col < 0 {
				new_pos.Col += max_x
			}
			robots[i].pos = new_pos
		}

		unique := true
		for y := range max_y {
			for x := range max_x {
				found := 0
				for _, robot := range robots {
					if robot.pos.Row == y && robot.pos.Col == x {
						found++
					}
				}
				if found > 1 {
					unique = false
				}
			}
		}
		if unique {
			fmt.Println(i + 1)
			break
		}
	}

	for y := range max_y {
		for x := range max_x {
			found := false
			for _, robot := range robots {
				if robot.pos.Row == y && robot.pos.Col == x {
					found = true
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
