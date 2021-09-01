package wordcount

import (
	"regexp"
	"strings"
)

// Frequency type defines a alias for map[string]int
type Frequency map[string]int

// WordCount function calculates the frequency of an word in a frase
func WordCount(input string) Frequency {
	res := make(Frequency)
	for _, word := range regexp.MustCompile("(\\'*[^0-9A-Za-z']+\\'*)+").Split(strings.ToLower(input)+" ", -1) {
		if word == "" {
			continue
		}
		if frequency, exist := res[word]; exist {
			res[word] = frequency + 1
		} else {
			res[word] = 1
		}
	}
	return res
}
