package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Machine struct {
	// Use bits so we can cache in the solve function
	lights  uint64
	target  uint64
	buttons [][]int
}

func (m Machine) Solve() int {
	type State struct {
		lights  uint64
		presses int
	}

	visited := map[State]bool{}

	queue := utils.Queue[State]{}
	queue.Push(State{lights: m.lights, presses: 0})

	for !queue.IsEmpty() {
		current := queue.PopFront()

		if current.lights == m.target {
			return current.presses
		}

		if visited[current] {
			continue
		}
		visited[current] = true

		// Try pressing each button
		for i := range m.buttons {
			newLights := current.lights
			buttons := m.buttons[i]
			for _, btnIdx := range buttons {
				newLights ^= (1 << btnIdx) // Toggle bit at btnIdx
			}
			queue.Push(State{lights: newLights, presses: current.presses + 1})
		}
	}
	return -1 // No solution found
}

func part1(input []string) {
	defer utils.Timer("part1")()

	machines := []Machine{}

	for _, line := range input {
		parts := strings.Split(line, " ")
		lightStr := parts[0][1 : len(parts[0])-1]
		buttonStrs := parts[1 : len(parts)-1]

		var lights uint64 = 0
		var target uint64 = 0

		for i, ch := range lightStr {
			if ch == '#' {
				target |= (1 << i) // Set bit i in target
			}
		}

		buttons := [][]int{}
		for _, btnStr := range buttonStrs {
			btnParts := strings.Split(btnStr[1:len(btnStr)-1], ",")
			btn := []int{}
			for _, part := range btnParts {
				btn = append(btn, utils.IntegerOf(part))
			}
			buttons = append(buttons, btn)
		}
		machines = append(machines, Machine{lights: lights, buttons: buttons, target: target})
	}

	results := []int{}
	for _, machine := range machines {
		result := machine.Solve()
		results = append(results, result)
	}

	fmt.Printf("Part 1: %v\n", utils.Sum(results))
}

type Machine2 struct {
	joltages            []int
	joltageRequirements []int
	buttons             [][]int
}

func (m Machine2) SolveZ3() int {
	var smt bytes.Buffer

	// declare vars (one integer per button)
	numButtons := len(m.buttons)
	for i := 0; i < numButtons; i++ {
		// (declare-const b0 Int) ...
		smt.WriteString(fmt.Sprintf("(declare-const b%d Int)\n", i))
	}

	// Constraints:
	for i := 0; i < numButtons; i++ {
		// (assert (>= b0 0)) ...
		smt.WriteString(fmt.Sprintf("(assert (>= b%d 0))\n", i))
	}

	// Convert joltage counter into sum of buttons
	numCounters := len(m.joltageRequirements)
	for cIdx := 0; cIdx < numCounters; cIdx++ {
		// Construct sum: (+ b0 b2 b5 ...)
		// Only include buttons that actually affect this counter (cIdx)
		var sumParts []string
		for bIdx, btnEffects := range m.buttons {
			// Check if button bIdx affects counter cIdx
			affects := false
			for _, eff := range btnEffects {
				if eff == cIdx {
					affects = true
					break
				}
			}
			if affects {
				sumParts = append(sumParts, fmt.Sprintf("b%d", bIdx))
			}
		}

		// (assert (= (+ b0 b1) 10))
		line := fmt.Sprintf("(assert (= (+ %s) %d))\n", strings.Join(sumParts, " "), m.joltageRequirements[cIdx])
		smt.WriteString(line)
	}

	// Goal is min total button presses
	var allButtons []string
	for i := 0; i < numButtons; i++ {
		allButtons = append(allButtons, fmt.Sprintf("b%d", i))
	}
	// (minimize (+ b0 b1 ...))
	smt.WriteString(fmt.Sprintf("(minimize (+ %s))\n", strings.Join(allButtons, " ")))

	smt.WriteString("(check-sat)\n")
	smt.WriteString("(get-model)\n")

	cmd := exec.Command("z3", "-in")
	cmd.Stdin = &smt
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running Z3:", err)
		return 0
	}
	output := string(outputBytes)

	// Output format is:
	// sat
	// (model
	//   (define-fun b0 () Int
	//     5)
	//   (define-fun b1 () Int
	//     10)
	// )

	re := regexp.MustCompile(`\(define-fun b(\d+) \(\) Int\s+(\d+)\)`)
	matches := re.FindAllStringSubmatch(output, -1)

	totalPresses := 0
	for _, match := range matches {
		// Get second int (first is button index)
		val, _ := strconv.Atoi(match[2])
		totalPresses += val
	}

	return totalPresses
}

func part2(input []string) {
	defer utils.Timer("part2")()

	machines := []Machine2{}

	for _, line := range input {
		parts := strings.Split(line, " ")
		buttonStrs := parts[1 : len(parts)-1]
		joltageStr := parts[len(parts)-1]
		joltageStr = joltageStr[1 : len(joltageStr)-1]
		joltageParts := strings.Split(joltageStr, ",")

		joltages := make([]int, len(joltageParts))
		joltageRequirements := make([]int, len(joltageParts))

		for i, str := range joltageParts {
			joltageRequirements[i] = utils.IntegerOf(str)
		}

		buttons := [][]int{}
		for _, btnStr := range buttonStrs {
			btnParts := strings.Split(btnStr[1:len(btnStr)-1], ",")
			btn := []int{}
			for _, part := range btnParts {
				btn = append(btn, utils.IntegerOf(part))
			}
			buttons = append(buttons, btn)
		}
		machines = append(machines, Machine2{joltages: joltages, joltageRequirements: joltageRequirements, buttons: buttons})
	}

	results := []int{}
	for i, machine := range machines {
		fmt.Printf("Solving machine %d/%d\n", i+1, len(machines))
		result := machine.SolveZ3()
		fmt.Printf("Result: %d\n", result)
		results = append(results, result)
	}

	fmt.Printf("Part 2: %v\n", utils.Sum(results))
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
