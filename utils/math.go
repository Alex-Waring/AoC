package utils

import "strconv"

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

// Makes a slice from min to max, will panic of min is larger than
// max
func MakeRange(min int, max int) []int {
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
