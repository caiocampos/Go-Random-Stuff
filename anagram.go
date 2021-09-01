package anagram

import (
	"strings"
	"unicode"
)

// Detect function seeks for the anagrams of a word in a list
func Detect(subject string, candidates []string) []string {
	res := []string{}
	countSubj := make(map[rune]int8)
	for _, letter := range subject {
		letter = unicode.ToLower(letter)
		if frequency, exist := countSubj[letter]; exist {
			countSubj[letter] = frequency + 1
		} else {
			countSubj[letter] = 1
		}
	}
	countCand := make(map[rune]int8)
	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}
		for k := range countSubj {
			countCand[k] = 0
		}
		ok := true
		for _, letter := range candidate {
			letter = unicode.ToLower(letter)
			if _, exist := countSubj[letter]; !exist {
				ok = false
				continue
			}
			if countCand[letter]++; countCand[letter] > countSubj[letter] {
				ok = false
				continue
			}
		}
		if ok {
			for letter := range countCand {
				if countCand[letter] != countSubj[letter] {
					ok = false
				}
			}
		}
		if ok {
			res = append(res, candidate)
		}
	}
	return res
}
