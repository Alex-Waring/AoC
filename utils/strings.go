package utils

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
