package main

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type State struct {
	robot1_pos      utils.Position
	last_move       string
	numbers_entered string
}

type HState struct {
	robot1_pos      utils.Position
	last_move       string
	numbers_entered string
	instructions    string
}

type Cache struct {
	robot1_pos      utils.Position
	last_move       string
	numbers_entered string
}

func solve(input []string, pads int) {
	defer utils.Timer("part1")()

	keypad := utils.NewBoard()
	keypad[utils.Position{Row: 0, Col: 0}] = "7"
	keypad[utils.Position{Row: 0, Col: 1}] = "8"
	keypad[utils.Position{Row: 0, Col: 2}] = "9"
	keypad[utils.Position{Row: 1, Col: 0}] = "4"
	keypad[utils.Position{Row: 1, Col: 1}] = "5"
	keypad[utils.Position{Row: 1, Col: 2}] = "6"
	keypad[utils.Position{Row: 2, Col: 0}] = "1"
	keypad[utils.Position{Row: 2, Col: 1}] = "2"
	keypad[utils.Position{Row: 2, Col: 2}] = "3"
	keypad[utils.Position{Row: 3, Col: 0}] = "X"
	keypad[utils.Position{Row: 3, Col: 1}] = "0"
	keypad[utils.Position{Row: 3, Col: 2}] = "A"

	keypad_lookup := map[string]utils.Position{}
	for pos, key := range keypad {
		keypad_lookup[key] = pos
	}

	controller := utils.NewBoard()
	controller[utils.Position{Row: 0, Col: 0}] = "X"
	controller[utils.Position{Row: 0, Col: 1}] = "^"
	controller[utils.Position{Row: 0, Col: 2}] = "A"
	controller[utils.Position{Row: 1, Col: 0}] = "<"
	controller[utils.Position{Row: 1, Col: 1}] = "v"
	controller[utils.Position{Row: 1, Col: 2}] = ">"

	numeric_map := map[int]int{
		0: 319,
		1: 85,
		2: 143,
		3: 286,
		4: 789,
	}
	total := 0

	controller_lookup := map[string]utils.Position{}
	for pos, key := range controller {
		controller_lookup[key] = pos
	}

	for i, line := range input {
		pq := make(utils.PriorityQueue, 1)
		pq[0] = utils.NewItem(State{
			robot1_pos:      keypad_lookup["A"],
			last_move:       "A",
			numbers_entered: "",
		}, 0, 0)
		heap.Init(&pq)

		// The cache doesn't care about the instructions
		seen := map[Cache]bool{}
		var finalLength int

		keys := []string{"^", "v", "<", ">", "A"}

		for pq.Len() > 0 {
			item := heap.Pop(&pq).(*utils.Item)
			state := item.GetValue().(State)
			robot1_pos := state.robot1_pos
			last_move := state.last_move
			numbers_entered := state.numbers_entered

			// Check if we've won
			if state.numbers_entered == line {
				finalLength = item.GetPriority()
				break
			}

			// Check if we're off track
			if !strings.HasPrefix(line, numbers_entered) {
				continue
			}

			// Check if we're over the gap
			if num, ok := keypad[robot1_pos]; !ok || num == "X" {
				continue
			}

			// Check if we've seen this state before
			cache := Cache{
				robot1_pos:      robot1_pos,
				last_move:       last_move,
				numbers_entered: numbers_entered,
			}
			if _, ok := seen[cache]; ok {
				continue
			}
			seen[cache] = true

			// Try and press the last controller
			for _, key := range keys {
				new_robot1_pos := robot1_pos
				new_out := numbers_entered
				new_robot1_pos, output := doKeypad(new_robot1_pos, key, keypad)
				if output != "" {
					new_out += output
				}
				cost_move := heuristic(key, last_move, pads, controller_lookup, controller) + item.GetPriority()
				newItem := utils.NewItem(State{
					robot1_pos:      new_robot1_pos,
					last_move:       key,
					numbers_entered: new_out,
				}, cost_move, 0)
				heap.Push(&pq, newItem)
				pq.Update(newItem, cost_move)
			}
		}
		total += numeric_map[i] * finalLength
	}
	fmt.Println(total)
}

var heuristicCache = map[string]int{}

// Layers is the number of robots we have to control
func heuristic(move string, prev_move string, layers int, keypad_lookup map[string]utils.Position, keypad utils.Board) int {
	// Quick caching
	key := fmt.Sprintf("%s%s%d", move, prev_move, layers)
	if val, ok := heuristicCache[key]; ok {
		return val
	}

	// We've moved the robot, so now return
	if layers == 0 {
		return 1
	}
	start_pos := keypad_lookup[prev_move]

	pq := make(utils.PriorityQueue, 1)
	pq[0] = utils.NewItem(HState{
		robot1_pos:      start_pos,
		last_move:       "A",
		numbers_entered: "",
		instructions:    "",
	}, 0, 0)
	heap.Init(&pq)
	cache := map[string]int{}

	// Use a priority queue and return when we find the answer, IE get the right move
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*utils.Item)
		state := item.GetValue().(HState)
		cost := item.GetPriority()
		top_pos := state.robot1_pos
		last_move := state.last_move
		output := state.numbers_entered
		instructions := state.instructions

		// Check we're in a valid position
		if val, ok := keypad[top_pos]; !ok || val == "X" {
			continue
		}

		// Check if we've solved it and got the move
		if output == move {
			heuristicCache[key] = cost
			return cost
		}

		// Cache, when the position and last move are the same, we've seen this before
		seen_key := fmt.Sprintf("%s%s", top_pos.String(), last_move)
		if _, ok := cache[seen_key]; ok {
			continue
		}
		cache[seen_key] = cost

		for _, move := range []string{"^", "v", "<", ">", "A"} {
			new_pos := top_pos
			new_pos, new_output := doMovePad(new_pos, move, keypad)
			new_instructions := instructions + move
			cost_move := heuristic(move, last_move, layers-1, keypad_lookup, keypad)
			new_cost := cost + cost_move
			newItem := utils.NewItem(HState{
				robot1_pos:      new_pos,
				last_move:       move,
				numbers_entered: output + new_output,
				instructions:    new_instructions,
			}, new_cost, 0)
			heap.Push(&pq, newItem)
			pq.Update(newItem, new_cost)
		}
	}
	panic("no solution")
}

func doKeypad(pos utils.Position, move string, keypad utils.Board) (utils.Position, string) {
	switch move {
	case "A":
		return pos, keypad[pos]
	case "v":
		return pos.Move(utils.Down, 1), ""
	case "^":
		return pos.Move(utils.Up, 1), ""
	case "<":
		return pos.Move(utils.Left, 1), ""
	case ">":
		return pos.Move(utils.Right, 1), ""
	}
	panic("not handled")
}

func doMovePad(pos utils.Position, move string, movepad utils.Board) (utils.Position, string) {
	switch move {
	case "A":
		return pos, movepad[pos]
	case "v":
		return pos.Move(utils.Down, 1), ""
	case "^":
		return pos.Move(utils.Up, 1), ""
	case "<":
		return pos.Move(utils.Left, 1), ""
	case ">":
		return pos.Move(utils.Right, 1), ""
	}
	panic("not handled")
}

func main() {
	input := utils.ReadInput("input.txt")
	solve(input, 2)
	solve(input, 25)
}
