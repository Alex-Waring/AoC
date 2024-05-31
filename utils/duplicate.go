package utils

func DuplicateList[T any](list []T) []T {
	duplicate := make([]T, len(list))
	copy(duplicate, list)
	return duplicate
}
