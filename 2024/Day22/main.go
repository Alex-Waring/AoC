package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type CheatSheat struct {
	diffs  [2000]int
	values [2000]int
	totals map[string]int
}

func part1(input []string) {
	defer utils.Timer("part1")()

	total := 0

	for _, line := range input {
		initial := utils.IntegerOf(line)

		for i := 0; i < 2000; i++ {
			initial = mutate(initial)
		}

		total += initial
	}

	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	cheat_sheats := []CheatSheat{}

	for _, line := range input {
		initial := utils.IntegerOf(line)

		cheat_sheat := CheatSheat{
			totals: map[string]int{},
		}

		for i := 0; i < 2000; i++ {
			new := mutate(initial)
			cheat_sheat.diffs[i] = onesDigit(new) - onesDigit(initial)
			cheat_sheat.values[i] = onesDigit(new)
			initial = new
		}

		cheat_sheats = append(cheat_sheats, cheat_sheat)
	}

	for _, cheat_sheat := range cheat_sheats {
		for m := 3; m < 2000; m++ {
			c_i, c_j, c_k, c_l := cheat_sheat.diffs[m-3], cheat_sheat.diffs[m-2], cheat_sheat.diffs[m-1], cheat_sheat.diffs[m]
			key := fmt.Sprintf("%d,%d,%d,%d", c_i, c_j, c_k, c_l)

			if _, ok := cheat_sheat.totals[key]; !ok {
				cheat_sheat.totals[key] = cheat_sheat.values[m]
			}
		}
	}

	result := 0

	for i := -9; i < 10; i++ {
		fmt.Println(i)
		for j := -9; j < 10; j++ {
			for k := -9; k < 10; k++ {
				for l := -9; l < 10; l++ {
					total := 0

					for _, cheat_sheat := range cheat_sheats {
						if val, ok := cheat_sheat.totals[fmt.Sprintf("%d,%d,%d,%d", i, j, k, l)]; ok {
							total += val
						}
					}

					if total > result {
						result = total
					}
				}
			}
		}
	}
	fmt.Println(result)

}

var mutateCache map[int]int

func onesDigit(n int) int {
	return n % 10
}

func mutate(n int) int {
	if val, ok := mutateCache[n]; ok {
		return val
	}
	step1 := ((n * 64) ^ n) % 16777216
	step2 := ((step1 / 32) ^ step1) % 16777216
	step3 := ((step2 * 2048) ^ step2) % 16777216
	mutateCache[n] = step3
	return step3
}

func main() {
	input := utils.ReadInput("input.txt")
	mutateCache = map[int]int{}
	part1(input)
	part2(input)
}
