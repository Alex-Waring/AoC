package main

import (
	"fmt"
	"sort"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	corners := [][2]int{}

	for _, line := range input {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		corners = append(corners, [2]int{x, y})
	}

	// Largest rectangle should be one of two possible pairs:
	// - Clostest to top right (min (x-y)) and closest to bottom left (max (x-y))
	// - Closest to top left (min total) and closest to bottom right (max total)
	topRightIdx, bottomLeftIdx := 0, 0
	minDiff, maxDiff := corners[0][0]-corners[0][1], corners[0][0]-corners[0][1]
	topLeftIdx, bottomRightIdx := 0, 0
	minTotal, maxTotal := corners[0][0]+corners[0][1], corners[0][0]+corners[0][1]

	for i, corner := range corners {
		diff := corner[0] - corner[1]
		if diff < minDiff {
			minDiff = diff
			topRightIdx = i
		}
		if diff > maxDiff {
			maxDiff = diff
			bottomLeftIdx = i
		}

		total := corner[0] + corner[1]
		if total < minTotal {
			minTotal = total
			topLeftIdx = i
		}
		if total > maxTotal {
			maxTotal = total
			bottomRightIdx = i
		}
	}

	// The corners are inclusive, so we need to add 1 to the area calculation
	area1 := (corners[bottomLeftIdx][0] - corners[topRightIdx][0] + 1) * (corners[bottomLeftIdx][1] - corners[topRightIdx][1] + 1)
	area2 := (corners[bottomRightIdx][0] - corners[topLeftIdx][0] + 1) * (corners[bottomRightIdx][1] - corners[topLeftIdx][1] + 1)

	if area1 > area2 {
		fmt.Printf("Part 1: %d\n", area1)
	} else {
		fmt.Printf("Part 1: %d\n", area2)
	}
}

type Rect struct {
	x1, y1, x2, y2 int
}

func (r Rect) Area() int {
	return (r.x2 - r.x1 + 1) * (r.y2 - r.y1 + 1)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	corners := [][2]int{}

	for _, line := range input {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		corners = append(corners, [2]int{x, y})
	}

	// Generate all pairs of corners, sort so x and y are right way round
	var pairs []Rect
	n := len(corners)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := corners[i][0], corners[i][1]
			c, d := corners[j][0], corners[j][1]

			rect := Rect{
				x1: utils.Min(a, c),
				y1: utils.Min(b, d),
				x2: utils.Max(a, c),
				y2: utils.Max(b, d),
			}
			pairs = append(pairs, rect)
		}
	}

	// Sort pairs by area descending
	sort.Slice(pairs, func(i, j int) bool {
		areaI := pairs[i].Area()
		areaJ := pairs[j].Area()
		return areaI > areaJ
	})

	// Generate lines (consecutive pairs of corners including wrap-around)
	var lines []Rect
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		p, q := corners[i][0], corners[i][1]
		r, s := corners[j][0], corners[j][1]

		line := Rect{
			x1: utils.Min(p, r),
			y1: utils.Min(q, s),
			x2: utils.Max(p, r),
			y2: utils.Max(q, s),
		}
		lines = append(lines, line)
	}

	// Find the largest rectangle that doesn't intersect any line
	for _, rect := range pairs {

		intersects := false
		for _, line := range lines {
			// Check if line intersects rectangle
			if line.x1 < rect.x2 && line.y1 < rect.y2 && line.x2 > rect.x1 && line.y2 > rect.y1 {
				intersects = true
				break
			}
		}

		if !intersects {
			fmt.Printf("Part 2: %d\n", rect.Area())
			return
		}
	}
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
