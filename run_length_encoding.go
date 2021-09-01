package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode function implements the Run-length encoding
func RunLengthEncode(input string) string {
	var current rune
	var b strings.Builder
	count := 0
	reset := func(r rune) {
		count, current = 1, r
	}
	write := func() {
		switch {
		case count == 0:
			return
		case count > 1:
			b.WriteString(strconv.Itoa(count))
			fallthrough
		default:
			b.WriteRune(current)
		}
	}
	for _, letter := range input {
		if letter != current {
			write()
			reset(letter)
		} else {
			count++
		}
	}
	write()
	return b.String()
}

// RunLengthDecode function implements the Run-length decoding
func RunLengthDecode(input string) string {
	var bRes strings.Builder
	var bn strings.Builder
	for _, letter := range input {
		if unicode.IsNumber(letter) {
			bn.WriteRune(letter)
		} else {
			n := 1
			if bn.Len() > 0 {
				n, _ = strconv.Atoi(bn.String())
				bn.Reset()
			}
			bRes.WriteString(strings.Repeat(string(letter), n))
			/*
				for i := 0; i < n; i++ {
					bRes.WriteRune(letter)
				}
			*/
		}
	}
	return bRes.String()
}
