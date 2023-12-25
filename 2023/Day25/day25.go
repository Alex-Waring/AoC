package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
	"gonum.org/v1/gonum/graph/network"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	input := utils.ReadInput("input.txt")
	ids := map[string]int64{}

	g := simple.NewUndirectedGraph()

	for _, line := range input {
		vertex := strings.Split(line, ":")[0]
		var v_id int64

		if id, exists := ids[vertex]; exists {
			v_id = id
		} else {
			v_id = int64(len(ids))
			ids[vertex] = int64(v_id)
		}

		edges := strings.Split(strings.TrimLeft(strings.Split(line, ":")[1], " "), " ")
		for _, edge := range edges {
			var e_id int64
			if id, exists := ids[edge]; exists {
				e_id = id
			} else {
				e_id = int64(len(ids))
				ids[edge] = int64(e_id)
			}

			g.SetEdge(simple.Edge{
				F: simple.Node(v_id),
				T: simple.Node(e_id),
			})
		}
	}

	betweenness_map := network.Betweenness(g)
	keys := make([]int, 0, len(betweenness_map))

	for key := range betweenness_map {
		keys = append(keys, int(key))
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return betweenness_map[int64(keys[i])] < betweenness_map[int64(keys[j])]
	})

	// After sorting by 'betweenness' we have 6 nodes that are more in the middle than others
	// Split them and we have two systems
	node_1 := keys[len(keys)-6]
	node_2 := keys[len(keys)-5]
	node_3 := keys[len(keys)-4]
	node_4 := keys[len(keys)-3]
	node_5 := keys[len(keys)-2]
	node_6 := keys[len(keys)-1]

	g.RemoveEdge(int64(node_1), int64(node_2))
	g.RemoveEdge(int64(node_4), int64(node_3))
	g.RemoveEdge(int64(node_5), int64(node_6))

	community1 := path.DijkstraAllFrom(g.Node(int64(node_1)), g)

	nodes_1 := 0
	nodes_2 := 0

	// If we can reach a node, it's in one community, if we can't it's in another.
	// THis works for any node
	for _, key := range keys {
		path, _ := community1.AllTo(int64(key))

		if len(path) > 0 {
			nodes_1++
		} else {
			nodes_2++
		}
	}

	fmt.Println(nodes_1 * nodes_2)
}
