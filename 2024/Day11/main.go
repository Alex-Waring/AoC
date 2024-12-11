package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	data := utils.StringToIntList(strings.Fields(input[0]))
	fmt.Println(solve(data, 25))
}

func part2(input []string) {
	defer utils.Timer("part2")()

	data := utils.StringToIntList(strings.Fields(input[0]))
	fmt.Println(solve(data, 75))
}

func solve(data []int, blinks int) int {
	stones := map[int]int{}

	for _, v := range data {
		stones[v] = 1
	}

	for range blinks {
		stones = calc(stones)
	}

	total := 0
	for _, v := range stones {
		total += v
	}
	return total
}

func calc(stones map[int]int) map[int]int {
	output := map[int]int{}

	for stone, count := range stones {
		if stone == 0 {
			addStone(output, 1, count)
		} else if len(fmt.Sprintf("%d", stone))%2 == 0 {
			nums := split(stone)
			addStone(output, nums[0], count)
			addStone(output, nums[1], count)
		} else {
			addStone(output, stone*2024, count)
		}
	}
	return output
}

func addStone(stones map[int]int, engraved int, toAdd int) map[int]int {
	if _, ok := stones[engraved]; !ok {
		stones[engraved] = toAdd
	} else {
		stones[engraved] += toAdd
	}
	return stones
}

func split(num int) []int {
	str := fmt.Sprintf("%d", num)
	half := len(str) / 2

	num1 := str[:half]
	num2 := str[half:]

	return []int{utils.IntegerOf(num1), utils.IntegerOf(num2)}
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
