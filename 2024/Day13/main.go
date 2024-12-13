package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Puzzle struct {
	Ax int
	Ay int
	Bx int
	By int
	X  int
	Y  int
}

func part1(input []string) {
	defer utils.Timer("part1")()

	puzzles := []Puzzle{}

	for i := 0; i < len(input); i = i + 4 {
		Ax := utils.IntegerOf(strings.TrimSuffix(strings.Split(strings.Fields(input[i])[2], "+")[1], ","))
		Ay := utils.IntegerOf(strings.Split(strings.Fields(input[i])[3], "+")[1])
		Bx := utils.IntegerOf(strings.TrimSuffix(strings.Split(strings.Fields(input[i+1])[2], "+")[1], ","))
		By := utils.IntegerOf(strings.Split(strings.Fields(input[i+1])[3], "+")[1])
		X := utils.IntegerOf(strings.TrimSuffix(strings.Split(strings.Fields(input[i+2])[1], "=")[1], ","))
		Y := utils.IntegerOf(strings.Split(strings.Fields(input[i+2])[2], "=")[1])
		puzzles = append(puzzles, Puzzle{Ax, Ay, Bx, By, X, Y})
	}

	total := 0

	for _, p := range puzzles {
		divisor := (p.Ax*p.By - p.Ay*p.Bx)
		A := float64(p.X*p.By-p.Y*p.Bx) / float64(divisor)
		B := float64(p.Ax*p.Y-p.Ay*p.X) / float64(divisor)

		if A >= 0 && B >= 0 {
			if A <= 100 && B <= 100 {
				if A == float64(int(A)) && B == float64(int(B)) {
					total += int(A*3 + B)
				}
			}
		}
	}
	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()

	puzzles := []Puzzle{}

	for i := 0; i < len(input); i = i + 4 {
		Ax := utils.IntegerOf(strings.TrimSuffix(strings.Split(strings.Fields(input[i])[2], "+")[1], ","))
		Ay := utils.IntegerOf(strings.Split(strings.Fields(input[i])[3], "+")[1])
		Bx := utils.IntegerOf(strings.TrimSuffix(strings.Split(strings.Fields(input[i+1])[2], "+")[1], ","))
		By := utils.IntegerOf(strings.Split(strings.Fields(input[i+1])[3], "+")[1])
		X := utils.IntegerOf(strings.TrimSuffix(strings.Split(strings.Fields(input[i+2])[1], "=")[1], ",")) + 10000000000000
		Y := utils.IntegerOf(strings.Split(strings.Fields(input[i+2])[2], "=")[1]) + 10000000000000
		puzzles = append(puzzles, Puzzle{Ax, Ay, Bx, By, X, Y})
	}

	total := 0

	for _, p := range puzzles {
		divisor := (p.Ax*p.By - p.Ay*p.Bx)
		A := float64(p.X*p.By-p.Y*p.Bx) / float64(divisor)
		B := float64(p.Ax*p.Y-p.Ay*p.X) / float64(divisor)

		if A >= 0 && B >= 0 {
			if A == float64(int(A)) && B == float64(int(B)) {
				total += int(A*3 + B)
			}
		}
	}
	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
