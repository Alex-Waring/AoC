package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Alex-Waring/AoC/utils"
)

type Range struct {
	start    int
	end      int
	interval int
}

type SourceToDest struct {
	ranges []Range
}

func parse_maps(lines []string) SourceToDest {
	ranges := []Range{}
	for _, line := range utils.RemoveSliceSpaces(lines) {
		values := strings.Fields(line)
		destination_start := utils.IntegerOf(values[0])
		source_start := utils.IntegerOf(values[1])
		map_range := utils.IntegerOf(values[2])

		ranges = append(ranges, Range{
			source_start, source_start + map_range, destination_start - source_start,
		})
	}
	return SourceToDest{ranges: ranges}
}

func lookup_map(input SourceToDest, lookup int) int {
	for _, Range := range input.ranges {
		if lookup >= Range.start && lookup <= Range.end {
			return lookup + Range.interval
		}
	}
	return lookup
}

func main() {
	defer utils.Timer("main")()

	raw_input, _ := os.ReadFile("input.txt")

	sections := strings.Split(string(raw_input), "\n\n")

	seeds := strings.Fields(strings.Split(sections[0], ":")[1])
	seed_to_soil := parse_maps(strings.Split(strings.Split(sections[1], ":\n")[1], "\n"))
	soil_to_fertilizer := parse_maps(strings.Split(strings.Split(sections[2], ":\n")[1], "\n"))
	fertilizer_to_water := parse_maps(strings.Split(strings.Split(sections[3], ":\n")[1], "\n"))
	water_to_light := parse_maps(strings.Split(strings.Split(sections[4], ":\n")[1], "\n"))
	light_to_temperature := parse_maps(strings.Split(strings.Split(sections[5], ":\n")[1], "\n"))
	temperature_to_humidity := parse_maps(strings.Split(strings.Split(sections[6], ":\n")[1], "\n"))
	humidity_to_location := parse_maps(strings.Split(strings.Split(sections[7], ":\n")[1], "\n"))

	locations := []int{}
	for _, seed := range seeds {
		soil := lookup_map(seed_to_soil, utils.IntegerOf(seed))
		fertilizer := lookup_map(soil_to_fertilizer, soil)
		water := lookup_map(fertilizer_to_water, fertilizer)
		light := lookup_map(water_to_light, water)
		temp := lookup_map(light_to_temperature, light)
		humidity := lookup_map(temperature_to_humidity, temp)
		locations = append(locations, lookup_map(humidity_to_location, humidity))
	}
	fmt.Println(slices.Min(locations))

	new_locations := []int{}

	// It's not optimised but it works...
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println(seeds[i*2])
		start := utils.IntegerOf(seeds[i])
		seed_range := utils.IntegerOf(seeds[i+1])
		for seed := start; seed < start+seed_range; seed++ {
			soil := lookup_map(seed_to_soil, seed)
			fertilizer := lookup_map(soil_to_fertilizer, soil)
			water := lookup_map(fertilizer_to_water, fertilizer)
			light := lookup_map(water_to_light, water)
			temp := lookup_map(light_to_temperature, light)
			humidity := lookup_map(temperature_to_humidity, temp)
			new_location := lookup_map(humidity_to_location, humidity)
			new_locations = append(new_locations, new_location)
		}
	}
	fmt.Println(slices.Min(new_locations))
}
