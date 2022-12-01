package main

import (
	"fmt"
	"unicode"

	"github.com/Alex-Waring/AoC/utils"
)

func findNumber(line string, x int) int {
	number := 0
	for x > 0 && unicode.IsDigit(rune(line[x-1])) {
		x--
	}

	for x < len(line) && unicode.IsDigit(rune(line[x])) {
		number = (number * 10) + int(line[x]-'0')
		x++
	}
	return number
}

func main() {
	lines := utils.ReadInput("Input.txt")
	parts := []int{}

	delta := []int{-1, 0, 1}
	number := 0
	part := false

	for y, line := range lines {
		if part && number != 0 {
			parts = append(parts, number)
		}
		part = false
		number = 0

		for x, charecter := range line {
			if unicode.IsDigit(charecter) {
				// shunt number and add new number
				number = (number * 10) + int(charecter-'0')

				for _, y_diff := range delta {
					for _, x_diff := range delta {
						y2 := y + y_diff
						x2 := x + x_diff

						if x2 >= 0 && x2 < len(line) && y2 >= 0 && y2 < len(lines) {
							checked_charecter := lines[y2][x2]
							if string(checked_charecter) != "." && !unicode.IsDigit(rune(checked_charecter)) {
								part = true
							}
						}
					}
				}
			} else {
				if part {
					parts = append(parts, number)
				}
				part = false
				number = 0
			}
		}
	}
	println(utils.Sum(parts))

	ratios := []int{}
	for y, line := range lines {
		for x, charecter := range line {
			if string(charecter) != "*" {
				continue
			}

			set := map[int]bool{}

			for _, y_diff := range delta {
				for _, x_diff := range delta {
					y2 := y + y_diff
					x2 := x + x_diff

					if x2 >= 0 && x2 < len(line) && y2 >= 0 && y2 < len(lines) {
						if unicode.IsDigit(rune(lines[y2][x2])) {
							num := findNumber(lines[y2], x2)
							set[num] = true
						}
					}
				}
			}
			fmt.Println(set)
			if len(set) == 2 {
				list := []int{}
				for key := range set {
					list = append(list, key)
				}
				ratios = append(ratios, list[0]*list[1])
			}
		}
	}
	println(utils.Sum(ratios))
}
