package dups

func RemoveDuplicates(input []string) []string {
	unique := make(map[string]bool, len(input))
	result := make([]string, 0, len(input))

	for _, v := range input {
		if _, ok := unique[v]; !ok {
			result = append(result, v)
			unique[v] = true
		}
	}

	return result
}
