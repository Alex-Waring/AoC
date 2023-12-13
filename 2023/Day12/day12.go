package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type GearSet struct {
	damaged_springs []string
	input           string
}

type Solver struct {
	cache map[string]int
}

func createCacheString(input string, damaged_springs []string) string {
	key := input

	for _, spring := range damaged_springs {
		key += ","
		key += spring
	}
	return key
}

func (s *Solver) solve(input string, damaged_springs []string, index int) int {

	cache_key := createCacheString(input, damaged_springs)
	if v, ok := s.cache[cache_key]; ok {
		return v
	}

	// Debug blocks
	// if index > 17 {
	// 	fmt.Print("Solving " + input + " for ")
	// 	fmt.Println(damaged_springs)
	// }

	// If there is no input left, return 1 if groups are empty, or return 0
	// This is the catch block when we have finished solving
	if len(input) == 0 {
		if len(damaged_springs) == 0 {
			s.cache[cache_key] = 1
			// No input and no groups means a win!
			return 1
		} else {
			s.cache[cache_key] = 0
			// We have no input but there are groups left, fail
			return 0
		}
	}

	// If the length of the input is less than the sum of the groups, it's not possible
	if len(input) < utils.SumListString(damaged_springs) {
		return 0
	}

	// If it starts with a ., strip that dot
	if strings.HasPrefix(input, ".") {
		result := s.solve(strings.TrimPrefix(input, "."), damaged_springs, index)
		s.cache[cache_key] = result
		return result
	}

	// If it starts with a ? then try both . and #
	if strings.HasPrefix(input, "?") {
		// This is the branch
		result := s.solve(strings.Replace(input, "?", "#", 1), damaged_springs, index) + s.solve(strings.Replace(input, "?", ".", 1), damaged_springs, index)
		s.cache[cache_key] = result
		return result
	}

	// Start finding the gear set
	if strings.HasPrefix(input, "#") {
		// Shouldn't really be here, returning zero for cases where it's not a possible set
		// Option 1 is we have no more groups left, we know we have a gear left in the input so this is a fail
		// Option 2 is if the length of the input is less than the first group, we shouldn't hit this because
		// of line 56 but it's a good catch
		if (len(damaged_springs) == 0) || (len(input) < utils.IntegerOf(damaged_springs[0])) {
			s.cache[cache_key] = 0
			return 0
		}

		// If there isn't enough room at the start for the first gear, return zero
		if strings.Contains(input[0:utils.IntegerOf(damaged_springs[0])], ".") {
			s.cache[cache_key] = 0
			return 0
		}

		// If we have more than one set left, solve the first set
		if len(damaged_springs) > 1 {
			// Catch if there wouldn't be enough space left for another set
			if len(input) < utils.IntegerOf(damaged_springs[0])+1 {
				s.cache[cache_key] = 0
				return 0
			}
			// Catch if the group we are tring to match would end in a # (there are no overlaps)
			if input[utils.IntegerOf(damaged_springs[0])] == '#' {
				s.cache[cache_key] = 0
				return 0
			}

			// Remove the set we matched and the first group and keep solving
			result := s.solve(input[utils.IntegerOf(damaged_springs[0])+1:], damaged_springs[1:], index)
			s.cache[cache_key] = result
			return result
		} else {
			// We have one group left, remove it and solve again
			result := s.solve(input[utils.IntegerOf(damaged_springs[0]):], damaged_springs[1:], index)
			s.cache[cache_key] = result
			return result
		}
	}
	panic("Should not be here")
}

func part1(spring_list []GearSet) {
	defer utils.Timer("part1")()
	total_arrangements := 0
	s := Solver{make(map[string]int)}
	for index, gearSet := range spring_list {
		total_arrangements += s.solve(gearSet.input, gearSet.damaged_springs, index)
	}
	fmt.Println(total_arrangements)
}

func part2(spring_list []GearSet) {
	defer utils.Timer("part2")()

	unfolded_spring_list := []GearSet{}

	for _, set := range spring_list {
		// At this point this is just easier
		new_input := set.input + "?" + set.input + "?" + set.input + "?" + set.input + "?" + set.input
		new_damaged_springs := set.damaged_springs
		for i := 0; i <= 3; i++ {
			new_damaged_springs = append(new_damaged_springs, set.damaged_springs...)
		}
		unfolded_spring_list = append(unfolded_spring_list, GearSet{input: new_input, damaged_springs: new_damaged_springs})
	}

	total_arrangements := 0
	s := Solver{make(map[string]int)}
	for index, gearSet := range unfolded_spring_list {
		total_arrangements += s.solve(gearSet.input, gearSet.damaged_springs, index)
	}
	fmt.Println(total_arrangements)
}

func main() {
	input := utils.ReadInput("input.txt")
	spring_list := []GearSet{}

	for _, line := range input {
		set := GearSet{
			input: strings.Split(line, " ")[0], damaged_springs: strings.Split(strings.Split(line, " ")[1], ","),
		}
		spring_list = append(spring_list, set)
	}
	part1(spring_list)
	part2(spring_list)
}
