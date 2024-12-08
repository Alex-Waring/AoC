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

	antinodes_chan := make(chan []utils.Position, len(locs))
	for _, f_loc := range locs {
		solve2(f_loc, antinodes_chan, board)
	}
	for i := 0; i < len(locs); i++ {
		antinode := <-antinodes_chan
		for _, pos := range antinode {
			antinodes[pos] = true
		}
	}

	fmt.Println(len(antinodes))
}

func solve2(f_loc []utils.Position, antinodes chan<- []utils.Position, board utils.Board) {
	return_antinodes := []utils.Position{}
	for i := 0; i < len(f_loc); i++ {
		for j := 0; j < len(f_loc); j++ {
			if i != j {
				row_diff := f_loc[i].Row - f_loc[j].Row
				col_diff := f_loc[i].Col - f_loc[j].Col
				// Set the locs as antinodes
				return_antinodes = append(return_antinodes, f_loc[i])
				return_antinodes = append(return_antinodes, f_loc[j])
				// generate antinodes until we hit a wall
				step_row := row_diff
				step_col := col_diff
			dir_1:
				for {
					antinode_1 := utils.Position{Row: f_loc[i].Row + row_diff, Col: f_loc[i].Col + col_diff}
					if _, ok := board[antinode_1]; ok {
						return_antinodes = append(return_antinodes, antinode_1)
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
						return_antinodes = append(return_antinodes, antinode_2)
					} else {
						break dir_2
					}
					row_diff += step_row
					col_diff += step_col
				}
			}
		}
	}
	antinodes <- return_antinodes
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
