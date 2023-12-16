package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

func hash(step string) int {
	current_value := 0
	for _, charecter := range step {
		ascii := int(charecter)
		current_value += ascii
		current_value = current_value * 17
		current_value = current_value % 256
	}
	return current_value
}

func part1(input string) {
	defer utils.Timer("part1")()

	sequence := strings.Split(input, ",")
	total := 0

	for _, step := range sequence {
		total += hash(step)
	}

	fmt.Println(total)
}

func part2(input string) {
	defer utils.Timer("part1")()

	sequence := strings.Split(input, ",")
	boxes := map[int][]string{}

	for _, step := range sequence {
		if strings.Contains(step, "=") {
			label := strings.Split(step, "=")[0]
			box := hash(label)
			focal_length := strings.Split(step, "=")[1]

			if contents, ok := boxes[box]; ok {
				label_in_box := false

				for i := 0; i < len(contents) && !label_in_box; i++ {
					if strings.HasPrefix(boxes[box][i], label) {
						boxes[box][i] = label + " " + focal_length
						label_in_box = true
					}
				}

				if !label_in_box {
					boxes[box] = append(boxes[box], label+" "+focal_length)
				}
			} else {
				boxes[box] = []string{label + " " + focal_length}
			}
		} else {
			label := strings.Split(step, "-")[0]
			box := hash(label)
			if contents, ok := boxes[box]; ok {
				label_in_box := false

				for i := 0; i < len(contents) && !label_in_box; i++ {
					if strings.HasPrefix(boxes[box][i], label) {
						boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
						label_in_box = true
					}
				}
			}
		}
	}

	values := []int{}
	for box_no, contents := range boxes {
		for slot, lense := range contents {
			slot_value := (1 + box_no) * (1 + slot) * utils.IntegerOf(strings.Split(lense, " ")[1])
			values = append(values, slot_value)
		}
	}

	fmt.Println(utils.Sum(values))
}

func main() {
	input := utils.ReadInput("input.txt")

	part1(input[0])
	part2(input[0])
}
