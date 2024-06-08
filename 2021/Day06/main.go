package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type State map[int]int

func solve(input []string, days int) {
	defer utils.Timer("part1")()
	state := State{}

	// Load State
	for _, age := range strings.Split(input[0], ",") {
		i := utils.IntegerOf(age)
		state[i]++
	}

	for day := 1; day <= days; day++ {
		new_state := State{}
		for age, count := range state {
			if age == 0 {
				// Reset
				new_state[6] += count
				// Birth
				new_state[8] += count
			} else {
				// Age
				new_state[age-1] += count
			}
		}
		state = new_state
	}

	total := 0
	for _, count := range state {
		total += count
	}
	fmt.Println("Total: ", total)
}

func main() {
	input := utils.ReadInput("input.txt")
	solve(input, 80)
	solve(input, 256)
}
