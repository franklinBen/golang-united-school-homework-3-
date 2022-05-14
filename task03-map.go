package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	for m := range input {
		keys = append(keys, m)
	}
	sort.Ints(keys)
	for _, m := range keys {
		result[m] = input[m]
	}
	return result
}
