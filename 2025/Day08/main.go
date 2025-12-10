package main

import (
	"fmt"
	"sort"

	"github.com/Alex-Waring/AoC/utils"
)

// Only bother calculating distance squared as it's faster and we're just comparing sizes
func threeDEuclideanDistanceSquared(a, b [3]int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	dz := a[2] - b[2]
	return dx*dx + dy*dy + dz*dz
}

type connection struct {
	p1, p2   [3]int
	distance int
}

func part1(input []string) {
	defer utils.Timer("part1")()
	puzzleInput := [][3]int{}
	for _, line := range input {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		puzzleInput = append(puzzleInput, [3]int{x, y, z})
	}

	// Find all possible connections and sort by distance
	allConnections := []connection{}
	for i, a := range puzzleInput {
		for j, b := range puzzleInput {
			if i < j { // Only consider each pair once
				dist := threeDEuclideanDistanceSquared(a, b)
				allConnections = append(allConnections, connection{a, b, dist})
			}
		}
	}

	// Sort by distance
	sort.Slice(allConnections, func(i, j int) bool {
		return allConnections[i].distance < allConnections[j].distance
	})

	circuits := [][][3]int{}
	inCircuit := map[[3]int]bool{}
	circuitLookup := map[[3]int]int{}

	for i := 0; i < 1000; i++ {
		conn := allConnections[i]
		p1, p2 := conn.p1, conn.p2

		// Check if already in same circuit
		if inCircuit[p1] && inCircuit[p2] && circuitLookup[p1] == circuitLookup[p2] {
			// fmt.Println("Connection", i+1, "skipped - already in same circuit")
			continue
		}

		if inCircuit[p1] && inCircuit[p2] {
			// merge circuits
			circuit1Index := circuitLookup[p1]
			circuit2Index := circuitLookup[p2]

			// Skip issues of trying to merge in wrong order
			if circuit1Index > circuit2Index {
				circuit1Index, circuit2Index = circuit2Index, circuit1Index
			}

			circuits[circuit1Index] = append(circuits[circuit1Index], circuits[circuit2Index]...)
			for _, point := range circuits[circuit2Index] {
				circuitLookup[point] = circuit1Index
			}
			circuits = append(circuits[:circuit2Index], circuits[circuit2Index+1:]...)
			for idx := circuit2Index; idx < len(circuits); idx++ {
				for _, point := range circuits[idx] {
					circuitLookup[point] = idx
				}
			}
		} else if inCircuit[p1] {
			circuitIndex := circuitLookup[p1]
			circuits[circuitIndex] = append(circuits[circuitIndex], p2)
			inCircuit[p2] = true
			circuitLookup[p2] = circuitIndex
		} else if inCircuit[p2] {
			circuitIndex := circuitLookup[p2]
			circuits[circuitIndex] = append(circuits[circuitIndex], p1)
			inCircuit[p1] = true
			circuitLookup[p1] = circuitIndex
		} else {
			newCircuit := [][3]int{p1, p2}
			circuits = append(circuits, newCircuit)
			inCircuit[p1] = true
			inCircuit[p2] = true
			circuitLookup[p1] = len(circuits) - 1
			circuitLookup[p2] = len(circuits) - 1
		}
		// fmt.Println("Connection", i+1, "made")
	}
	// fmt.Println("Circuits: ", circuits)

	threeLargersCircuits := [3]int{0, 0, 0}
	for _, circuit := range circuits {
		if len(circuit) > threeLargersCircuits[0] {
			threeLargersCircuits[2] = threeLargersCircuits[1]
			threeLargersCircuits[1] = threeLargersCircuits[0]
			threeLargersCircuits[0] = len(circuit)
		} else if len(circuit) > threeLargersCircuits[1] {
			threeLargersCircuits[2] = threeLargersCircuits[1]
			threeLargersCircuits[1] = len(circuit)
		} else if len(circuit) > threeLargersCircuits[2] {
			threeLargersCircuits[2] = len(circuit)
		}
	}
	fmt.Println("Part 1: ", threeLargersCircuits[0]*threeLargersCircuits[1]*threeLargersCircuits[2])
}

func largestCircuitSize(circuits [][][3]int) int {
	largest := 0
	for _, circuit := range circuits {
		if len(circuit) > largest {
			largest = len(circuit)
		}
	}
	return largest
}

func part2(input []string) {
	defer utils.Timer("part2")()

	puzzleInput := [][3]int{}
	for _, line := range input {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		puzzleInput = append(puzzleInput, [3]int{x, y, z})
	}

	// Find all possible connections and sort by distance
	allConnections := []connection{}
	for i, a := range puzzleInput {
		for j, b := range puzzleInput {
			if i < j { // Only consider each pair once
				dist := threeDEuclideanDistanceSquared(a, b)
				allConnections = append(allConnections, connection{a, b, dist})
			}
		}
	}

	// Sort by distance
	sort.Slice(allConnections, func(i, j int) bool {
		return allConnections[i].distance < allConnections[j].distance
	})

	circuits := [][][3]int{}
	inCircuit := map[[3]int]bool{}
	circuitLookup := map[[3]int]int{}

	// Process until all points are connected
	for i := 0; i < len(allConnections); i++ {
		conn := allConnections[i]
		p1, p2 := conn.p1, conn.p2

		// Check if already in same circuit
		if inCircuit[p1] && inCircuit[p2] && circuitLookup[p1] == circuitLookup[p2] {
			// fmt.Println("Connection", i+1, "skipped - already in same circuit")
			continue
		}

		if inCircuit[p1] && inCircuit[p2] {
			// merge circuits
			circuit1Index := circuitLookup[p1]
			circuit2Index := circuitLookup[p2]

			// Skip issues of trying to merge in wrong order
			if circuit1Index > circuit2Index {
				circuit1Index, circuit2Index = circuit2Index, circuit1Index
			}

			circuits[circuit1Index] = append(circuits[circuit1Index], circuits[circuit2Index]...)
			for _, point := range circuits[circuit2Index] {
				circuitLookup[point] = circuit1Index
			}
			circuits = append(circuits[:circuit2Index], circuits[circuit2Index+1:]...)
			for idx := circuit2Index; idx < len(circuits); idx++ {
				for _, point := range circuits[idx] {
					circuitLookup[point] = idx
				}
			}
		} else if inCircuit[p1] {
			circuitIndex := circuitLookup[p1]
			circuits[circuitIndex] = append(circuits[circuitIndex], p2)
			inCircuit[p2] = true
			circuitLookup[p2] = circuitIndex
		} else if inCircuit[p2] {
			circuitIndex := circuitLookup[p2]
			circuits[circuitIndex] = append(circuits[circuitIndex], p1)
			inCircuit[p1] = true
			circuitLookup[p1] = circuitIndex
		} else {
			newCircuit := [][3]int{p1, p2}
			circuits = append(circuits, newCircuit)
			inCircuit[p1] = true
			inCircuit[p2] = true
			circuitLookup[p1] = len(circuits) - 1
			circuitLookup[p2] = len(circuits) - 1
		}

		if largestCircuitSize(circuits) == len(puzzleInput) {
			fmt.Println("All points connected after", i+1, "connections")
			fmt.Println(p1[0] * p2[0])
			break
		}

	}
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
