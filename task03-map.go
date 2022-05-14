package homework

import "sort"

func sortMapValues(input map[int]string) []string {
	i := 0
	keys := make([]int, len(input))
	for k := range input {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	result := make([]string, len(input))
	for _, m := range keys {
		result[m] = input[m]
	}
	return result
}
