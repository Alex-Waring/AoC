package utils

import "sort"

func Median(slice []int) int {
	dataCopy := make([]int, len(slice))
	copy(dataCopy, slice)

	sort.Ints(dataCopy)
	l := len(dataCopy)
	if l%2 == 0 {
		return dataCopy[l/2]
	} else {
		return dataCopy[l/2+1]
	}
}

func Mean(slice []int) float64 {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return float64(sum) / float64(len(slice))
}
