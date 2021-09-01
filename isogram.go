package isogram

import "unicode"

// IsIsogram function verify if a word is an isogram
func IsIsogram(word string) bool {
	return IsIsogramVA(word)
}

// IsIsogramVA function verify if a word is an isogram (Version A)
func IsIsogramVA(word string) bool {
	letters := make(map[rune]bool)
	for _, letter := range word {
		if letter == '-' || letter == ' ' {
			continue
		}
		uLetter := unicode.ToUpper(letter)
		if letters[uLetter] {
			return false
		}
		letters[uLetter] = true
	}
	return true
}
