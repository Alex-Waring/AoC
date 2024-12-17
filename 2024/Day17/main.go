package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Stack struct {
	A int
	B int
	C int
}

func part1(A int, B int, C int, program []int) string {
	// defer utils.Timer("part1")()

	pointer := 0
	stack := Stack{A, B, C}
	result := ""

	for {
		if pointer >= len(program) {
			break
		}

		opcode := program[pointer]
		operand := program[pointer+1]

		new_pointer, str := performInstruction(opcode, operand, &stack, pointer)
		pointer = new_pointer
		result += str
	}
	return result
}

func performInstruction(opcode int, operand int, stack *Stack, pointer int) (int, string) {
	str := ""
	switch opcode {
	case 0:
		result := math.Floor(float64(float64(stack.A) / (math.Pow(float64(2), float64(comboOperation(operand, stack))))))
		stack.A = int(result)
	case 1:
		result := stack.B ^ operand
		stack.B = result
	case 2:
		result := comboOperation(operand, stack) % 8
		stack.B = result
	case 3:
		if stack.A != 0 {
			return operand, str
		}
	case 4:
		result := stack.B ^ stack.C
		stack.B = result
	case 5:
		result := comboOperation(operand, stack) % 8
		str = fmt.Sprintf("%d,", result)
	case 6:
		result := math.Floor(float64(float64(stack.A) / (math.Pow(float64(2), float64(comboOperation(operand, stack))))))
		stack.B = int(result)
	case 7:
		result := math.Floor(float64(float64(stack.A) / (math.Pow(float64(2), float64(comboOperation(operand, stack))))))
		stack.C = int(result)
	}
	return pointer + 2, str
}

func comboOperation(operand int, stack *Stack) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return stack.A
	case 5:
		return stack.B
	case 6:
		return stack.C
	}
	panic("Invalid operand")
}

func part2(input []string) {
	defer utils.Timer("part2")()
}

func main() {
	input := utils.ReadInput("input.txt")

	var A int
	var B int
	var C int
	var program []int

	var program_string string

	fmt.Sscanf(input[0], "Register A: %d", &A)
	fmt.Sscanf(input[1], "Register B: %d", &B)
	fmt.Sscanf(input[2], "Register C: %d", &C)
	program_string = strings.Split(input[4], ": ")[1]

	for _, i := range strings.Split(program_string, ",") {
		program = append(program, utils.IntegerOf(i))
	}

	p1 := part1(A, B, C, program)
	fmt.Println(p1)

	target := program_string + ","

	// Part 2

	a_to_start := 0

	for {
		var result string
	BRUTE_LOOP:
		for i := range 8 {
			a := a_to_start + i
			result = part1(a, B, C, program)
			if strings.HasSuffix(target, result) {
				a_to_start = a * 8
				fmt.Printf("%d: %s\n", a, result)
				break BRUTE_LOOP
			}
		}
		// At this point we're nearly there and need to brute force the last few values
		if result == strings.TrimPrefix(target, "2,") {
			break
		}
	}

	for {
		a_to_start++
		result := part1(a_to_start, B, C, program)
		if result == target {
			fmt.Printf("%d: %s\n", a_to_start, result)
			break
		}
	}
}
