package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher function cipher the input with the rotational method of encryption
func RotationalCipher(input string, key int) string {
	if key == 0 || key == 26 {
		return input
	}
	var b strings.Builder
	b.Grow(len(input))
	for _, el := range input {
		if unicode.IsLetter(el) {
			up := el < 'a'
			el = el + rune(key)
			if (el > 'z') || (up && el > 'Z') {
				el -= 26
			}
		}
		b.WriteRune(el)
	}
	return b.String()
}
