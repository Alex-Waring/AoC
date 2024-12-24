package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Wire struct {
	state int
	set   bool
}

type Gate struct {
	operation string
	inputs    []string
	output    string
}

func part1(input []string) {
	defer utils.Timer("part1")()

	wires := make(map[string]Wire)
	gates := make([]Gate, 0)

	for _, line := range input {
		if strings.Contains(line, ":") {
			name := strings.Split(line, ": ")[0]
			value := utils.IntegerOf(strings.Split(line, ": ")[1])
			wires[name] = Wire{state: value, set: true}
		} else if line == "" {
			continue
		} else {
			var op string
			var input1, input2, output string
			fmt.Sscanf(line, "%s %s %s -> %s", &input1, &op, &input2, &output)
			gates = append(gates, Gate{operation: op, inputs: []string{input1, input2}, output: output})

			for _, wire := range []string{input1, input2, output} {
				if _, ok := wires[wire]; !ok {
					wires[wire] = Wire{state: 0, set: false}
				}
			}
		}
	}

	outputs := make([]string, 0)

	for name := range wires {
		if strings.HasPrefix(name, "z") {
			outputs = append(outputs, name)
		}
	}

	q := utils.Queue[Gate]{}

	for _, gate := range gates {
		if wires[gate.inputs[0]].set && wires[gate.inputs[1]].set {
			q.Push(gate)
		}
	}

	for !q.IsEmpty() {
		gate := q.Pop()

		input1 := wires[gate.inputs[0]]
		input2 := wires[gate.inputs[1]]
		output := wires[gate.output]
		op := gate.operation

		// Shouldn't happen but no harm in checking
		if input1.set == false || input2.set == false {
			continue
		}

		output.state = performOperation(op, input1.state, input2.state)
		output.set = true

		wires[gate.output] = output

		for _, gate_to_add := range gates {
			if !utils.StringInSlice(gate.output, gate_to_add.inputs) {
				continue
			}
			if wires[gate_to_add.inputs[0]].set && wires[gate_to_add.inputs[1]].set {
				q.Push(gate_to_add)
			}
		}
	}

	keys := make([]string, 0, len(wires))
	for k := range wires {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	slices.Reverse(keys)
	result := ""
	for _, name := range keys {
		if strings.HasPrefix(name, "z") {
			result += fmt.Sprintf("%d", wires[name].state)
		}
	}
	value, _ := strconv.ParseInt(result, 2, 64)
	fmt.Println(result)
	fmt.Println(value)

}

func part2(input []string, to_swap int) {
	defer utils.Timer("part2")()

	wires := make(map[string]Wire)
	gates := make([]Gate, 0)

	for _, line := range input {
		if strings.Contains(line, ":") {
			name := strings.Split(line, ": ")[0]
			value := utils.IntegerOf(strings.Split(line, ": ")[1])
			wires[name] = Wire{state: value, set: true}
		} else if line == "" {
			continue
		} else {
			var op string
			var input1, input2, output string
			fmt.Sscanf(line, "%s %s %s -> %s", &input1, &op, &input2, &output)
			gates = append(gates, Gate{operation: op, inputs: []string{input1, input2}, output: output})

			for _, wire := range []string{input1, input2, output} {
				if _, ok := wires[wire]; !ok {
					wires[wire] = Wire{state: 0, set: false}
				}
			}
		}
	}

	keys := make([]string, 0, len(wires))
	for k := range wires {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	slices.Reverse(keys)
	result := ""
	add1 := ""
	add2 := ""
	for _, name := range keys {
		if strings.HasPrefix(name, "z") {
			result += fmt.Sprintf("%d", wires[name].state)
		} else if strings.HasPrefix(name, "x") {
			add1 += fmt.Sprintf("%d", wires[name].state)
		} else if strings.HasPrefix(name, "y") {
			add2 += fmt.Sprintf("%d", wires[name].state)
		}
	}
	value, _ := strconv.ParseInt(result, 2, 64)
	add1_value, _ := strconv.ParseInt(add1, 2, 64)
	add2_value, _ := strconv.ParseInt(add2, 2, 64)
	fmt.Printf("Expected adder: %d\n", add1_value+add2_value)
	fmt.Println(value)
}

func performOperation(op string, input1 int, input2 int) int {
	switch op {
	case "AND":
		return input1 & input2
	case "OR":
		return input1 | input2
	case "XOR":
		return input1 ^ input2
	default:
		panic(fmt.Sprintf("unknown operation %s", op))
	}
}

func checkFinished(wires map[string]Wire, outputs []string) bool {
	for _, output := range outputs {
		if !wires[output].set {
			return false
		}
	}
	return true
}

func main() {
	input := utils.ReadInput("input.txt")
	to_swap := os.Args[1]
	part1(input)
	part2(input, utils.IntegerOf(to_swap))
}
