package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Rock [][9]int

func intInSlice(a int, list [9]int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func copyGrid(grid [][9]int) [][9]int {
	return_grid := [][9]int{}

	for _, line := range grid {
		return_grid = append(return_grid, line)
	}
	return return_grid
}

func stripEmptyLines(grid [][9]int) [][9]int {
	for i := len(grid) - 1; i >= 0; i-- {
		if grid[i] == [9]int{1, 0, 0, 0, 0, 0, 0, 0, 1} {
			grid = grid[:i]
		}
	}
	return grid
}

func printGrid(landscape [][9]int, rock_pos [][9]int) {
	for j := len(landscape) - 1; j >= 0; j-- {
		for k := 0; k < len(landscape[j]); k++ {
			if j == 0 && (k == 0 || k == 8) {
				fmt.Print("+")
			} else if j == 0 {
				fmt.Print("-")
			} else if k == 0 || k == 8 {
				fmt.Print("|")
			} else if rock_pos[j][k] == 1 {
				fmt.Print("@")
			} else if landscape[j][k] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveDown(landscape [][9]int, rock_pos [][9]int, lines int) ([][9]int, bool) {
	return_grid := make([][9]int, len(rock_pos))

	for index, line := range rock_pos {
		if intInSlice(1, line) {
			for i := index; i < index+lines; i++ {
				return_grid[i-1] = rock_pos[i]

				// Check for a collision, if there is return the original
				for j := 0; j < 9; j++ {
					coll := return_grid[i-1][j] + landscape[i-1][j]
					if coll > 1 {
						return rock_pos, false
					}
				}
			}
			return return_grid, true
		}
	}
	panic("Rock not found in rock_pos")
}

func moveRock(landscape [][9]int, rock_pos [][9]int, direction string) ([][9]int, bool) {
	return_grid := copyGrid(rock_pos)

	// Keep track if we've crossed a rock, if we have and the current line
	// has no rock in it, we can return early
	crossed_rock := false

	if direction == ">" {
		for index, line := range rock_pos {
			// Create a new line, that is rock pos shifted to the right
			new_line := [9]int{}
			for i := 1; i < 9; i++ {
				new_line[i] = line[i-1]
			}

			// Check for a collision, if there is return the original
			for i := 0; i < 9; i++ {
				coll := new_line[i] + landscape[index][i]
				if coll > 1 {
					return rock_pos, false
				}
			}

			// No collision, replace the line
			return_grid[index] = new_line

			// See if we've passed a rock here
			if intInSlice(1, new_line) {
				crossed_rock = true
			} else if crossed_rock {
				// As there isn't a rock here, and we have crossed one, return
				return return_grid, true
			}
		}
	} else {
		for index, line := range rock_pos {
			// Create a new line, that is rock pos shifted to the left
			new_line := [9]int{}
			for i := 0; i < 8; i++ {
				new_line[i] = line[i+1]
			}

			// Check for a collision, if there is return the original
			for i := 0; i < 9; i++ {
				coll := new_line[i] + landscape[index][i]
				if coll > 1 {
					return rock_pos, false
				}
			}

			// No collision, replace the line
			return_grid[index] = new_line

			// See if we've passed a rock here
			if intInSlice(1, new_line) {
				crossed_rock = true
			} else if crossed_rock {
				// As there isn't a rock here, and we have crossed one, return
				return return_grid, true
			}
		}
	}
	return return_grid, true
}

func part1(input string, total_rocks int) {
	defer utils.Timer("part1")()
	// Store two grids, one for the landscape and one the rock
	// The landscape has walls
	landscape := [][9]int{}
	landscape = append(landscape, [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1})

	jets := utils.LinkedList{}
	for _, char := range input {
		jets.Insert(string(char))
	}
	utils.ConvertSinglyToCircular(&jets)

	rocks := utils.LinkedList{}
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 1, 1, 1, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 0, 1, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 1, 1, 0, 0, 0}, [9]int{0, 0, 0, 0, 1, 0, 0, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 1, 1, 0, 0, 0}, [9]int{0, 0, 0, 0, 0, 1, 0, 0, 0}, [9]int{0, 0, 0, 0, 0, 1, 0, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 1, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 1, 0, 0, 0, 0}})
	utils.ConvertSinglyToCircular(&rocks)

	current_rock := utils.GetFirst(&rocks)
	current_jet := utils.GetFirst(&jets)

	for i := 0; i < total_rocks; i++ {
		rock := current_rock.GetInfo().(Rock)

		// First of all place three blank walled lines
		for i := 0; i < 3; i++ {
			landscape = append(landscape, [9]int{1, 0, 0, 0, 0, 0, 0, 0, 1})
		}

		// Create a rock grid that is completely empty but the same length as the current grid, then stick
		// the rock at the top of it and walls on the landscape
		rock_pos := make([][9]int, len(landscape))
		for _, row := range rock {
			landscape = append(landscape, [9]int{1, 0, 0, 0, 0, 0, 0, 0, 1})
			rock_pos = append(rock_pos, row)
		}

		moving_down := true
		for moving_down {
			jet := current_jet.GetInfo().(string)
			// Move the rock according to the jet
			rock_pos, _ = moveRock(landscape, rock_pos, jet)

			// Try and move the rock down
			rock_pos, moving_down = moveDown(landscape, rock_pos, len(rock))
			current_jet = utils.GetNext(current_jet)
		}

		for i := 0; i < len(landscape); i++ {
			for j := 0; j < len(landscape[i]); j++ {
				landscape[i][j] += rock_pos[i][j]
			}
		}
		landscape = stripEmptyLines(landscape)

		current_rock = utils.GetNext(current_rock)
	}

	fmt.Println(len(landscape) - 1)
}

func part2(input string, total_rocks int) {
	defer utils.Timer("part2")()
	// Setup as per part 1
	landscape := [][9]int{}
	landscape = append(landscape, [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1})

	jets := utils.LinkedList{}
	for _, char := range input {
		jets.Insert(string(char))
	}
	utils.ConvertSinglyToCircular(&jets)

	rocks := utils.LinkedList{}
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 1, 1, 1, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 0, 1, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 1, 1, 0, 0, 0}, [9]int{0, 0, 0, 0, 1, 0, 0, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 1, 1, 0, 0, 0}, [9]int{0, 0, 0, 0, 0, 1, 0, 0, 0}, [9]int{0, 0, 0, 0, 0, 1, 0, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 0, 0, 0, 0, 0}})
	rocks.Insert(Rock{[9]int{0, 0, 0, 1, 1, 0, 0, 0, 0}, [9]int{0, 0, 0, 1, 1, 0, 0, 0, 0}})
	utils.ConvertSinglyToCircular(&rocks)

	current_rock := utils.GetFirst(&rocks)
	current_jet := utils.GetFirst(&jets)

	type entry struct {
		top_10_layers [10][9]int
		next_jet      *utils.Node
		rock          *utils.Node
	}
	cache := map[entry]int{}

	final_i := 0
	loop_length := 0

	for i := 0; i < total_rocks; i++ {
		rock := current_rock.GetInfo().(Rock)

		// Check the cache, only if we have enough layers
		if len(landscape) > 9 {
			e := entry{
				top_10_layers: [10][9]int(landscape[len(landscape)-11 : len(landscape)-1]),
				next_jet:      current_jet,
				rock:          current_rock,
			}
			if rocks, exists := cache[e]; exists {
				final_i = i
				loop_length = rocks
				break
			} else {
				cache[e] = i
			}
		}

		// First of all place three blank walled lines
		for i := 0; i < 3; i++ {
			landscape = append(landscape, [9]int{1, 0, 0, 0, 0, 0, 0, 0, 1})
		}

		// Create a rock grid that is completely empty but the same length as the current grid, then stick
		// the rock at the top of it and walls on the landscape
		rock_pos := make([][9]int, len(landscape))
		for _, row := range rock {
			landscape = append(landscape, [9]int{1, 0, 0, 0, 0, 0, 0, 0, 1})
			rock_pos = append(rock_pos, row)
		}

		moving_down := true
		for moving_down {
			jet := current_jet.GetInfo().(string)
			// Move the rock according to the jet
			rock_pos, _ = moveRock(landscape, rock_pos, jet)

			// Try and move the rock down
			rock_pos, moving_down = moveDown(landscape, rock_pos, len(rock))
			current_jet = utils.GetNext(current_jet)
		}

		for i := 0; i < len(landscape); i++ {
			for j := 0; j < len(landscape[i]); j++ {
				landscape[i][j] += rock_pos[i][j]
			}
		}
		landscape = stripEmptyLines(landscape)

		current_rock = utils.GetNext(current_rock)
	}
	fmt.Println(final_i)
	fmt.Println(loop_length)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input[0], 2022)
	part2(input[0], 1000000000000)
}
