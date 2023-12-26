package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

type entry struct {
	value int
	pos   int
}

func part1(data []entry, cll utils.LinkedList) {
	defer utils.Timer("part1")()
	var zero_e entry

	for _, num := range data {
		cll.Slide(num, num.value%(len(data)-1), len(data))
		if num.value == 0 {
			zero_e = num
		}
	}
	onek := utils.GetPosFrom(&cll, zero_e, 1000).(entry)
	twok := utils.GetPosFrom(&cll, zero_e, 2000).(entry)
	threek := utils.GetPosFrom(&cll, zero_e, 3000).(entry)
	fmt.Println(onek)
	fmt.Println(twok)
	fmt.Println(threek)

	fmt.Println(onek.value + twok.value + threek.value)
}

func part2(data []entry, cll utils.LinkedList, zero_e entry) {
	defer utils.Timer("part2")()

	for i := 0; i < 10; i++ {
		for _, num := range data {
			cll.Slide(num, num.value%(len(data)-1), len(data))
		}
	}

	onek := utils.GetPosFrom(&cll, zero_e, 1000).(entry)
	fmt.Println(onek)
	twok := utils.GetPosFrom(&cll, zero_e, 2000).(entry)
	fmt.Println(twok)
	threek := utils.GetPosFrom(&cll, zero_e, 3000).(entry)
	fmt.Println(threek)

	fmt.Println(onek.value + twok.value + threek.value)
}

func main() {
	input := utils.ReadInput("input.txt")

	data := []entry{}
	cll := utils.LinkedList{}

	for index, line := range input {
		num := utils.IntegerOf(line)
		e := entry{
			value: num,
			pos:   index,
		}
		data = append(data, e)
		cll.Insert(e)
	}
	utils.ConvertSinglyToCircular(&cll)
	part1(data, cll)

	data_p2 := []entry{}
	cll_p2 := utils.LinkedList{}
	for index, line := range input {
		num := utils.IntegerOf(line)
		e := entry{
			value: num * 811589153,
			pos:   index,
		}
		data_p2 = append(data_p2, e)
		cll_p2.Insert(e)
	}
	utils.ConvertSinglyToCircular(&cll_p2)
	part2(data_p2, cll_p2, entry{
		value: 0,
		pos:   2720,
	})
}
