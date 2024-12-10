package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()

	board := make(map[utils.Position]int)
	total := 0
	starts := []utils.Position{}

	dirs := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	for row, line := range input {
		for col, c := range line {
			board[utils.Position{Row: row, Col: col}] = int(c - '0')
			if int(c-'0') == 0 {
				starts = append(starts, utils.Position{Row: row, Col: col})
			}
		}
	}

	for _, start := range starts {
		q := utils.Queue[utils.Position]{}
		q.Push(start)
		hike_total := 0
		visited := make(map[utils.Position]bool)
		for !q.IsEmpty() {
			pos := q.Pop()
			height := board[pos]

			if height == 9 {
				hike_total++
				continue
			}

			for _, dir := range dirs {
				newPos := pos.Move(dir, 1)
				if _, exists := board[newPos]; !exists {
					continue
				}
				if visited[newPos] {
					continue
				}
				if board[newPos] == height+1 {
					visited[newPos] = true
					q.Push(newPos)
				}
			}
		}
		total += hike_total
	}

	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	board := make(map[utils.Position]int)
	total := 0
	starts := []utils.Position{}

	dirs := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	for row, line := range input {
		for col, c := range line {
			board[utils.Position{Row: row, Col: col}] = int(c - '0')
			if int(c-'0') == 0 {
				starts = append(starts, utils.Position{Row: row, Col: col})
			}
		}
	}

	for _, start := range starts {
		q := utils.Queue[utils.Position]{}
		q.Push(start)
		hike_total := 0
		for !q.IsEmpty() {
			pos := q.Pop()
			height := board[pos]

			if height == 9 {
				hike_total++
				continue
			}

			for _, dir := range dirs {
				newPos := pos.Move(dir, 1)
				if _, exists := board[newPos]; !exists {
					continue
				}
				if board[newPos] == height+1 {
					q.Push(newPos)
				}
			}
		}
		total += hike_total
	}

	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
