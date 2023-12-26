package main

import (
	"fmt"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Blueprint struct {
	number        int
	ore_cost      int
	clay_cost     int
	obsidian_cost [2]int // Ore, Clay
	geode_cost    [2]int // Ore, obsidian
}

type cache struct {
	minute         int
	ore            int
	clay           int
	obsidian       int
	geodes         int
	ore_robot      int
	clay_robot     int
	obsidian_robot int
	geode_robot    int
}

func advanceCache(c cache) cache {
	return_cache := cache{
		minute:         c.minute + 1,
		ore:            c.ore + c.ore_robot,
		clay:           c.clay + c.clay_robot,
		obsidian:       c.obsidian + c.obsidian_robot,
		geodes:         c.geodes + c.geode_robot,
		ore_robot:      c.ore_robot,
		clay_robot:     c.clay_robot,
		obsidian_robot: c.obsidian_robot,
		geode_robot:    c.geode_robot,
	}
	return return_cache
}

func remainingBuilt(timeLeft int, robots int, geodes int) int {
	result := geodes
	for i := timeLeft; i > 0; i-- {
		result += robots
		robots++
	}
	return result
}

func part1(blueprints map[int]Blueprint, minutes int) {
	defer utils.Timer("part1")()
	results := map[int]int{}

	for _, blueprint := range blueprints {
		q := utils.Queue[cache]{}
		q.Push(cache{
			minute:    0,
			ore_robot: 1,
		})

		// If we can only spend x ore per minute, dont but x+1 robots, here
		// we find the max per min
		max_ore_min := max(blueprint.ore_cost, blueprint.clay_cost, blueprint.obsidian_cost[0], blueprint.geode_cost[0])
		max_clay_min := blueprint.obsidian_cost[1]
		max_obs_min := blueprint.geode_cost[1]

		seen := map[cache]bool{}

		for !q.IsEmpty() {
			itter := q.Pop()
			// If we've reached targer minutes, see if we have more geodes and add
			if itter.minute >= minutes {
				if geodes, exists := results[blueprint.number]; exists {
					if geodes < itter.geodes {
						results[blueprint.number] = itter.geodes
					}
				} else {
					results[blueprint.number] = itter.geodes
				}
				continue
			}

			// If the time left * geode robots doesn't make number bigger than
			// current max, continue
			time_left := minutes - itter.minute
			if itter.geodes+remainingBuilt(time_left, itter.geode_robot, itter.geodes) <= results[blueprint.number] {
				continue
			}

			// If we've been at this point, continue
			if seen[itter] {
				continue
			} else {
				seen[itter] = true
			}

			// Branch on trying to build robots
			if itter.ore >= blueprint.ore_cost && itter.ore_robot < max_ore_min {
				new_cache := advanceCache(itter)
				new_cache.ore_robot++
				new_cache.ore -= blueprint.ore_cost
				q.Push(new_cache)
			}
			if itter.ore >= blueprint.clay_cost && itter.clay_robot < max_clay_min {
				new_cache := advanceCache(itter)
				new_cache.clay_robot++
				new_cache.ore -= blueprint.clay_cost
				q.Push(new_cache)
			}
			// Don't build obsidian if we have no clay robots
			if itter.clay_robot != 0 && itter.obsidian_robot < max_obs_min {
				if itter.ore >= blueprint.obsidian_cost[0] && itter.clay >= blueprint.obsidian_cost[1] {
					new_cache := advanceCache(itter)
					new_cache.obsidian_robot++
					new_cache.ore -= blueprint.obsidian_cost[0]
					new_cache.clay -= blueprint.obsidian_cost[1]
					q.Push(new_cache)
				}
			}
			if itter.ore >= blueprint.geode_cost[0] && itter.obsidian >= blueprint.geode_cost[1] {
				new_cache := advanceCache(itter)
				new_cache.geode_robot++
				new_cache.ore -= blueprint.geode_cost[0]
				new_cache.obsidian -= blueprint.geode_cost[1]
				q.Push(new_cache)
			}

			// We also have to push buying nothing
			new_cache := advanceCache(itter)
			q.Push(new_cache)
			// fmt.Println(results)
		}
	}

	result := 0
	for id, geodes := range results {
		result += id * geodes
	}
	fmt.Println(result)
}

func part2(blueprints map[int]Blueprint, minutes int) {
	defer utils.Timer("part2")()
	results := map[int]int{}

	for _, blueprint := range blueprints {
		q := utils.Queue[cache]{}
		q.Push(cache{
			minute:    0,
			ore_robot: 1,
		})

		// If we can only spend x ore per minute, dont but x+1 robots, here
		// we find the max per min
		max_ore_min := max(blueprint.ore_cost, blueprint.clay_cost, blueprint.obsidian_cost[0], blueprint.geode_cost[0])
		max_clay_min := blueprint.obsidian_cost[1]
		max_obs_min := blueprint.geode_cost[1]

		seen := map[cache]bool{}

		for !q.IsEmpty() {
			itter := q.Pop()
			// If we've reached targer minutes, see if we have more geodes and add
			if itter.minute >= minutes {
				if geodes, exists := results[blueprint.number]; exists {
					if geodes < itter.geodes {
						results[blueprint.number] = itter.geodes
					}
				} else {
					results[blueprint.number] = itter.geodes
				}
				continue
			}

			// If the time left * geode robots doesn't make number bigger than
			// current max, continue
			time_left := minutes - itter.minute
			if itter.geodes+remainingBuilt(time_left, itter.geode_robot, itter.geodes) <= results[blueprint.number] {
				continue
			}

			// If we've been at this point, continue
			if seen[itter] {
				continue
			} else {
				seen[itter] = true
			}

			// Branch on trying to build robots
			if itter.ore >= blueprint.ore_cost && itter.ore_robot < max_ore_min {
				new_cache := advanceCache(itter)
				new_cache.ore_robot++
				new_cache.ore -= blueprint.ore_cost
				q.Push(new_cache)
			}
			if itter.ore >= blueprint.clay_cost && itter.clay_robot < max_clay_min {
				new_cache := advanceCache(itter)
				new_cache.clay_robot++
				new_cache.ore -= blueprint.clay_cost
				q.Push(new_cache)
			}
			// Don't build obsidian if we have no clay robots
			if itter.clay_robot != 0 && itter.obsidian_robot < max_obs_min {
				if itter.ore >= blueprint.obsidian_cost[0] && itter.clay >= blueprint.obsidian_cost[1] {
					new_cache := advanceCache(itter)
					new_cache.obsidian_robot++
					new_cache.ore -= blueprint.obsidian_cost[0]
					new_cache.clay -= blueprint.obsidian_cost[1]
					q.Push(new_cache)
				}
			}
			if itter.ore >= blueprint.geode_cost[0] && itter.obsidian >= blueprint.geode_cost[1] {
				new_cache := advanceCache(itter)
				new_cache.geode_robot++
				new_cache.ore -= blueprint.geode_cost[0]
				new_cache.obsidian -= blueprint.geode_cost[1]
				q.Push(new_cache)
			}

			// We also have to push buying nothing
			new_cache := advanceCache(itter)
			q.Push(new_cache)
			// fmt.Println(results)
		}
	}

	result := 1
	for _, geodes := range results {
		result *= geodes
	}
	fmt.Println(result)
}

func main() {
	input := utils.ReadInput("input.txt")
	blueprints := map[int]Blueprint{}

	for index, line := range input {
		number := index + 1
		costs := strings.Split(strings.Split(line, ":")[1], ".")
		ore_cost := utils.IntegerOf(strings.Split(costs[0], " ")[5])
		clay_cost := utils.IntegerOf(strings.Split(costs[1], " ")[5])
		obsidian_cost := [2]int{utils.IntegerOf(strings.Split(costs[2], " ")[5]), utils.IntegerOf(strings.Split(costs[2], " ")[8])}
		geode_cost := [2]int{utils.IntegerOf(strings.Split(costs[3], " ")[5]), utils.IntegerOf(strings.Split(costs[3], " ")[8])}

		blueprints[number] = Blueprint{
			number:        number,
			ore_cost:      ore_cost,
			clay_cost:     clay_cost,
			obsidian_cost: obsidian_cost,
			geode_cost:    geode_cost,
		}
	}

	part1(blueprints, 24)
	part2_blueprints := map[int]Blueprint{
		1: blueprints[1],
		2: blueprints[2],
		3: blueprints[3],
	}
	part2(part2_blueprints, 32)
}
