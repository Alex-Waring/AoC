package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func add(a, b int) int {
	r := a + b
	return r % 100
}

func sub(a, b int) int {
	r := a - b
	r = r % 100
	if r < 0 {
		r += 100
	}
	return r
}

func addWithCount(a, b int) (int, int) {
	count := 0
	for i := 0; i < b; i++ {
		a = a + 1
		if a == 100 {
			a = 0
			count++
		}
	}
	return a, count
}

func subWithCount(a, b int) (int, int) {
	count := 0
	for i := 0; i < b; i++ {
		a = a - 1
		if a == -1 {
			a = 99
		}
		if a == 0 {
			count++
		}
	}
	return a, count
}

func part1(input []string) {
	defer utils.Timer("part1")()
	initialPos := 50
	pos := initialPos
	timesAt0 := 0

	for _, line := range input {
		var move int
		dir := string(line[0])
		fmt.Sscanf(line[1:], "%d", &move)

		if dir == "L" {
			pos = sub(pos, move)
		} else if dir == "R" {
			pos = add(pos, move)
		}

		if pos == 0 {
			timesAt0++
		}
	}
	fmt.Println("Final position:", pos)
	fmt.Println("Times at 0:", timesAt0)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	initialPos := 50
	pos := initialPos
	timesAt0 := 0

	for _, line := range input {
		var move int
		dir := string(line[0])
		fmt.Sscanf(line[1:], "%d", &move)

		if dir == "L" {
			newPos, count := subWithCount(pos, move)
			pos = newPos
			timesAt0 += count
		} else if dir == "R" {
			newPos, count := addWithCount(pos, move)
			pos = newPos
			timesAt0 += count
		}
	}
	fmt.Println("Final position:", pos)
	fmt.Println("Times at 0:", timesAt0)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
