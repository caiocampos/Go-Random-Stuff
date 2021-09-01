// Package acronym provides ways to abbreviate phrases
package acronym

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate function abbreviates a phrase
func Abbreviate(s string) string {
	var b strings.Builder
	for _, word := range regexp.MustCompile("[\\s-]+").Split(s, -1) {
		letter, err := GetFistLetter(word)
		if err == nil {
			b.WriteRune(unicode.ToUpper(letter))
		}
	}
	return b.String()
}

// GetFistLetter function tries to search the first letter on a word
func GetFistLetter(word string) (rune, error) {
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			return letter, nil
		}
	}
	return '.', errors.New("The input has no letters")
}
