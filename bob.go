package bob

import (
	"strings"
	"unicode"
)

// Hey function responses based on the input
func Hey(remark string) string {
	upper, lower, alphanum := false, false, false
	for _, letter := range remark {
		if !upper && unicode.IsUpper(letter) {
			upper, alphanum = true, true
			if lower {
				break
			}
			continue
		}
		if !lower && unicode.IsLower(letter) {
			lower, alphanum = true, true
			if upper {
				break
			}
			continue
		}
		if !alphanum && unicode.IsNumber(letter) {
			alphanum = true
			if upper && lower {
				break
			}
		}
	}
	if strings.HasSuffix(strings.TrimSpace(remark), "?") {
		if upper && !lower && alphanum {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if !lower {
		if upper {
			return "Whoa, chill out!"
		}
		if !alphanum {
			return "Fine. Be that way!"
		}
	}
	return "Whatever."
}
