package main

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string, bytes int) bool {
	// defer utils.Timer("part1")()

	board := utils.NewBoard()

	for row := 0; row < 71; row++ {
		for col := 0; col < 71; col++ {
			board[utils.NewPosition(row, col)] = "."
		}
	}

	for i := 0; i < bytes; i++ {
		mem := strings.Split(input[i], ",")
		board[utils.NewPosition(utils.IntegerOf(mem[1]), utils.IntegerOf(mem[0]))] = "#"
	}

	pq := make(utils.PriorityQueue, 1)
	pq[0] = utils.NewItem(utils.NewPosition(0, 0), 0, 0)
	heap.Init(&pq)

	cost_so_far := map[utils.Position]int{}
	cost_so_far[utils.NewPosition(0, 0)] = 0

	target := utils.NewPosition(70, 70)
	dirs := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*utils.Item)
		pos := item.GetValue().(utils.Position)
		cost := item.GetPriority()

		if pos == target {
			// fmt.Println(cost)
			return true
		}

		for _, dir := range dirs {
			newPos := pos.Move(dir, 1)
			newCost := cost + 1
			if tile, ok := board[newPos]; ok && tile != "#" {
				if cost, ok := cost_so_far[newPos]; !ok || cost > newCost {
					cost_so_far[newPos] = newCost
					newItem := utils.NewItem(newPos, newCost, 0)
					heap.Push(&pq, newItem)
					pq.Update(newItem, newCost)
				}
			}
		}
	}
	return false
}

func part2(input []string) {
	defer utils.Timer("part2")()

	max := 3450

	start_bytes := 3450
	previous_check := 1024

	for {
		can_escape := part1(input, start_bytes)

		if can_escape {
			new_start_bytes := (start_bytes + max) / 2
			previous_check = start_bytes
			start_bytes = new_start_bytes
		} else {
			if start_bytes-previous_check == 1 {
				fmt.Println(input[previous_check])
				return
			} else {
				start_bytes = (start_bytes + previous_check) / 2
			}
		}
	}
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input, 1024)
	part2(input)
}
