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
