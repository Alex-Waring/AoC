package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	positions := utils.StringToIntList(strings.Split(input[0], ","))

	// The median is the optiomal position to minimize the sum of distances
	// https://en.wikipedia.org/wiki/Median#Optimality_property
	target := utils.Median(positions)

	distance := 0
	for _, p := range positions {
		distance += utils.Abs(target - p)
	}
	fmt.Println("Total distance: ", distance)

}

func part2(input []string) {
	defer utils.Timer("part2")()
	positions := utils.StringToIntList(strings.Split(input[0], ","))

	// The mean is probably the optimal position as the distance is a triangular number
	// (n^2 + n) / 2
	// If it was just n^2, the median would be optimal but it's close enough
	low_target := int(math.Floor(utils.Mean(positions)))
	high_target := int(math.Ceil(utils.Mean(positions)))

	low_distance := 0
	high_distance := 0
	for _, p := range positions {
		low_hops := utils.Abs(low_target - p)
		low_distance += (low_hops*low_hops + low_hops) / 2
		high_hops := utils.Abs(high_target - p)
		high_distance += (high_hops*high_hops + high_hops) / 2
	}
	// We've calculated the low and high distances by rounding up and down, one will be right
	fmt.Println("Total distance: ", min(low_distance, high_distance))
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
