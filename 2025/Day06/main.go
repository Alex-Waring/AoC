package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

const PUZZLE_LENGTH = 4

func part1(input []string) {
	defer utils.Timer("part1")()
	numbers := [][PUZZLE_LENGTH]int{}
	operations := []rune{}
	for i, line := range input {
		if i < len(input)-1 {
			nums := strings.Fields(line)
			for j, num := range nums {
				if len(numbers) <= j {
					numbers = append(numbers, [PUZZLE_LENGTH]int{})
				}
				numbers[j][i] = utils.IntegerOf(num)
			}
		} else {
			ops := strings.Fields(line)
			for _, op := range ops {
				operations = append(operations, rune(op[0]))
			}
		}
	}

	total := 0
	for i, op := range operations {
		switch op {
		case '+':
			miniTotal := 0
			for j := 0; j < PUZZLE_LENGTH; j++ {
				miniTotal += numbers[i][j]
			}
			total += miniTotal
		case '*':
			prod := 1
			for j := 0; j < PUZZLE_LENGTH; j++ {
				prod *= numbers[i][j]
			}
			total += prod
		}
	}
	fmt.Println("Part 1:", total)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	numbers := [][PUZZLE_LENGTH]string{}
	puzzleSizes := []int{}

	// parse the operations to also get the max length of each column
	operationsString := input[len(input)-1]
	operations := strings.Fields(operationsString)
	prepareForLengths := strings.ReplaceAll(operationsString, "*", "+")
	columnSizes := strings.Split(prepareForLengths, "+")
	for _, col := range columnSizes {
		puzzleSizes = append(puzzleSizes, len(col))
	}

	// Manual tweaking due to how my IDE handles the files, remove the first one and set the last one to 3
	puzzleSizes = puzzleSizes[1:]
	puzzleSizes[len(puzzleSizes)-1] = 3

	// Nunmbers pulled from puzzleSizes[i] then space
	for i := 0; i < len(input)-1; i++ {
		line := input[i] + "    " // add spaces to handle vscode trimming lines
		pos := 0
		for j := 0; j < len(puzzleSizes); j++ {
			size := puzzleSizes[j]
			num := line[pos : pos+size]
			pos += size + 1 // plus one for the space
			if len(numbers) <= j {
				numbers = append(numbers, [PUZZLE_LENGTH]string{})
			}
			numbers[j][i] = num
		}

	}

	// Convert the numbers into columns
	actualNumbers := [][]int{}
	for i := range numbers {
		actualNumbers = append(actualNumbers, []int{})
		length := len(numbers[i][0])
		for pos := length - 1; pos >= 0; pos-- {
			newNum := ""
			for j := PUZZLE_LENGTH - 1; j >= 0; j-- {
				newNum = string(numbers[i][j][pos]) + newNum
			}
			actualNumbers[i] = append(actualNumbers[i], utils.IntegerOf(strings.TrimSpace(newNum)))
		}
	}

	// Do the operations
	total := 0
	for i, op := range operations {
		switch op {
		case "+":
			sum := 0
			for _, num := range actualNumbers[i] {
				sum += num
			}
			total += sum
		case "*":
			prod := 1
			for _, num := range actualNumbers[i] {
				prod *= num
			}
			total += prod
		}
	}
	fmt.Println("Part 2:", total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
