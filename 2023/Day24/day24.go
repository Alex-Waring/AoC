package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
	"github.com/davidkleiven/gononlin/nonlin"
)

type Position struct {
	x int
	y int
	z int
}

type Velocity struct {
	x int
	y int
	z int
}

type Hailstone struct {
	pos Position
	vel Velocity
	m   float64
	c   float64
}

func part1(hail []Hailstone, lower_bound int, upper_bound int) {
	// Given we have y=mx+c, and we know an intersection happens when y is the same
	// m1 * x + c1 = m2 * x + c2 solves to
	// x = (c2-c1)/(m1-m2)
	// y = m1*(c2-c1)/(m1-m2) + c1
	total := 0
	for index, hailstone := range hail {
		for i := index + 1; i < len(hail); i++ {
			hailstone2 := hail[i]

			x := (hailstone2.c - hailstone.c) / (hailstone.m - hailstone2.m)
			y := hailstone.m*x + hailstone.c

			// to figure out if it's in the past, grab the x value and see if the diff matches the
			// vel sign for both hailstones
			if hailstone.vel.x < 0 {
				if x > float64(hailstone.pos.x) {
					// x velocity is -ve but we have a +ve x dir, we've gone back in time
					continue
				}
			} else {
				if x < float64(hailstone.pos.x) {
					// x velocity is +ve but we have a -ve x dir, we've gone back in time
					continue
				}
			}
			if hailstone2.vel.x < 0 {
				if x > float64(hailstone2.pos.x) {
					// x velocity is -ve but we have a +ve x dir, we've gone back in time
					continue
				}
			} else {
				if x < float64(hailstone2.pos.x) {
					// x velocity is +ve but we have a -ve x dir, we've gone back in time
					continue
				}
			}

			if x >= float64(lower_bound) && x <= float64(upper_bound) && y >= float64(lower_bound) && y <= float64(upper_bound) {
				total++
			}
		}
	}
	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	hail := []Hailstone{}

	for _, line := range input {
		pos_raw := strings.Split(strings.Split(line, "@")[0], ",")
		vel_raw := strings.Split(strings.Split(line, "@")[1], ",")

		new_hail := Hailstone{
			pos: Position{
				x: utils.IntegerOf(strings.ReplaceAll(pos_raw[0], " ", "")),
				y: utils.IntegerOf(strings.ReplaceAll(pos_raw[1], " ", "")),
				z: utils.IntegerOf(strings.ReplaceAll(pos_raw[2], " ", "")),
			},
			vel: Velocity{
				x: utils.IntegerOf(strings.ReplaceAll(vel_raw[0], " ", "")),
				y: utils.IntegerOf(strings.ReplaceAll(vel_raw[1], " ", "")),
				z: utils.IntegerOf(strings.ReplaceAll(vel_raw[2], " ", "")),
			},
		}
		hail = append(hail, new_hail)
	}

	// for each hailstone, take a step forward in time to calculate y = mx + c
	for index := range hail {
		hailstone := hail[index]
		y1, x1 := float64(hailstone.pos.y), float64(hailstone.pos.x)
		y2, x2 := float64(hailstone.pos.y+hailstone.vel.y), float64(hailstone.pos.x+hailstone.vel.x)

		hail[index].m = (y2 - y1) / (x2 - x1)
		hail[index].c = y1 - hail[index].m*x1
	}

	part1(hail, 200000000000000, 400000000000000)

	// Part 2 using https://github.com/davidkleiven/gononlin

	// We have 9 variables: x, y, z, dy, dy, dz and t1, t2, t3 for the first three collisions
	//                      0  1  2  3   4   5      6   7   8

	// And this wonderful set of problems
	problem := nonlin.Problem{
		F: func(out, x []float64) {
			out[0] = x[0] + x[3]*x[6] - float64(hail[0].pos.x) - x[6]*float64(hail[0].vel.x)
			out[1] = x[1] + x[4]*x[6] - float64(hail[0].pos.y) - x[6]*float64(hail[0].vel.y)
			out[2] = x[2] + x[5]*x[6] - float64(hail[0].pos.z) - x[6]*float64(hail[0].vel.z)
			out[3] = x[0] + x[3]*x[7] - float64(hail[1].pos.x) - x[7]*float64(hail[1].vel.x)
			out[4] = x[1] + x[4]*x[7] - float64(hail[1].pos.y) - x[7]*float64(hail[1].vel.y)
			out[5] = x[2] + x[5]*x[7] - float64(hail[1].pos.z) - x[7]*float64(hail[1].vel.z)
			out[6] = x[0] + x[3]*x[8] - float64(hail[2].pos.x) - x[8]*float64(hail[2].vel.x)
			out[7] = x[1] + x[4]*x[8] - float64(hail[2].pos.y) - x[8]*float64(hail[2].vel.y)
			out[8] = x[2] + x[5]*x[8] - float64(hail[2].pos.z) - x[8]*float64(hail[2].vel.z)
		},
	}

	solver := nonlin.NewtonKrylov{
		// Maximum number of Newton iterations
		Maxiter: 400000000000000000,

		// Stepsize used to appriximate jacobian with finite differences
		StepSize: 50000000000,

		// Tolerance for the solution
		Tol: 1,
	}
	// Have a guess, takes some iterations and eyeballing
	x0 := []float64{400000000000000.0, 300000000000000.0, 300000000000000.0, -300.0, 0, 0, 500000000000, 700000000000, 200000000000}
	res := solver.Solve(problem, x0)
	fmt.Println(res.X)
	// This chucks out a solution that needs rounding
	fmt.Println(res.X[0] + res.X[1] + res.X[2])
}
