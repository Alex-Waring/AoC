package main

import (
	"fmt"

	"github.com/Alex-Waring/AoC/utils"
)

func part1(input []string) {
	defer utils.Timer("part1")()
	data := input[0]

	diskData := []string{}
	id := 0

	for i := 0; i < len(data); i++ {
		length := utils.IntegerOf(string(data[i]))
		if i%2 == 0 {
			for j := 0; j < length; j++ {
				diskData = append(diskData, fmt.Sprintf("%d", id))
			}
			id++
		} else {
			for j := 0; j < length; j++ {
				diskData = append(diskData, ".")
			}
		}
	}

	compressedData := []string{}

compression_loop:
	for {
		// If we start with a number, add that list of numbers to the output
		if diskData[0] != "." {
			compressedData = append(compressedData, diskData[0])
			diskData = diskData[1:]
		} else {
			// Otherwise take the end number
		suffix_loop:
			for {
				if diskData[len(diskData)-1] == "." {
					diskData = diskData[:len(diskData)-1]
				} else {
					break suffix_loop
				}
			}
			compressedData = append(compressedData, diskData[len(diskData)-1])
			diskData = diskData[1 : len(diskData)-1]
		}
		if len(diskData) == 0 {
			break compression_loop
		}
	}

	total := 0
	for i := 0; i < len(compressedData); i++ {
		total += i * utils.IntegerOf(compressedData[i])
	}
	fmt.Println(total)
}

func part2(input []string) {
	defer utils.Timer("part2")()
	data := input[0]

	type File struct {
		size  int
		ID    int
		blank bool
	}
	diskData := []File{}
	id := 0

	for i := 0; i < len(data); i++ {
		length := utils.IntegerOf(string(data[i]))
		if i%2 == 0 {
			diskData = append(diskData, File{
				size:  length,
				ID:    id,
				blank: false,
			})
			id++
		} else {
			diskData = append(diskData, File{
				size:  length,
				blank: true,
			})
		}
	}

	id--
	for {
		// Find the file with ID
		var loc int
		var file_to_move File
		for i, file := range diskData {
			if file.ID == id {
				loc = i
				file_to_move = file
			}
		}
		if file_to_move.size == 0 {
			break
		}
		// Move up through the compressed data and try and place it
		for i, file := range diskData {
			if !file.blank {
				continue
			}
			if !(file.size >= file_to_move.size) {
				continue
			}
			// If we've passed the file, move on
			if i >= loc {
				break
			}
			// Construct the new list which is:
			// start up to space
			// moved file
			// empty space remaining
			// list up until old moved file loc
			// empty space for old file
			// list past old file loc
			new_data := make([]File, 0)
			for j := 0; j < i; j++ {
				new_data = append(new_data, diskData[j])
			}
			new_data = append(new_data, file_to_move)
			if file_to_move.size < file.size {
				new_data = append(new_data, File{
					size:  file.size - file_to_move.size,
					blank: true,
				})
			}
			new_data = append(new_data, diskData[i+1:loc]...)
			new_data = append(new_data, File{
				size:  file_to_move.size,
				blank: true,
			})
			new_data = append(new_data, diskData[loc+1:]...)
			diskData = new_data
			break
		}
		if id == 0 {
			break
		} else {
			id--
		}
	}

	compressed_data := []string{}

	for _, file := range diskData {
		for i := 0; i < file.size; i++ {
			if file.blank {
				compressed_data = append(compressed_data, ".")
			} else {
				compressed_data = append(compressed_data, fmt.Sprintf("%d", file.ID))
			}
		}
	}
	total := 0
	for i := 0; i < len(compressed_data); i++ {
		if compressed_data[i] == "." {
			continue
		}
		total += i * utils.IntegerOf(compressed_data[i])
	}
	fmt.Println(total)
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}
