package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func Sum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func main() {
	lines := utils.ReadInput("input.txt")

	powers := []int{}

	red_regex, _ := regexp.Compile("[0-9]* red")
	green_regex, _ := regexp.Compile("[0-9]* green")
	blue_regex, _ := regexp.Compile("[0-9]* blue")

	for _, input := range lines {
		input = strings.Split(input, ":")[1]
		rounds := strings.Split(input, ";")

		var red_cubes []int
		var green_cubes []int
		var blue_cubes []int

		for _, round := range rounds {
			red_string := red_regex.FindString(round)
			green_string := green_regex.FindString(round)
			blue_string := blue_regex.FindString(round)

			if red_string != "" {
				red_cubes = append(red_cubes, utils.IntegerOf(strings.Split(red_string, " ")[0]))
			}
			if green_string != "" {
				green_cubes = append(green_cubes, utils.IntegerOf(strings.Split(green_string, " ")[0]))
			}
			if blue_string != "" {
				blue_cubes = append(blue_cubes, utils.IntegerOf(strings.Split(blue_string, " ")[0]))
			}
		}
		power := slices.Max(red_cubes) * slices.Max(blue_cubes) * slices.Max(green_cubes)
		powers = append(powers, power)
	}
	fmt.Println(utils.Sum(powers))
}
