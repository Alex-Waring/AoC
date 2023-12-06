package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1() {
	defer utils.Timer("part1")()
	lines := utils.ReadInput("input.txt")

	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distance := strings.Fields(strings.Split(lines[1], ":")[1])

	victories := []int{}
	for i := 0; i < len(times); i++ {
		victory := 0
		time_int := utils.IntegerOf(times[i])
		distance_int := utils.IntegerOf(distance[i])
		for time := 0; time < time_int; time++ {
			time_remaining := time_int - time
			round_distance := time_remaining * time
			if round_distance > distance_int {
				victory++
			}
		}
		victories = append(victories, victory)
	}

	fmt.Println(utils.Multiply(victories))
}

func part2() {
	defer utils.Timer("part1")()
	lines := utils.ReadInput("input.txt")

	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distance := strings.Fields(strings.Split(lines[1], ":")[1])

	// victories := []int{}
	distance_merge := ""
	time_merge := ""
	for i := 0; i < len(times); i++ {
		distance_merge += distance[i]
		time_merge += times[i]
	}
	distance_int := utils.IntegerOf(distance_merge)
	time_int := utils.IntegerOf(time_merge)

	time := 0
	for (time * (time_int - time)) < distance_int {
		time++
	}
	min_time := time

	time = time_int
	for (time * (time_int - time)) < distance_int {
		time--
	}
	max_time := time

	fmt.Println(max_time - min_time + 1)
}

func main() {
	part1()
	part2()
}
