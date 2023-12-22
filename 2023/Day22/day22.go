package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Location struct {
	x int
	y int
	z int
}

type Brick struct {
	index int
	front Location
	back  Location
}

// Return if there is a block above
func block_below(grid map[Location]int, bricks map[int]Brick, brick_index int) (bool, int) {

	brick := bricks[brick_index]

	// if z is 1, we are already on the bottom
	if min(brick.front.z, brick.back.z) == 1 {
		return true, 0
	}

	for x := brick.front.x; x <= brick.back.x; x++ {
		for y := brick.front.y; y <= brick.back.y; y++ {

			z := min(brick.front.z, brick.back.z)
			below := Location{x: x, y: y, z: z - 1}

			if block, exists := grid[below]; exists {
				return true, block
			}

		}
	}
	return false, 0
}

// Return a set of all of the blocks below a block
func bricks_below(grid map[Location]int, bricks map[int]Brick, brick_index int) map[int]bool {
	return_bricks := make(map[int]bool)
	brick := bricks[brick_index]

	for x := brick.front.x; x <= brick.back.x; x++ {
		for y := brick.front.y; y <= brick.back.y; y++ {
			z := min(brick.front.z, brick.back.z)

			below := Location{x: x, z: z - 1, y: y}

			if below_index, exists := grid[below]; exists {
				return_bricks[below_index] = true
			}
		}
	}
	return return_bricks
}

// return a set of all of the blocks above a block
func bricks_above(grid map[Location]int, bricks map[int]Brick, brick_index int) map[int]bool {
	return_bricks := make(map[int]bool)
	brick := bricks[brick_index]

	for x := brick.front.x; x <= brick.back.x; x++ {
		for y := brick.front.y; y <= brick.back.y; y++ {
			z := max(brick.front.z, brick.back.z)

			above := Location{x: x, z: z + 1, y: y}

			if above_index, exists := grid[above]; exists {
				return_bricks[above_index] = true
			}
		}
	}
	return return_bricks
}

func copy_grid(grid map[Location]int) map[Location]int {
	return_grid := map[Location]int{}
	for loc, index := range grid {
		return_grid[loc] = index
	}
	return return_grid
}

func copy_bricks(bricks map[int]Brick) map[int]Brick {
	return_bricks := map[int]Brick{}
	for loc, index := range bricks {
		return_bricks[loc] = index
	}
	return return_bricks
}

func main() {
	input := utils.ReadInput("input.txt")
	grid := map[Location]int{}
	bricks := map[int]Brick{}

	max_x := 0
	max_y := 0
	max_z := 0

	for index, line := range input {
		front := strings.Split(line, "~")[0]
		back := strings.Split(line, "~")[1]

		front_x, front_y, front_z := strings.Split(front, ",")[0], strings.Split(front, ",")[1], strings.Split(front, ",")[2]
		back_x, back_y, back_z := strings.Split(back, ",")[0], strings.Split(back, ",")[1], strings.Split(back, ",")[2]

		brick := Brick{
			index: index,
			front: Location{x: utils.IntegerOf(front_x), y: utils.IntegerOf(front_y), z: utils.IntegerOf(front_z)},
			back:  Location{x: utils.IntegerOf(back_x), y: utils.IntegerOf(back_y), z: utils.IntegerOf(back_z)},
		}

		for x := brick.front.x; x <= brick.back.x; x++ {
			for y := brick.front.y; y <= brick.back.y; y++ {
				for z := brick.front.z; z <= brick.back.z; z++ {
					loc := Location{x: x, y: y, z: z}
					grid[loc] = index
					max_x = max(max_x, x)
					max_y = max(max_y, y)
					max_z = max(max_z, z)
				}
			}
		}

		bricks[index] = brick
	}

	// Loop through the bricks, keeping track if any have falled, and settle them to the bottom
	rotation := false
	for !rotation {
		rotation = true
		for index, brick := range bricks {
			is_block_below, _ := block_below(grid, bricks, index)

			if !is_block_below {
				rotation = false
				for x := brick.front.x; x <= brick.back.x; x++ {
					for y := brick.front.y; y <= brick.back.y; y++ {
						for z := brick.front.z; z <= brick.front.z; z++ {
							loc := Location{x: x, y: y, z: z}
							delete(grid, loc)
							loc.z = loc.z - 1
							grid[loc] = index

							new_brick := Brick{
								index: index,
								front: Location{x: brick.front.x, y: brick.front.y, z: brick.front.z - 1},
								back:  Location{x: brick.back.x, y: brick.back.y, z: brick.back.z - 1},
							}
							bricks[index] = new_brick
						}
					}
				}
			}
		}
	}

	// For every brick, if there is a brick above it check if that brick has multiple bricks below it
	// If every brick above a brick has multiple supporting it, it can be disintegrated
	disintegrated := make(map[int]bool)
	for index, _ := range bricks {
		bricks_above := bricks_above(grid, bricks, index)
		can_disintegrate := true

		for brick := range bricks_above {
			bricks_below := bricks_below(grid, bricks, brick)
			if len(bricks_below) <= 1 {
				can_disintegrate = false
			}
		}
		if can_disintegrate {
			disintegrated[index] = true
		}
	}

	fmt.Println(len(disintegrated))

	// For part 2, brute force. For every brick dissintegrate it and loop through every brick, counting
	// the number of times we fall
	total := 0
	for index, brick := range bricks {
		grid_copy := copy_grid(grid)
		bricks_copy := copy_bricks(bricks)
		delete(bricks_copy, index)

		// Disintegrate the brick
		for x := brick.front.x; x <= brick.back.x; x++ {
			for y := brick.front.y; y <= brick.back.y; y++ {
				for z := brick.front.z; z <= brick.front.z; z++ {
					loc := Location{x: x, y: y, z: z}
					delete(grid_copy, loc)
				}
			}
		}
		bricks_fallen := 0

		// See what's fallen
		rotation := false
		for !rotation {
			rotation = true
			for index, brick := range bricks_copy {
				is_block_below, _ := block_below(grid_copy, bricks_copy, index)

				if !is_block_below {
					rotation = false
					for x := brick.front.x; x <= brick.back.x; x++ {
						for y := brick.front.y; y <= brick.back.y; y++ {
							for z := brick.front.z; z <= brick.front.z; z++ {
								loc := Location{x: x, y: y, z: z}
								delete(grid_copy, loc)
								loc.z = loc.z - 1
								grid_copy[loc] = index

								new_brick := Brick{
									index: index,
									front: Location{x: brick.front.x, y: brick.front.y, z: brick.front.z - 1},
									back:  Location{x: brick.back.x, y: brick.back.y, z: brick.back.z - 1},
								}
								bricks_copy[index] = new_brick
								bricks_fallen++
							}
						}
					}
				}
			}
		}
		total += bricks_fallen
	}
	fmt.Println(total)
}
