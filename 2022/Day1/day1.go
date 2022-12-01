package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(array []string) int {  
	result := 0  
	for _, v := range array {  
		value, err := strconv.Atoi(v)
		check(err)
		result += value
	}  
	return result  
}

func findmax(array []int) int {
	max := 0
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return max
}

func findmax3sum(array []int) int {
	max := [3]int{0, 0, 0}
	for _, v := range array {
		if v > max[0] {
			max[0] = v
		} else if v > max[1] {
			max[1] = v
		} else if v > max[2] {
			max[2] = v
		}
	}
	result := 0
	for _, v := range max {
		result += v
	}
	return result
}

func getElves(input string) []int {
	elves := strings.Split(input, "\n\n")
	elvessum := []int{}

	for _, value := range elves {
		total := sum(strings.Split(value, "\n"))
		elvessum = append(elvessum, total)
	}

	return elvessum
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)
	output := getElves(string(input))
	check(err)
	fmt.Println(findmax(output))
	fmt.Println(findmax3sum(output))
}