package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	increased := 0

	for i := 1; i < len(input); i++ {
		if utils.IntegerOf(input[i]) > utils.IntegerOf(input[i-1]) {
			increased++
		}
	}
	fmt.Println(increased)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	increased := 0

	for i := 1; i < len(input)-2; i++ {
		currentWindow := utils.IntegerOf(input[i]) + utils.IntegerOf(input[i+1]) + utils.IntegerOf(input[i+2])
		previousWindow := utils.IntegerOf(input[i-1]) + utils.IntegerOf(input[i]) + utils.IntegerOf(input[i+1])

		if currentWindow > previousWindow {
			increased++
		}
	}
	fmt.Println(increased)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
