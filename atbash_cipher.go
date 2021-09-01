package atbash

import (
	"strings"
	"unicode"
)

// Atbash function cipher the input with the atbash method of encryption
func Atbash(input string) string {
	b := strings.Builder{}
	counter := 0
	for _, el := range input {
		if counter == 5 {
			b.WriteRune(' ')
			counter = 0
		}
		if unicode.IsNumber(el) {
			b.WriteRune(el)
			counter++
		} else if unicode.IsLetter(el) {
			el = unicode.ToLower(el)
			el = 'z' - (el - 'a')
			b.WriteRune(el)
			counter++
		}
	}
	return strings.TrimRightFunc(b.String(), unicode.IsSpace)
}
