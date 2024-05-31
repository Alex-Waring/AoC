package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	pos := utils.NewPosition(0, 0)

	for _, line := range input {
		instructions := strings.Split(line, " ")
		direction := instructions[0]
		distance := utils.IntegerOf(instructions[1])

		switch direction {
		case "forward":
			pos = pos.Move(utils.Right, distance)
		case "up":
			pos = pos.Move(utils.Up, distance)
		case "down":
			pos = pos.Move(utils.Down, distance)
		}
	}
	fmt.Println(pos.Col * pos.Row)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	pos := utils.NewPosition(0, 0)
	aim := 0

	for _, line := range input {
		instructions := strings.Split(line, " ")
		direction := instructions[0]
		distance := utils.IntegerOf(instructions[1])

		switch direction {
		case "forward":
			pos = pos.Move(utils.Right, distance)
			pos = pos.Move(utils.Down, aim*distance)
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	fmt.Println(pos.Col * pos.Row)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
