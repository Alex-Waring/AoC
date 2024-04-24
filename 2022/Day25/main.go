package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	total := "0"
	for _, line := range input {
		total = addSnafu(total, line)
	}
	fmt.Println(total)

}

func part2() {
	defer utils.Timer("part2")()
}

func addSnafu(a string, b string) string {
	// Pad the string
	for len(a) != len(b) {
		if len(a) < len(b) {
			a = "0" + a
		} else {
			b = "0" + b
		}
	}

	a = utils.Reverse(a)
	b = utils.Reverse(b)
	total := ""

	remainder := "0"
	// Add them together
	for i := 0; i < len(a); i++ {
		aChar := string(a[i])
		bChar := string(b[i])
		t, r := addChars(aChar, bChar, remainder)
		remainder = r
		total += t
	}
	if remainder != "0" {
		total = remainder + total
	}
	return utils.Reverse(total)
}

func addChars(a string, b string, r string) (string, string) {
	// Convert to int
	decimalConverter := map[string]int{
		"=": -2,
		"-": -1,
		"0": 0,
		"1": 1,
		"2": 2,
	}
	snafuConverter := map[int]string{
		-2: "=",
		-1: "-",
		0:  "0",
		1:  "1",
		2:  "2",
	}

	// Add the numbers
	total := decimalConverter[a] + decimalConverter[b] + decimalConverter[r]

	// Check if we're over
	if total > 2 {
		return snafuConverter[total-5], "1"
	}
	// Check if we're under
	if total < -2 {
		return snafuConverter[total+5], "-"
	}
	// Otherwise we're in the clear
	return snafuConverter[total], "0"

}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
}
