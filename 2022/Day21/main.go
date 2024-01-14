package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type operation int

const (
	add operation = iota
	multiply
	divide
	subtract
	equal
)

type Monkey struct {
	name      string
	number    int
	operation operation
	inputs    []string
}

func part1(monkeys map[string]Monkey) {
	defer utils.Timer("part1")()
	found_root := false

	for !found_root {
		for _, monkey := range monkeys {
			if monkey.number != 0 {
				continue
			}

			input_1 := monkeys[monkey.inputs[0]].number
			input_2 := monkeys[monkey.inputs[1]].number

			if input_1 != 0 && input_2 != 0 {
				// The source monkeys have shouted so we can calculate the number

				switch monkey.operation {
				case add:
					monkey.number = input_1 + input_2
				case multiply:
					monkey.number = input_1 * input_2
				case divide:
					monkey.number = input_1 / input_2
				case subtract:
					monkey.number = input_1 - input_2
				}
			}

			monkeys[monkey.name] = monkey

			if monkey.name == "root" && monkey.number != 0 {
				fmt.Println(monkey.number)
				found_root = true
				break
			}
		}
	}
}

func part2(monkeys map[string]Monkey, input []string) {
	defer utils.Timer("part2")()
	found_i := false

	// Do a linear search, we guess a number, if it's too large then we move to half the lower
	// bound, if it's too small we move to half the upper bound
	// This works because the input only affects the left side of the tree, and increases it
	// decreases the output for input, increased for example
	lower_bound := 0
	upper_bound := 100000000000000
	i := (upper_bound - lower_bound) / 2
	too_high := false
	hit_loop := false

	for !found_i {
		found_i, too_high, hit_loop = calculateLoops(monkeys, i)

		if found_i {
			fmt.Println(i - 1)
		}

		if hit_loop {
			i++
		} else if too_high {
			upper_bound = i
			i = lower_bound + ((upper_bound - lower_bound) / 2)
		} else {
			lower_bound = i
			i = lower_bound + ((upper_bound - lower_bound) / 2)
		}

		// Reset monkeys
		monkeys = setMonkeys(input)
	}
}

func calculateLoops(monkeys map[string]Monkey, human int) (bool, bool, bool) {
	found_root := false

	humn := monkeys["humn"]
	humn.number = human
	monkeys["humn"] = humn

	for !found_root {
		for _, monkey := range monkeys {
			if monkey.number != 0 {
				continue
			}

			input_1 := monkeys[monkey.inputs[0]].number
			input_2 := monkeys[monkey.inputs[1]].number

			if input_1 != 0 && input_2 != 0 {
				// The source monkeys have shouted so we can calculate the number

				// First check if this is root
				if monkey.name == "root" {
					// Number are equal, we've found the root
					if input_1 == input_2 {
						return true, true, false
					} else {
						// Numbers are not equal, we've failed
						if input_1 < input_2 {
							// We need to reduce the humn input
							return false, true, false
						} else {
							// We need to increase the humn input
							return false, false, false
						}
					}
				}

				switch monkey.operation {
				case add:
					monkey.number = input_1 + input_2
				case multiply:
					monkey.number = input_1 * input_2
				case divide:
					monkey.number = input_1 / input_2
				case subtract:
					monkey.number = input_1 - input_2
				}

				// If the number is still zero, we've hit a loop and can break
				// we should increase the humn input by 1 and try again
				if monkey.number == 0 {
					return false, true, true
				}
			}

			monkeys[monkey.name] = monkey

			if monkey.name == "root" && monkey.number != 0 {
				fmt.Println(monkey.number)
				found_root = true
				break
			}
		}
	}
	// Should never get here
	return false, false, false
}

func setMonkeys(input []string) map[string]Monkey {
	monkeys := map[string]Monkey{}
	for _, line := range input {
		name := strings.Split(line, ": ")[0]
		math := strings.Split(line, ": ")[1]

		if len(strings.Split(math, " ")) == 1 {
			number := utils.IntegerOf(math)
			monkeys[name] = Monkey{name: name, number: number}
		} else {
			operation := operation(0)
			if strings.Contains(math, "+") {
				operation = add
			} else if strings.Contains(math, "*") {
				operation = multiply
			} else if strings.Contains(math, "/") {
				operation = divide
			} else if strings.Contains(math, "-") {
				operation = subtract
			}

			input_1 := strings.Split(math, " ")[0]
			input_2 := strings.Split(math, " ")[2]

			monkeys[name] = Monkey{name: name, operation: operation, inputs: []string{input_1, input_2}}
		}
	}

	return monkeys
}

func main() {
	input := utils.ReadInput("input.txt")
	monkeys := setMonkeys(input)

	part1(monkeys)

	monkeys2 := setMonkeys(input)
	part2(monkeys2, input)
}
