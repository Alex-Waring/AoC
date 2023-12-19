package utils

import "strings"

func RemoveSliceSpaces(list []string) []string {
	return_list := []string{}

	for _, item := range list {
		if item != "" && item != " " {
			return_list = append(return_list, item)
		}
	}
	return return_list
}

func InsertIntoList(a []rune, index int, value rune) []rune {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func SliceFilledWithRune(size int, r rune) []rune {
	data := make([]rune, size)
	for i := 0; i < size; i++ {
		data[i] = r
	}
	return data
}

func SumListString(input []string) int {
	sum := 0
	for _, str := range input {
		sum += IntegerOf(str)
	}
	return sum
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	e += s + e - 1
	return str[s:e]
}
