package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Position struct {
	x int
	y int
	z int
}

func sumCovered(cubes map[Position]int) int {
	covered := 0
	for _, c := range cubes {
		covered += c
	}
	return covered
}

func part1(cubes map[Position]int) {
	defer utils.Timer("part1")()

	for pos := range cubes {
		covered := 0

		above := Position{pos.x, pos.y, pos.z + 1}
		if _, exists := cubes[above]; exists {
			covered++
		}

		below := Position{pos.x, pos.y, pos.z - 1}
		if _, exists := cubes[below]; exists {
			covered++
		}

		north := Position{pos.x + 1, pos.y, pos.z}
		if _, exists := cubes[north]; exists {
			covered++
		}

		south := Position{pos.x - 1, pos.y, pos.z}
		if _, exists := cubes[south]; exists {
			covered++
		}

		east := Position{pos.x, pos.y + 1, pos.z}
		if _, exists := cubes[east]; exists {
			covered++
		}

		west := Position{pos.x, pos.y - 1, pos.z}
		if _, exists := cubes[west]; exists {
			covered++
		}
		cubes[pos] = covered
	}

	surface_area := len(cubes)*6 - sumCovered(cubes)
	fmt.Println(surface_area)
}

func part2(cubes map[Position]int) {
	defer utils.Timer("part2")()
	// Looking at the input, everything is happening in a 20*20*20 area
	// Fill up the area with steam, staring at -1,-1,-1, keep track of the number of
	// times the steam is blocked
	// Some room has been give to go round the bubble
	blocked := 0

	q := utils.Queue[Position]{}
	q.Push(Position{x: -1, y: -1, z: -1})

	seen := map[Position]bool{}

	for !q.IsEmpty() {
		steam := q.Pop()

		if _, exists := seen[steam]; exists {
			continue
		} else {
			seen[steam] = true
		}

		// move up
		above := Position{x: steam.x, y: steam.y, z: steam.z + 1}
		if _, exists := cubes[above]; exists {
			blocked++
		} else if above.z < 22 {
			q.Push(above)
		}

		below := Position{x: steam.x, y: steam.y, z: steam.z - 1}
		if _, exists := cubes[below]; exists {
			blocked++
		} else if below.z > -1 {
			q.Push(below)
		}

		north := Position{x: steam.x + 1, y: steam.y, z: steam.z}
		if _, exists := cubes[north]; exists {
			blocked++
		} else if north.x < 22 {
			q.Push(north)
		}

		south := Position{x: steam.x - 1, y: steam.y, z: steam.z}
		if _, exists := cubes[south]; exists {
			blocked++
		} else if south.x > -1 {
			q.Push(south)
		}

		east := Position{x: steam.x, y: steam.y + 1, z: steam.z}
		if _, exists := cubes[east]; exists {
			blocked++
		} else if east.y < 22 {
			q.Push(east)
		}

		west := Position{x: steam.x, y: steam.y - 1, z: steam.z}
		if _, exists := cubes[west]; exists {
			blocked++
		} else if west.y > -1 {
			q.Push(west)
		}
	}
	fmt.Println(blocked)
}

func main() {
	input := utils.ReadInput("input.txt")
	cubes := map[Position]int{}

	for _, line := range input {
		coords := strings.Split(line, ",")
		new_pos := Position{
			x: utils.IntegerOf(coords[0]),
			y: utils.IntegerOf(coords[1]),
			z: utils.IntegerOf(coords[2]),
		}
		cubes[new_pos] = 0
	}
	part1(cubes)
	part2(cubes)
}
