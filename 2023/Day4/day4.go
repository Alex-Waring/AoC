package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func main() {
	defer utils.Timer("main")()
	lines := utils.ReadInput("Input.txt")

	card_points := []int{}
	line_matches := map[int]int{}

	// Part 1
	for game_no, line := range lines {
		round_points := 0
		line = strings.Split(line, ": ")[1]
		winning_numbers := utils.RemoveSliceSpaces(strings.Split(strings.Split(line, " | ")[0], " "))
		card_numbers := utils.RemoveSliceSpaces(strings.Split(strings.Split(line, " | ")[1], " "))
		count_wins := 0

		for _, number := range card_numbers {
			if slices.Contains(winning_numbers, number) {
				count_wins++
				if round_points == 0 {
					round_points++
				} else {
					round_points = round_points * 2
				}
			}
		}
		// Storing the number of matches so we can use the result in part 2
		line_matches[game_no+1] = count_wins
		card_points = append(card_points, round_points)
	}
	fmt.Println(utils.Sum(card_points))

	// Part 2, queue based but used the cached results from part 1
	queue := utils.MakeRange(1, len(lines))
	won_rounds := []int{}

	for len(queue) > 0 {
		processed_game := queue[0]
		games_won := line_matches[processed_game]

		for i := 1; i <= games_won; i++ {
			queue = append(queue, (i + processed_game))
		}
		won_rounds = append(won_rounds, processed_game)
		queue = queue[1:]
	}
	fmt.Println(len(won_rounds))
}
