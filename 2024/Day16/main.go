package main

import (
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

	queue := utils.Queue[Path]{}
	queue.Push(Path{loc: start, steps: 0, turns: 0})
	scores := []int{}
	seen := map[utils.Location]int{}

	for !queue.IsEmpty() {
		path := queue.Pop()
		loc := path.loc
		steps := path.steps
		turns := path.turns

		// Check if finished
		if board[loc.Pos] == "E" {
			scores = append(scores, steps+1000*turns)
			continue
		}

		// Try to move in the directin we're facing
		newLoc := utils.Location{Pos: loc.Pos.Move(loc.Dir, 1), Dir: loc.Dir}
		new_cost := steps + 1 + 1000*turns
		if board[newLoc.Pos] != "#" {
			if cost, ok := seen[newLoc]; !ok || cost > new_cost {
				queue.Push(Path{loc: newLoc, steps: steps + 1, turns: turns})
				seen[newLoc] = new_cost
			}
		}

		// Try to turn left
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Left)}
		new_cost = steps + 1000*(turns+1)
		if cost, ok := seen[newLoc]; !ok || cost > new_cost {
			queue.Push(Path{loc: newLoc, steps: steps, turns: turns + 1})
			seen[newLoc] = new_cost
		}

		// Try to turn right
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Right)}
		new_cost = steps + 1000*(turns+1)
		if cost, ok := seen[newLoc]; !ok || cost > new_cost {
			queue.Push(Path{loc: newLoc, steps: steps, turns: turns + 1})
			seen[newLoc] = new_cost
		}
	}

	fmt.Println(utils.FindMin(scores))
	return utils.FindMin(scores)
}

type StoredPath struct {
	loc   utils.Location
	steps int
	turns int
	path  map[utils.Position]bool
}

func part2(input []string, min int) {
	defer utils.Timer("part2")()

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

	queue := utils.Queue[StoredPath]{}
	queue.Push(StoredPath{loc: start, steps: 0, turns: 0, path: make(map[utils.Position]bool)})
	scores := []int{}
	seen := map[utils.Location]int{}

	type path_score struct {
		score int
		path  []utils.Position
	}

	finishing_paths := []path_score{}

	for !queue.IsEmpty() {
		path := queue.Pop()
		loc := path.loc
		steps := path.steps
		turns := path.turns

		// Slight cheat, use the answer from part 1 to prune the search space
		if steps+1000*turns > min {
			continue
		}

		// Check if finished
		if board[loc.Pos] == "E" {
			scores = append(scores, steps+1000*turns)
			new_finishing_path := []utils.Position{}
			for pos := range path.path {
				new_finishing_path = append(new_finishing_path, pos)
			}
			finishing_paths = append(finishing_paths, path_score{score: steps + 1000*turns, path: new_finishing_path})
			continue
		}

		// Try to move in the directin we're facing
		newLoc := utils.Location{Pos: loc.Pos.Move(loc.Dir, 1), Dir: loc.Dir}
		new_cost := steps + 1 + 1000*turns
		if board[newLoc.Pos] != "#" {
			if cost, ok := seen[newLoc]; !ok || cost >= new_cost {
				new_path := duplicatePath(path.path)
				new_path[newLoc.Pos] = true
				queue.Push(StoredPath{loc: newLoc, steps: steps + 1, turns: turns, path: new_path})
				seen[newLoc] = new_cost
			}
		}

		// Try to turn left
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Left)}
		new_cost = steps + 1000*(turns+1)
		if cost, ok := seen[newLoc]; !ok || cost >= new_cost {
			queue.Push(StoredPath{loc: newLoc, steps: steps, turns: turns + 1, path: path.path})
			seen[newLoc] = new_cost
		}

		// Try to turn right
		newLoc = utils.Location{Pos: loc.Pos, Dir: loc.Dir.Turn(utils.Right)}
		new_cost = steps + 1000*(turns+1)
		if cost, ok := seen[newLoc]; !ok || cost >= new_cost {
			queue.Push(StoredPath{loc: newLoc, steps: steps, turns: turns + 1, path: path.path})
			seen[newLoc] = new_cost
		}
	}

	on_finish := map[utils.Position]bool{}

	for _, path := range finishing_paths {
		if path.score == min {
			for _, pos := range path.path {
				on_finish[pos] = true
			}
		}
	}
	fmt.Println(len(on_finish) + 1)

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
