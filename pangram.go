package pangram

import "unicode"

// IsPangram function verify if a phrase is a pangram
func IsPangram(input string) bool {
	count := 0
	letters := make(map[rune]int8)
	for _, letter := range input {
		if !unicode.IsLetter(letter) {
			continue
		}
		uLetter := unicode.ToUpper(letter)
		if frequency, exist := letters[uLetter]; exist {
			letters[uLetter] = frequency + 1
		} else {
			letters[uLetter] = 1
			count++
		}
	}
	return count == 26
}
