package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	keys := [][5]int{}
	locks := [][5]int{}

	for i := 0; i < len(input); i = i + 8 {
		if input[i] == "....." {
			keys = append(keys, parseKeys([6]string{input[i+1], input[i+2], input[i+3], input[i+4], input[i+5], input[i+6]}))
		} else {
			locks = append(locks, parseLocks([6]string{input[i+1], input[i+2], input[i+3], input[i+4], input[i+5], input[i+6]}))
		}
	}

	possible := 0

	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(locks); j++ {
			key_to_try := keys[i]
			lock_to_try := locks[j]

			lock_fits := true

			for k := 0; k < 5; k++ {
				if key_to_try[k]+lock_to_try[k] > 5 {
					lock_fits = false
					break
				}
			}
			if lock_fits {
				possible += 1
			}
		}
	}
	fmt.Println(possible)

}

func parseLocks(input [6]string) [5]int {
	returnValue := [5]int{}
	for i := 0; i < 5; i++ {
		returnValue[i] = 0
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			if input[i][j] == '#' {
				returnValue[j] += 1
			}
		}
	}
	return returnValue
}

func parseKeys(input [6]string) [5]int {
	returnValue := [5]int{}
	for i := 0; i < 5; i++ {
		returnValue[i] = 0
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if input[i][j] == '#' {
				returnValue[j] += 1
			}
		}
	}
	return returnValue
}

func part2(input []string) {
	defer utils.Timer("part2")()
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
