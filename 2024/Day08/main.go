package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	board := utils.NewBoard()
	locs := map[string][]utils.Position{}
	antinodes := map[utils.Position]bool{}

	for row, line := range input {
		for col, char := range line {
			board[utils.Position{Row: row, Col: col}] = string(char)
			if string(char) != "." {
				locs[string(char)] = append(locs[string(char)], utils.Position{Row: row, Col: col})
			}
		}
	}

	for _, f_loc := range locs {
		for i := 0; i < len(f_loc); i++ {
			for j := 0; j < len(f_loc); j++ {
				if i != j {
					row_diff := f_loc[i].Row - f_loc[j].Row
					col_diff := f_loc[i].Col - f_loc[j].Col
					antinode_1 := utils.Position{Row: f_loc[i].Row + row_diff, Col: f_loc[i].Col + col_diff}
					antinode_2 := utils.Position{Row: f_loc[j].Row - row_diff, Col: f_loc[j].Col - col_diff}
					if _, ok := board[antinode_1]; ok {
						antinodes[antinode_1] = true
					}
					if _, ok := board[antinode_2]; ok {
						antinodes[antinode_2] = true
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))

}

func part2(input []string) {
	defer utils.Timer("part2")()

	board := utils.NewBoard()
	locs := map[string][]utils.Position{}
	antinodes := map[utils.Position]bool{}

	for row, line := range input {
		for col, char := range line {
			board[utils.Position{Row: row, Col: col}] = string(char)
			if string(char) != "." {
				locs[string(char)] = append(locs[string(char)], utils.Position{Row: row, Col: col})
			}
		}
	}

	for _, f_loc := range locs {
		for i := 0; i < len(f_loc); i++ {
			for j := 0; j < len(f_loc); j++ {
				if i != j {
					row_diff := f_loc[i].Row - f_loc[j].Row
					col_diff := f_loc[i].Col - f_loc[j].Col
					// Set the locs as antinodes
					antinodes[f_loc[i]] = true
					antinodes[f_loc[j]] = true
					// generate antinodes until we hit a wall
					f_antinodes := []utils.Position{}
					step_row := row_diff
					step_col := col_diff
				dir_1:
					for {
						antinode_1 := utils.Position{Row: f_loc[i].Row + row_diff, Col: f_loc[i].Col + col_diff}
						if _, ok := board[antinode_1]; ok {
							f_antinodes = append(f_antinodes, antinode_1)
						} else {
							break dir_1
						}
						row_diff += step_row
						col_diff += step_col
					}
					row_diff = f_loc[i].Row - f_loc[j].Row
					col_diff = f_loc[i].Col - f_loc[j].Col
				dir_2:
					for {
						antinode_2 := utils.Position{Row: f_loc[j].Row - row_diff, Col: f_loc[j].Col - col_diff}
						if _, ok := board[antinode_2]; ok {
							f_antinodes = append(f_antinodes, antinode_2)
						} else {
							break dir_2
						}
						row_diff += step_row
						col_diff += step_col
					}
					for _, antinode := range f_antinodes {
						antinodes[antinode] = true
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
