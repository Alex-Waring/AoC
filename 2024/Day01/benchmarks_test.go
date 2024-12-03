package main

import (
	"testing"

	"github.com/Alex-Waring/AoC/utils"
)

func BenchmarkPart1(b *testing.B) {
	input := utils.ReadInput("input.txt")
	part1(input)
}

func BenchmarkPart2(b *testing.B) {
	input := utils.ReadInput("input.txt")
	part2(input)
}
