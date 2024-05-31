package main

import (
	"fmt"
	"strconv"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	bitCounter := []int{}

	for i := 0; i < len(input[0]); i++ {
		bitCounter = append(bitCounter, 0)
	}

	for _, line := range input {
		for i, bit := range line {
			if bit == '1' {
				bitCounter[i]++
			} else {
				bitCounter[i]--
			}
		}
	}

	gamma := ""
	epsilon := ""

	for _, bit := range bitCounter {
		if bit > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gamma_int, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_int, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gamma_int * epsilon_int)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	oxygenList := utils.DuplicateList(input)
	scrubberList := utils.DuplicateList(input)

	oxygen := oxygenCounter(oxygenList)
	scrubber := scrubberCounter(scrubberList)

	fmt.Println(oxygen * scrubber)

}

func oxygenCounter(input []string) int {
	for i := 0; i < len(input[0]); i++ {
		bitCounter := 0
		for _, line := range input {
			if line[i] == '1' {
				bitCounter++
			} else {
				bitCounter--
			}
		}
		newList := []string{}
		for _, line := range input {
			if bitCounter >= 0 && line[i] == '1' {
				newList = append(newList, line)
			} else if bitCounter < 0 && line[i] == '0' {
				newList = append(newList, line)
			}
		}
		input = newList
		if len(input) == 1 {
			result, _ := strconv.ParseInt(input[0], 2, 64)
			return int(result)
		}
	}
	panic("No solution found")
}

func scrubberCounter(input []string) int {
	for i := 0; i < len(input[0]); i++ {
		bitCounter := 0
		for _, line := range input {
			if line[i] == '1' {
				bitCounter--
			} else {
				bitCounter++
			}
		}
		newList := []string{}
		for _, line := range input {
			if bitCounter > 0 && line[i] == '1' {
				newList = append(newList, line)
			} else if bitCounter <= 0 && line[i] == '0' {
				newList = append(newList, line)
			}
		}
		input = newList
		if len(input) == 1 {
			result, _ := strconv.ParseInt(input[0], 2, 64)
			return int(result)
		}
	}
	panic("No solution found")
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
