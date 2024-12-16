package main

import (
	"container/heap"
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Path struct {
	loc   utils.Location
	steps int
	turns int
}

func part1(input []string) int {
	defer utils.Timer("part1")()

	board := utils.NewBoard()
	start := utils.NewLocation(0, 0, utils.Right)

	for row, line := range input {
		for col, char := range line {
			board[utils.NewPosition(row, col)] = string(char)
			if char == 'S' {
				start = utils.NewLocation(row, col, utils.Right)
			}
		}
	}

	pq := make(utils.PriorityQueue, 1)
	pq[0] = utils.NewItem(start, 0, 0)
	heap.Init(&pq)

	cost_so_far := map[utils.Location]int{}
	cost_so_far[start] = 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*utils.Item)
		loc := item.GetValue().(utils.Location)
		cost := item.GetPriority()

		// Check if finished
		if board[loc.Pos] == "E" {
			fmt.Println(cost)
			return cost
		}

		// Try to move in the directin we're facing
		newLoc := utils.Location{Pos: loc.Pos.Move(loc.Dir, 1), Dir: loc.Dir}
		new_cost := cost + 1
		if board[newLoc.Pos] != "#" {
			if cost, ok := cost_so_far[newLoc]; !ok || cost > new_cost {
				cost_so_far[newLoc] = new_cost
				new_item := utils.NewItem(newLoc, new_cost, 0)
				heap.Push(&pq, new_item)
				pq.Update(new_item, new_cost)
			}
		}

		// Try to turn left
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Left)}
		new_cost = cost + 1000
		if cost, ok := cost_so_far[newLoc]; !ok || cost > new_cost {
			cost_so_far[newLoc] = new_cost
			new_item := utils.NewItem(newLoc, new_cost, 0)
			heap.Push(&pq, new_item)
			pq.Update(new_item, new_cost)
		}

		// Try to turn right
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Right)}
		new_cost = cost + 1000
		if cost, ok := cost_so_far[newLoc]; !ok || cost > new_cost {
			cost_so_far[newLoc] = new_cost
			new_item := utils.NewItem(newLoc, new_cost, 0)
			heap.Push(&pq, new_item)
			pq.Update(new_item, new_cost)
		}
	}

	panic("No solution found")
}

type StoredPath struct {
	loc  utils.Location
	path map[utils.Position]bool
	cost int
}

func part2(input []string, min int) {
	defer utils.Timer("part2")()

	board := utils.NewBoard()
	start := utils.NewLocation(0, 0, utils.Right)
	end := utils.NewPosition(0, 0)

	for row, line := range input {
		for col, char := range line {
			board[utils.NewPosition(row, col)] = string(char)
			if char == 'S' {
				start = utils.NewLocation(row, col, utils.Right)
			}
			if char == 'E' {
				end = utils.NewPosition(row, col)
			}
		}
	}

	pq := make(utils.PriorityQueue, 1)
	pq[0] = utils.NewItem(StoredPath{
		loc:  start,
		path: map[utils.Position]bool{start.Pos: true},
		cost: 0,
	}, end.Manhattan(start.Pos), 0)
	heap.Init(&pq)

	cost_so_far := map[utils.Location]int{}
	cost_so_far[start] = 0

	finishing_paths := []StoredPath{}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*utils.Item)
		item_value := item.GetValue().(StoredPath)
		loc := item_value.loc
		cost := item_value.cost

		// Check if finished
		if board[item_value.loc.Pos] == "E" {
			finishing_paths = append(finishing_paths, item_value)
			continue
		}

		// If we are over the minimum, we can stop
		if cost > min {
			continue
		}

		// Try to move in the directin we're facing
		newLoc := utils.Location{Pos: loc.Pos.Move(loc.Dir, 1), Dir: loc.Dir}
		new_cost := cost + 1
		if board[newLoc.Pos] != "#" {
			if cost, ok := cost_so_far[newLoc]; !ok || cost >= new_cost {
				cost_so_far[newLoc] = new_cost
				new_path := duplicatePath(item_value.path)
				new_path[newLoc.Pos] = true
				new_item := utils.NewItem(StoredPath{
					loc:  newLoc,
					path: new_path,
					cost: new_cost,
				}, new_cost+end.Manhattan(newLoc.Pos), 0)
				heap.Push(&pq, new_item)
				pq.Update(new_item, new_cost)
			}
		}

		// Try to turn left
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Left)}
		new_cost = cost + 1000
		if cost, ok := cost_so_far[newLoc]; !ok || cost >= new_cost {
			cost_so_far[newLoc] = new_cost
			new_path := duplicatePath(item_value.path)
			new_path[newLoc.Pos] = true
			new_item := utils.NewItem(StoredPath{
				loc:  newLoc,
				path: new_path,
				cost: new_cost,
			}, new_cost+end.Manhattan(newLoc.Pos), 0)
			heap.Push(&pq, new_item)
			pq.Update(new_item, new_cost)
		}

		// Try to turn right
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Right)}
		new_cost = cost + 1000
		if cost, ok := cost_so_far[newLoc]; !ok || cost >= new_cost {
			cost_so_far[newLoc] = new_cost
			new_path := duplicatePath(item_value.path)
			new_path[newLoc.Pos] = true
			new_item := utils.NewItem(StoredPath{
				loc:  newLoc,
				path: new_path,
				cost: new_cost,
			}, new_cost+end.Manhattan(newLoc.Pos), 0)
			heap.Push(&pq, new_item)
			pq.Update(new_item, new_cost)
		}
	}

	on_finish := map[utils.Position]bool{}

	for _, path := range finishing_paths {
		for pos := range path.path {
			on_finish[pos] = true
		}
	}
	fmt.Println(len(on_finish))

}

func duplicatePath(path map[utils.Position]bool) map[utils.Position]bool {
	new_path := make(map[utils.Position]bool)
	for pos := range path {
		new_path[pos] = true
	}
	return new_path
}

func main() {
	input := utils.ReadInput("input.txt")
	min := part1(input)
	part2(input, min)
}
