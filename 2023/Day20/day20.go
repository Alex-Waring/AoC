package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type FlipFlop struct {
	on bool
}

type Conjucation struct {
	recieved map[string]int
}

type Module struct {
	// b, f, or c
	moduleType  string
	name        string
	flipFlip    FlipFlop
	conjucation Conjucation
	destination []string
}

type Packet struct {
	module *Module
	pulse  int
	source string
}

func push_button(modules map[string]*Module, loop int) (int, int) {
	q := utils.Queue[Packet]{}

	low_pulses := 0
	high_pulses := 0

	// Push the button
	low_pulses++
	for _, destination := range modules["broadcast"].destination {
		q.Push(Packet{
			module: modules[destination],
			pulse:  0,
			source: "broadcast",
		})
		low_pulses++
	}

	for !q.IsEmpty() {
		packet := q.Pop()

		if packet.module.moduleType == "b" {
			for _, destination := range packet.module.destination {
				q.Push(Packet{
					module: modules[destination],
					pulse:  packet.pulse,
					source: packet.module.name,
				})
				if packet.pulse == 1 {
					high_pulses++
				} else {
					low_pulses++
				}
			}
		} else if packet.module.moduleType == "f" {
			// Flip Flops ignore high pulses
			if packet.pulse == 0 {
				// If on, send a low pulse and turn off
				if packet.module.flipFlip.on {
					for _, destination := range packet.module.destination {
						q.Push(Packet{
							module: modules[destination],
							pulse:  0,
							source: packet.module.name,
						})
						low_pulses++
					}
					packet.module.flipFlip.on = false
				} else {
					// if off, send a high pulse and turn on
					for _, destination := range packet.module.destination {
						q.Push(Packet{
							module: modules[destination],
							pulse:  1,
							source: packet.module.name,
						})
						high_pulses++
					}
					packet.module.flipFlip.on = true
				}
			}
		} else if packet.module.moduleType == "c" {
			packet.module.conjucation.recieved[packet.source] = packet.pulse

			// If it remembers high for all, send a low, otherwise high
			return_pulse := 0
			for _, pulse := range packet.module.conjucation.recieved {
				if pulse == 0 {
					return_pulse = 1
				}
			}

			for _, destination := range packet.module.destination {
				q.Push(Packet{
					module: modules[destination],
					pulse:  return_pulse,
					source: packet.module.name,
				})
				if return_pulse == 1 {
					high_pulses++
				} else {
					low_pulses++
				}
			}
		}
	}

	return low_pulses, high_pulses
}

func part2(modules map[string]*Module, loop int, low_counts map[string]int) map[string]int {
	q := utils.Queue[Packet]{}

	low_pulses := 0
	high_pulses := 0

	sources := []string{"nh", "xm", "tr", "dr"}

	// Push the button
	low_pulses++
	for _, destination := range modules["broadcast"].destination {
		q.Push(Packet{
			module: modules[destination],
			pulse:  0,
			source: "broadcast",
		})
		low_pulses++
	}

	for !q.IsEmpty() {
		packet := q.Pop()

		if utils.StringInSlice(packet.module.name, sources) {
			if packet.pulse == 0 {
				if _, ok := low_counts[packet.module.name]; !ok {
					low_counts[packet.module.name] = loop
				}
			}
		}

		if packet.module.moduleType == "b" {
			for _, destination := range packet.module.destination {
				q.Push(Packet{
					module: modules[destination],
					pulse:  packet.pulse,
					source: packet.module.name,
				})
				if packet.pulse == 1 {
					high_pulses++
				} else {
					low_pulses++
				}
			}
		} else if packet.module.moduleType == "f" {
			// Flip Flops ignore high pulses
			if packet.pulse == 0 {
				// If on, send a low pulse and turn off
				if packet.module.flipFlip.on {
					for _, destination := range packet.module.destination {
						q.Push(Packet{
							module: modules[destination],
							pulse:  0,
							source: packet.module.name,
						})
						low_pulses++
					}
					packet.module.flipFlip.on = false
				} else {
					// if off, send a high pulse and turn on
					for _, destination := range packet.module.destination {
						q.Push(Packet{
							module: modules[destination],
							pulse:  1,
							source: packet.module.name,
						})
						high_pulses++
					}
					packet.module.flipFlip.on = true
				}
			}
		} else if packet.module.moduleType == "c" {
			packet.module.conjucation.recieved[packet.source] = packet.pulse

			// If it remembers high for all, send a low, otherwise high
			return_pulse := 0
			for _, pulse := range packet.module.conjucation.recieved {
				if pulse == 0 {
					return_pulse = 1
				}
			}

			for _, destination := range packet.module.destination {
				q.Push(Packet{
					module: modules[destination],
					pulse:  return_pulse,
					source: packet.module.name,
				})
				if return_pulse == 1 {
					high_pulses++
				} else {
					low_pulses++
				}
			}
		}
	}
	return low_counts
}

func main() {
	input := utils.ReadInput("input.txt")
	modules := map[string]*Module{}

	// hard coding this because
	output_module := Module{
		moduleType: "o",
		name:       "output",
	}
	modules["rx"] = &output_module

	for _, line := range input {
		sections := strings.Split(line, "->")

		if strings.HasPrefix(sections[0], "broadcast") {
			new_module := Module{
				moduleType: "b",
				name:       "broadcast",
			}
			destinations := strings.Split(sections[1], ",")
			for _, destination := range destinations {
				new_module.destination = append(new_module.destination, strings.Replace(destination, " ", "", -1))
			}
			modules["broadcast"] = &new_module
		} else if strings.HasPrefix(sections[0], "%") {
			name := strings.Trim(strings.Trim(sections[0], "%"), " ")
			new_module := Module{
				moduleType: "f",
				name:       name,
				flipFlip:   FlipFlop{on: false},
			}
			destinations := strings.Split(sections[1], ",")
			for _, destination := range destinations {
				new_module.destination = append(new_module.destination, strings.Replace(destination, " ", "", -1))
			}
			modules[name] = &new_module
		} else if strings.HasPrefix(sections[0], "&") {
			name := strings.Trim(strings.Trim(sections[0], "&"), " ")
			new_module := Module{
				moduleType: "c",
				name:       name,
				conjucation: Conjucation{
					recieved: make(map[string]int),
				},
			}
			destinations := strings.Split(sections[1], ",")
			for _, destination := range destinations {
				new_module.destination = append(new_module.destination, strings.Replace(destination, " ", "", -1))
			}
			modules[name] = &new_module
		}
	}

	// Prep all the conjucation modules
	for _, module := range modules {
		for _, destination := range module.destination {
			if modules[destination].moduleType == "c" {
				modules[destination].conjucation.recieved[module.name] = 0
			}
		}
	}

	// total_low := 0
	// total_high := 0
	// for i := 0; i < 1000; i++ {
	// 	low, high := push_button(modules, i+1)
	// 	total_high += high
	// 	total_low += low
	// }
	// fmt.Println(total_high * total_low)

	// Comment out part 2 to do part 1
	low_counts := make(map[string]int)
	i := 1
	for len(low_counts) < 4 {
		low_counts = part2(modules, i, low_counts)
		i++
	}
	fmt.Println(low_counts)
	fmt.Println(utils.LCM(low_counts["dr"], low_counts["nh"], low_counts["xm"], low_counts["tr"]))
}
