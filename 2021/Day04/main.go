package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Grid struct {
	values  [5][5]int
	checked [5][5]bool
}

func part1(grids []Grid, calls []int) {
	defer utils.Timer("part1")()

	for _, call := range calls {
		for i := 0; i < len(grids); i++ {
			for row := 0; row < 5; row++ {
				for col := 0; col < 5; col++ {
					if grids[i].values[row][col] == call {
						grids[i].checked[row][col] = true
					}
				}
			}
			if checkGridPassed(grids[i]) {
				fmt.Println(scoreGrid(grids[i], call))
				return
			}
		}
	}
}

func part2(grids []Grid, calls []int) {
	defer utils.Timer("part2")()

	for _, call := range calls {
		newGrids := []Grid{}
		for i := 0; i < len(grids); i++ {
			for row := 0; row < 5; row++ {
				for col := 0; col < 5; col++ {
					if grids[i].values[row][col] == call {
						grids[i].checked[row][col] = true
					}
				}
			}
			if !checkGridPassed(grids[i]) {
				newGrids = append(newGrids, grids[i])
			}
		}
		if len(newGrids) == 0 {
			fmt.Println(scoreGrid(grids[0], call))
			return
		}
		grids = newGrids
	}
}

func checkGridPassed(grid Grid) bool {
	for row := 0; row < 5; row++ {
		rowPassed := true
		for col := 0; col < 5; col++ {
			rowPassed = rowPassed && grid.checked[row][col]
		}
		if rowPassed {
			return true
		}
	}
	for col := 0; col < 5; col++ {
		colPassed := true
		for row := 0; row < 5; row++ {
			colPassed = colPassed && grid.checked[row][col]
		}
		if colPassed {
			return true
		}
	}
	return false
}

func scoreGrid(grid Grid, call int) int {
	score := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !grid.checked[row][col] {
				score += grid.values[row][col]
			}
		}
	}
	return score * call
}

func main() {
	file := "input.txt"
	input := utils.ReadInput(file)

	calls := []int{}
	for _, score := range strings.Split(input[0], ",") {
		calls = append(calls, utils.IntegerOf(score))
	}

	raw_input, _ := os.ReadFile(file)
	grids := []Grid{}

	for i, grid := range strings.Split(string(raw_input), "\n\n") {
		if i == 0 {
			continue
		}
		newGrid := Grid{}
		for row, line := range strings.Split(grid, "\n") {
			for col, num := range strings.Fields(line) {
				newGrid.values[row][col] = utils.IntegerOf(num)
			}
		}
		grids = append(grids, newGrid)
	}

	part1(grids, calls)

	grids = []Grid{}

	for i, grid := range strings.Split(string(raw_input), "\n\n") {
		if i == 0 {
			continue
		}
		newGrid := Grid{}
		for row, line := range strings.Split(grid, "\n") {
			for col, num := range strings.Fields(line) {
				newGrid.values[row][col] = utils.IntegerOf(num)
			}
		}
		grids = append(grids, newGrid)
	}
	part2(grids, calls)
}
