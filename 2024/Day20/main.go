package main

import (
	"container/heap"
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Path struct {
	Position utils.Position
	Visited  []utils.Position
}

func part1(input []string) {
	defer utils.Timer("part1")()
	solve(input, 2)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	solve(input, 20)
}

func solve(input []string, cheat int) {
	board := utils.NewBoard()
	var start utils.Position

	for row, line := range input {
		for col, char := range line {
			board[utils.NewPosition(row, col)] = string(char)
			if char == 'S' {
				start = utils.NewPosition(row, col)
			}
		}
	}

	pq := make(utils.PriorityQueue, 1)
	pq[0] = utils.NewItem(Path{
		Position: start,
	}, 0, 0)
	heap.Init(&pq)

	cost_so_far := map[utils.Position]int{}
	cost_so_far[start] = 0

	dirs := []utils.Direction{utils.Up, utils.Down, utils.Left, utils.Right}

	// Find the shortest path to the exit
	var shortestPath Path

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*utils.Item)
		pos := item.GetValue().(Path).Position
		cost := item.GetPriority()
		visted := item.GetValue().(Path).Visited

		if board[pos] == "E" {
			shortestPath = item.GetValue().(Path)
			shortestPath.Visited = append(shortestPath.Visited, pos)
			break
		}

		for _, dir := range dirs {
			newPos := pos.Move(dir, 1)
			newCost := cost + 1
			if tile, ok := board[newPos]; ok && tile != "#" {
				if cost, ok := cost_so_far[newPos]; !ok || cost > newCost {
					cost_so_far[newPos] = newCost
					newItem := utils.NewItem(Path{
						Position: newPos,
						Visited:  append(visted, pos),
					}, newCost, 0)
					heap.Push(&pq, newItem)
					pq.Update(newItem, newCost)
				}
			}
		}
	}

	distance_map := map[utils.Position]int{}

	for distance, pos := range shortestPath.Visited {
		distance_map[pos] = distance
	}

	cheats := map[int]int{}

	// For every pair, if manhatten distance is less than two then find the cheat
	for pos1, distance1 := range distance_map {
		for pos2, distance2 := range distance_map {
			if pos1 == pos2 {
				continue
			}
			if pos1.Manhattan(pos2) == 1 {
				// This is on the path, we can skip
				continue
			}
			if pos1.Manhattan(pos2) <= cheat {
				cheat := utils.Abs(distance1-distance2) - pos1.Manhattan(pos2)
				// fmt.Printf("Pos1: %v, Pos2: %v, Cheat: %d\n", pos1, pos2, cheat)
				if cheat == 0 {
					continue
				}
				cheats[cheat]++
			}
		}
	}

	answer := 0

	for cheat, count := range cheats {
		if cheat >= 100 {
			answer += count / 2
		}
	}
	fmt.Println(answer)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
