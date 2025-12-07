package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func digitList(line string) []int {
	digits := []int{}
	for _, char := range line {
		digit := int(char - '0')
		digits = append(digits, digit)
	}
	return digits
}

func findMaxAndRecordIndex(arr []int) (int, int) {
	max := arr[0]
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}

	return max, maxIndex
}

func part1(input []string) {
	defer utils.Timer("part1")()

	results := []int{}

	for _, line := range input {
		// convert to list of digits
		digits := digitList(line)
		// Find first digit, max of all but the last one
		firstDigit, index := findMaxAndRecordIndex(digits[:len(digits)-1])
		// Find max of all to the right
		secondDigit, _ := findMaxAndRecordIndex(digits[index+1:])
		number := utils.IntegerOf(fmt.Sprintf("%d%d", firstDigit, secondDigit))
		results = append(results, number)
	}

	fmt.Println(utils.Sum(results))
}

func part2(input []string) {
	defer utils.Timer("part2")()

	results := []int{}

	for _, line := range input {
		// convert to list of digits
		digits := digitList(line)
		numbers := []int{}

		// Collect all but last digit
		for i := 12; i > 1; i-- {
			digitsToConsider := digits[:len(digits)-i+1]
			foundDigit, index := findMaxAndRecordIndex(digitsToConsider)
			digits = digits[index+1:]
			numbers = append(numbers, foundDigit)
		}
		// Get last one
		lastDigit, _ := findMaxAndRecordIndex(digits)
		numbers = append(numbers, lastDigit)

		numberString := ""
		for _, num := range numbers {
			numberString += fmt.Sprintf("%d", num)
		}
		number := utils.IntegerOf(numberString)
		results = append(results, number)
	}

	fmt.Println(utils.Sum(results))
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
