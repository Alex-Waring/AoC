package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/Alex-Waring/AoC/utils"
)

func main() {
	lines := utils.ReadInput("input.txt")

	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	_lines := []string{}
	for _, line := range lines {
		_numbers := []string{}

		for index, charecter := range line {
			if unicode.IsNumber(charecter) {
				_numbers = append(_numbers, string(charecter))
			}
			for word, number := range numbers {
				if strings.HasPrefix(line[index:], word) {
					_numbers = append(_numbers, number)
				}
			}
		}
		_lines = append(_lines, (_numbers[0] + _numbers[len(_numbers)-1]))
	}

	var answer int
	for _, line := range _lines {
		_num, _ := strconv.Atoi(line)
		answer += _num
	}
	fmt.Print(answer)

}
