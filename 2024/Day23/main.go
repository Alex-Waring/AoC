package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Node struct {
	name  string
	edges map[string]bool
}

func part1(input []string) {
	defer utils.Timer("part1")()

	nodes := make(map[string]Node)

	for _, line := range input {
		node1 := strings.Split(line, "-")[0]
		node2 := strings.Split(line, "-")[1]

		if _, ok := nodes[node1]; !ok {
			nodes[node1] = Node{name: node1, edges: map[string]bool{node2: true}}
		} else {
			nodes[node1].edges[node2] = true
		}

		if _, ok := nodes[node2]; !ok {
			nodes[node2] = Node{name: node2, edges: map[string]bool{node1: true}}
		} else {
			nodes[node2].edges[node1] = true
		}
	}

	// Networks is a set of alphabetically sorted strings of connected nodes
	networks := make(map[[3]string]bool)

	// For every node, loop through the pairs in it's edges and see if they are connected
	for _, node := range nodes {

		for edge1 := range node.edges {
			for edge2 := range node.edges {
				if edge1 == edge2 {
					continue
				}
				// Check if edge1 is connected to edge2
				if _, ok := nodes[edge1].edges[edge2]; ok {
					// Sort the strings alphabetically
					network := [3]string{node.name, edge1, edge2}
					slices.Sort(network[:])
					networks[network] = true
				}
			}
		}
	}

	total := 0
NETWORK_LOOP:
	for network := range networks {
		for _, node := range network {
			if strings.HasPrefix(node, "t") {
				total++
				continue NETWORK_LOOP
			}
		}
	}
	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	nodes := make(map[string]Node)

	for _, line := range input {
		node1 := strings.Split(line, "-")[0]
		node2 := strings.Split(line, "-")[1]

		if _, ok := nodes[node1]; !ok {
			nodes[node1] = Node{name: node1, edges: map[string]bool{node2: true}}
		} else {
			nodes[node1].edges[node2] = true
		}

		if _, ok := nodes[node2]; !ok {
			nodes[node2] = Node{name: node2, edges: map[string]bool{node1: true}}
		} else {
			nodes[node2].edges[node1] = true
		}
	}

	cliques := make([]map[string]bool, 0)

	// Bron-Kerbosch algorithm
	// Initiated with R and X empty, and P containing all nodes
	R := make([]string, 0)
	P := make([]string, 0)
	X := make([]string, 0)
	for name := range nodes {
		P = append(P, name)
	}
	BronKerbosch(R, P, X, nodes, &cliques)

	// I shouldn't need to convert this to a set, but I got the algorithm wrong
	// and ended up with duplicates
	max := map[string]bool{}
	for _, clique := range cliques {
		if len(clique) > len(max) {
			max = clique
		}
	}

	clique := make([]string, 0)
	for node := range max {
		clique = append(clique, node)
	}

	slices.Sort(clique)
	fmt.Println(strings.Join(clique, ","))
}

// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func BronKerbosch(R []string, P []string, X []string, nodes map[string]Node, cliques *[]map[string]bool) {
	if len(P) == 0 && len(X) == 0 {
		// Report R as a maximal clique
		new_clique := make(map[string]bool)
		for _, node := range R {
			new_clique[node] = true
		}
		*cliques = append(*cliques, new_clique)
		return
	}
	for i := 0; i < len(P); i++ {
		node_name := P[i]
		newR := append(R, node_name) // R u {v}

		newP := make([]string, 0) // P n N(v)
		for _, neighbour := range P {
			if nodes[node_name].edges[neighbour] {
				newP = append(newP, neighbour)
			}
		}

		newX := make([]string, 0) // X n N(v)
		for _, neighbour := range X {
			if nodes[node_name].edges[neighbour] {
				newX = append(newX, neighbour)
			}
		}

		BronKerbosch(newR, newP, newX, nodes, cliques)
		P = utils.Remove(P, i)   // P \ {v}
		X = append(X, node_name) // X u {v}
		i--                      // backtrack
	}
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
