package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func difference(numbers []int) ([]int, bool) {
	set := map[int]bool{}
	differences := []int{}

	for i := 1; i < len(numbers); i++ {
		difference := numbers[i] - numbers[i-1]
		differences = append(differences, difference)
		set[difference] = true
	}

	if reflect.DeepEqual(map[int]bool{0: true}, set) {
		return differences, true
	} else {
		return differences, false
	}
}

func part1(lines []string) {
	defer utils.Timer("part1")()
	new_values := []int{}

	for _, line := range lines {
		numbers := strings.Fields(line)
		numbers_int := []int{}
		tree := [][]int{}

		for _, number := range numbers {
			numbers_int = append(numbers_int, utils.IntegerOf(number))
		}
		tree = append(tree, numbers_int)

		zeros := false
		for !zeros {
			next_differences := []int{}
			next_differences, zeros = difference(tree[len(tree)-1])
			tree = append(tree, next_differences)
		}

		tree[len(tree)-1] = append(tree[len(tree)-1], 0)
		for i := len(tree) - 2; i >= 0; i-- {
			value_to_add := tree[i+1][len(tree[i+1])-1]
			new_value := tree[i][len(tree[i])-1] + value_to_add
			tree[i] = append(tree[i], new_value)
		}
		new_values = append(new_values, tree[0][len(tree[0])-1])
	}
	fmt.Println(utils.Sum(new_values))
}

func part2(lines []string) {
	defer utils.Timer("part2")()
	part2_values := []int{}

	for _, line := range lines {
		numbers := strings.Fields(line)
		numbers_int := []int{}
		tree := [][]int{}

		for _, number := range numbers {
			numbers_int = append(numbers_int, utils.IntegerOf(number))
		}
		tree = append(tree, numbers_int)

		zeros := false
		for !zeros {
			next_differences := []int{}
			next_differences, zeros = difference(tree[len(tree)-1])
			tree = append(tree, next_differences)
		}

		tree[len(tree)-1] = append(tree[len(tree)-1], 0)
		for i := len(tree) - 2; i >= 0; i-- {
			value_to_add := tree[i+1][0]
			new_value := tree[i][0] - value_to_add
			tree[i] = append([]int{new_value}, tree[i]...)
		}
		part2_values = append(part2_values, tree[0][0])
	}
	fmt.Println(utils.Sum(part2_values))
}

func main() {
	lines := utils.ReadInput("input.txt")
	part1(lines)
	part2(lines)
}
