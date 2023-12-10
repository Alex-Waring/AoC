package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type Pipe struct {
	pipe         rune
	directions   [][]int
	value        int
	visited      bool
	map_value    string
	loop         bool
	scanned_loop string
}

func createPipe(pipe rune) Pipe {
	if pipe == '|' {
		return Pipe{
			pipe: '|', directions: [][]int{{1, 0}, {-1, 0}}, value: 0, visited: false, map_value: "|", loop: false,
		}
	} else if pipe == '-' {
		return Pipe{
			pipe: '-', directions: [][]int{{0, 1}, {0, -1}}, value: 0, visited: false, map_value: "-", loop: false,
		}
	} else if pipe == 'L' {
		return Pipe{
			pipe: 'L', directions: [][]int{{-1, 0}, {0, 1}}, value: 0, visited: false, map_value: "L", loop: false,
		}
	} else if pipe == 'J' {
		return Pipe{
			pipe: 'J', directions: [][]int{{-1, 0}, {0, -1}}, value: 0, visited: false, map_value: "J", loop: false,
		}
	} else if pipe == '7' {
		return Pipe{
			pipe: '7', directions: [][]int{{1, 0}, {0, -1}}, value: 0, visited: false, map_value: "7", loop: false,
		}
	} else if pipe == 'F' {
		return Pipe{
			pipe: 'F', directions: [][]int{{1, 0}, {0, 1}}, value: 0, visited: false, map_value: "F", loop: false,
		}
	} else if pipe == '.' {
		return Pipe{
			pipe: '.', directions: [][]int{}, value: 0, visited: false, map_value: ".", loop: false,
		}
	} else {
		return Pipe{
			pipe: 'S', directions: [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}, value: 0, visited: false, map_value: "S", loop: true,
		}
	}
}

func translateCoords(one []int, two []int) []int {
	return []int{one[0] + two[0], one[1] + two[1]}
}

func connect(dest_pipe Pipe, connect []int) bool {
	for _, connection := range dest_pipe.directions {
		if connection[0]*-1 == connect[0] && connection[1]*-1 == connect[1] {
			return true
		}
	}
	return false
}

func main() {
	lines := utils.ReadInput("input.txt")
	grid_pipes := make(map[int]map[int]Pipe)
	start := []int{0, 0}

	fmt.Println(len(lines[1]))
	for y, line := range lines {
		for x, pipe := range line {
			if _, ok := grid_pipes[y]; ok {
				grid_pipes[y][x] = createPipe(pipe)
			} else {
				grid_pipes[y] = make(map[int]Pipe)
				grid_pipes[y][x] = createPipe(pipe)
			}
			if pipe == 'S' {
				start = []int{y, x}
			}
		}
	}

	current_coords := [][]int{start}
	jumps := []int{}
	for len(current_coords) != 0 {
		processed_pipe := current_coords[0]
		pipe := grid_pipes[processed_pipe[0]][processed_pipe[1]]
		pipe.visited = true
		grid_pipes[processed_pipe[0]][processed_pipe[1]] = pipe
		for _, next := range pipe.directions {
			new_coord := translateCoords(processed_pipe, next)
			new_pipe := grid_pipes[new_coord[0]][new_coord[1]]
			if connect(new_pipe, next) && (new_pipe.value <= pipe.value) && !new_pipe.visited {
				current_coords = append(current_coords, new_coord)
				new_pipe.value = pipe.value + 1
				new_pipe.map_value = fmt.Sprint(new_pipe.value)
				new_pipe.loop = true
				grid_pipes[new_coord[0]][new_coord[1]] = new_pipe
			}
		}
		current_coords = current_coords[1:]
	}

	for y := 0; y < len(grid_pipes); y++ {
		for x := 0; x < len(grid_pipes[y]); x++ {
			if grid_pipes[y][x].loop {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
			jumps = append(jumps, grid_pipes[y][x].value)
		}
		fmt.Println()
	}
	for y := 0; y < len(grid_pipes); y++ {
		for x := 0; x < len(grid_pipes[y]); x++ {
			pipe := grid_pipes[y][x]
			if pipe.loop {
				if pipe.pipe == '-' || pipe.pipe == 'L' || pipe.pipe == 'J' {
					pipe.scanned_loop = "."
					grid_pipes[y][x] = pipe
					fmt.Print(".")
				} else if pipe.pipe == 'S' {
					fmt.Print("S")
				} else {
					pipe.scanned_loop = "*"
					grid_pipes[y][x] = pipe
					fmt.Print("*")
				}
			} else {
				pipe.scanned_loop = "."
				grid_pipes[y][x] = pipe
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	inside := 0
	for y := 0; y < len(grid_pipes); y++ {
		crossed := 0
		for x := 0; x < len(grid_pipes[y]); x++ {
			pipe := grid_pipes[y][x]
			if pipe.scanned_loop == "*" {
				crossed++
				fmt.Print("*")
			} else if pipe.loop {
				fmt.Print(string(pipe.pipe))
			} else if crossed%2 == 0 {
				fmt.Print("O")
			} else {
				inside++
				fmt.Print("I")
			}
		}
		fmt.Println()
	}
	fmt.Println(inside)
}
