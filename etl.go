package etl

import "strings"

// Transform function transfroms a data in old format to a new format
func Transform(etl map[int][]string) map[string]int {
	res := make(map[string]int)
	for score, keyList := range etl {
		for _, key := range keyList {
			res[strings.ToLower(key)] = score
		}
	}
	return res
}
