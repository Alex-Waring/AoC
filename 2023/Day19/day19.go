package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Condition struct {
	attr         string
	greater_than string
	value        int
	destination  string
}

type Workflow struct {
	conditions        []Condition
	final_destination string
}

type Isolated_Workflow struct {
	attr         string
	greater_than string
	value        int
	destination  string
	fail_con     string
}

type Part map[string]int

type Attr struct {
	lower int
	upper int
}

type Part_Range struct {
	values        map[string]Attr
	next_workflow string
}

func edit_part(values map[string]Attr, attr string, lower int, upper int) map[string]Attr {
	return_map := map[string]Attr{}

	for key, value := range values {
		if key == attr {
			return_map[attr] = Attr{
				lower: lower,
				upper: upper,
			}
		} else {
			return_map[key] = value
		}
	}
	return return_map
}

func (w Isolated_Workflow) Process(part Part_Range) []Part_Range {
	if w.value == 0 {
		return []Part_Range{{
			values:        part.values,
			next_workflow: w.destination,
		}}
	}

	return_parts := []Part_Range{}

	previous_upper := part.values[w.attr].upper
	previous_lower := part.values[w.attr].lower

	if w.greater_than == "<" {
		// If the range includes part that would be successfull
		if part.values[w.attr].lower < w.value {
			// include the win con
			win_con := Part_Range{
				values:        edit_part(part.values, w.attr, previous_lower, w.value-1),
				next_workflow: w.destination,
			}
			return_parts = append(return_parts, win_con)

			// Include the fail con
			fail_con := Part_Range{
				values:        edit_part(part.values, w.attr, w.value, previous_upper),
				next_workflow: w.fail_con,
			}
			return_parts = append(return_parts, fail_con)
		} else {
			return_parts = append(return_parts, Part_Range{
				values:        part.values,
				next_workflow: w.fail_con,
			})
		}
	} else {
		if part.values[w.attr].upper > w.value {
			// include the win con
			win_con := Part_Range{
				values:        edit_part(part.values, w.attr, w.value+1, previous_upper),
				next_workflow: w.destination,
			}
			return_parts = append(return_parts, win_con)

			// Include the fail con
			fail_con := Part_Range{
				values:        edit_part(part.values, w.attr, previous_lower, w.value),
				next_workflow: w.fail_con,
			}
			return_parts = append(return_parts, fail_con)
		} else {
			return_parts = append(return_parts, Part_Range{
				values:        part.values,
				next_workflow: w.fail_con,
			})
		}
	}

	return return_parts
}

func (w Workflow) Process(part Part) string {
	for _, con := range w.conditions {
		if con.greater_than == "<" {
			if part[con.attr] < con.value {
				return con.destination
			}
		} else {
			if part[con.attr] > con.value {
				return con.destination
			}
		}
	}
	return w.final_destination
}

func parse_workflows2(raw_workflows []string) map[string]Isolated_Workflow {
	workflows := make(map[string]Isolated_Workflow)
	for _, workflow := range raw_workflows {
		name := strings.Split(workflow, "{")[0]
		steps := strings.Split(strings.Trim(strings.Split(workflow, "{")[1], "}"), ",")

		workflows[name+"_"+fmt.Sprintf("%d", len(steps)-1)] = Isolated_Workflow{destination: steps[len(steps)-1] + "_0"}

		for i := 0; i < len(steps)-1; i++ {
			attr := string(steps[i][0])
			greater_than := string(steps[i][1])
			value := strings.Split(strings.Split(steps[i], greater_than)[1], ":")[0]
			destination := strings.Split(steps[i], ":")[1]

			workflows[name+"_"+fmt.Sprintf("%d", i)] = Isolated_Workflow{
				attr:         attr,
				greater_than: greater_than,
				value:        utils.IntegerOf(value),
				destination:  destination + "_0",
				fail_con:     name + "_" + fmt.Sprintf("%d", i+1),
			}
		}

	}
	return workflows
}

func parse_workflows(raw_workflows []string) map[string]Workflow {
	worklows := make(map[string]Workflow)
	for _, workflow := range raw_workflows {
		name := strings.Split(workflow, "{")[0]
		steps := strings.Split(strings.Trim(strings.Split(workflow, "{")[1], "}"), ",")

		new_workflow := Workflow{final_destination: steps[len(steps)-1]}

		for i := 0; i < len(steps)-1; i++ {
			attr := string(steps[i][0])
			greater_than := string(steps[i][1])
			value := strings.Split(strings.Split(steps[i], greater_than)[1], ":")[0]
			destination := strings.Split(steps[i], ":")[1]

			new_workflow.conditions = append(new_workflow.conditions, Condition{
				attr:         attr,
				greater_than: greater_than,
				value:        utils.IntegerOf(value),
				destination:  destination,
			})
		}

		worklows[name] = new_workflow
	}
	return worklows
}

func part2(raw_workflows []string, raw_parts []string) {
	defer utils.Timer("part2")()
	worklows := parse_workflows2(raw_workflows)

	initial_part := Part_Range{
		values: map[string]Attr{
			"x": {lower: 1, upper: 4000},
			"m": {lower: 1, upper: 4000},
			"a": {lower: 1, upper: 4000},
			"s": {lower: 1, upper: 4000},
		},
		next_workflow: "in_0",
	}
	q := utils.Queue[Part_Range]{}
	q.Push(initial_part)

	accepted := []Part_Range{}

	for !q.IsEmpty() {
		part := q.Pop()
		results := worklows[part.next_workflow].Process(part)

		for _, result := range results {
			if result.next_workflow == "A_0" {
				accepted = append(accepted, result)
			} else if result.next_workflow == "R_0" {
				continue
			} else {
				q.Push(result)
			}
		}
		fmt.Println(q)
	}

	result := 0

	for _, part := range accepted {
		part_result := (part.values["x"].upper - part.values["x"].lower + 1) * (part.values["m"].upper - part.values["m"].lower + 1) * (part.values["a"].upper - part.values["a"].lower + 1) * (part.values["s"].upper - part.values["s"].lower + 1)
		result += part_result
	}
	fmt.Println(result)

}

func part1(raw_workflows []string, parts []string) {
	defer utils.Timer("part1")()
	worklows := parse_workflows(raw_workflows)
	parsed_parts := []Part{}

	for _, part := range parts {
		part = strings.Trim(part, "{")
		part = strings.Trim(part, "}")

		sections := strings.Split(part, ",")

		new_part := Part{}
		total := 0

		for _, section := range sections {
			key, value := strings.Split(section, "=")[0], strings.Split(section, "=")[1]
			total += utils.IntegerOf(value)
			new_part[key] = utils.IntegerOf(value)
		}
		new_part["total"] = total
		parsed_parts = append(parsed_parts, new_part)
	}

	result := 0
	for _, part := range parsed_parts {
		destination := worklows["in"].Process(part)

		for destination != "A" && destination != "R" {
			destination = worklows[destination].Process(part)
		}
		if destination == "A" {
			result += part["total"]
		}
	}
	fmt.Println(result)
}

func main() {
	raw_input, _ := os.ReadFile("input.txt")

	blocks := strings.Split(string(raw_input), "\n\n")

	workflows := utils.RemoveSliceSpaces(strings.Split(blocks[0], "\n"))
	parts := utils.RemoveSliceSpaces(strings.Split(blocks[1], "\n"))

	part1(workflows, parts)
	part2(workflows, parts)
}
