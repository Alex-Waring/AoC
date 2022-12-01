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
