package utils

import (
	"strconv"
	"strings"
)

func Sum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

// This function swallows the error return of strconv Atoi so it
// can be used inline. This is bad but faster to type
func IntegerOf(str string) int {
	num, err := strconv.Atoi(str)
	Check(err)
	return num
}

// Makes a slice from a to b
func MakeRange(start int, end int) []int {
	var max int
	var min int
	if start > end {
		max = start
		min = end
	} else {
		max = end
		min = start
	}

	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func Multiply(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = res * arr[i]
	}
	return res
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Abs returns the absolute value.
func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMax(arr []int) int {
	// Initialize the variables to hold the maximum and minimum values to draw comparisons.
	max := arr[0]
	// Iterate over the array
	for i := 1; i < len(arr); i++ {
		// if the current element is greater than the present maximum
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func FindMin(arr []int) int {
	// Initialize the variables to hold the maximum and minimum values to draw comparisons.
	min := arr[0]
	// Iterate over the array
	for i := 1; i < len(arr); i++ {
		// if the current element is smaller than the present minimum
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}

// Returns true if a ends with b
func EndsWith(a, b int) bool {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	return strings.HasSuffix(aStr, bStr)
}

// trims the int b off of a
func RemoveSuffix(a, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	return IntegerOf(strings.TrimSuffix(aStr, bStr))
}

func IsDivisible(a, b int) bool {
	if b == 0 {
		return false // Avoid division by zero
	}
	return a%b == 0
}
