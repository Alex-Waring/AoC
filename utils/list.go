package utils

func Remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func StringToIntList(slice []string) []int {
	var res []int
	for _, s := range slice {
		res = append(res, IntegerOf(s))
	}
	return res
}

func IntListContains(slice []int, value int) bool {
	for _, i := range slice {
		if i == value {
			return true
		}
	}
	return false
}

func StringListContains(slice []string, value string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}
