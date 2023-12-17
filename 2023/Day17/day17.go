package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func solve(board map[utils.Position]int, target utils.Position, minStraight int, maxStraight int) int {
	type state struct {
		concurrent_straight int
		location            utils.Location
	}
	type cache struct {
		state
		heatLoss int
	}

	// Create a queue that will order based on the lowest heat loss
	q := priorityqueue.NewWith(func(a, b any) int {
		p1 := a.(cache).heatLoss
		p2 := b.(cache).heatLoss
		return p1 - p2
	})

	// Create two new starting positions on the queue, skipping the first position
	q.Enqueue(cache{
		state: state{
			concurrent_straight: 1,
			location:            utils.NewLocation(0, 1, utils.Right),
		},
	})
	q.Enqueue(cache{
		state: state{
			concurrent_straight: 1,
			location:            utils.NewLocation(1, 0, utils.Down),
		},
	})
	visited := make(map[state]int)

	for !q.Empty() {
		c, _ := q.Dequeue()
		e := c.(cache)
		position := e.location.Pos

		// If the board does not have this position, skip
		if _, exists := board[position]; !exists {
			continue
		}

		heat := board[position] + e.heatLoss
		// The queue is ordered, so if we are at the target we have the lowest heat loss
		// We also need to check that we have passed the min number of striaghts for the ultra
		// crucible
		if position == target && e.concurrent_straight >= minStraight {
			fmt.Println(e.concurrent_straight)
			return heat
		}

		// If we have already been here, with less heat, then end the process
		if previos_heat_loss, exists := visited[e.state]; exists {
			if previos_heat_loss <= heat {
				continue
			}
		}
		// Cache the result
		visited[e.state] = heat

		// go left and right if we have moved more than the min number of straight
		if e.concurrent_straight >= minStraight {
			q.Enqueue(cache{
				state: state{
					location:            e.location.Turn(utils.Left, 1),
					concurrent_straight: 1,
				},
				heatLoss: heat,
			})

			q.Enqueue(cache{
				state: state{
					location:            e.location.Turn(utils.Right, 1),
					concurrent_straight: 1,
				},
				heatLoss: heat,
			})
		}

		// If we havn't done the max straight, go straight
		if e.concurrent_straight < maxStraight {
			q.Enqueue(cache{
				state: state{
					location:            e.location.Straight(1),
					concurrent_straight: e.concurrent_straight + 1,
				},
				heatLoss: heat,
			})
		}
	}
	panic("you screwed up")
}

func main() {
	input := utils.ReadInput("input.txt")
	board := make(map[utils.Position]int, len(input)*len(input[0]))
	for row, line := range input {
		runes := []rune(line)
		for col, c := range runes {
			board[utils.NewPosition(row, col)] = int(c - '0')
		}
	}
	target := utils.Position{Row: len(input) - 1, Col: len(input[0]) - 1}

	part1_answer := solve(board, target, 0, 3)
	fmt.Println(part1_answer)

	part2_answer := solve(board, target, 4, 10)
	fmt.Println(part2_answer)
}
