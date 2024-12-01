package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	list_1 := []int{}
	list_2 := []int{}

	for _, i := range input {
		number_1 := strings.Split(i, "   ")[0]
		number_2 := strings.Split(i, "   ")[1]

		list_1 = append(list_1, utils.IntegerOf(number_1))
		list_2 = append(list_2, utils.IntegerOf(number_2))
	}

	sort.Ints(list_1)
	sort.Ints(list_2)

	total := 0

	for i := 0; i < len(list_1); i++ {
		total += utils.Abs(list_2[i] - list_1[i])
	}
	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	list_1 := []int{}
	list_2 := []int{}

	for _, i := range input {
		number_1 := strings.Split(i, "   ")[0]
		number_2 := strings.Split(i, "   ")[1]

		list_1 = append(list_1, utils.IntegerOf(number_1))
		list_2 = append(list_2, utils.IntegerOf(number_2))
	}

	total := 0

	for i := 0; i < len(list_1); i++ {
		sub_total := 0
		current := list_1[i]
		for j := 0; j < len(list_1); j++ {
			comp := list_2[j]
			if current == comp {
				sub_total += 1
			}
		}
		total += sub_total * list_1[i]
	}
	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
