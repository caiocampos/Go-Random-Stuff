package romannumerals

import (
	"errors"
	"strings"
)

type romanNum struct {
	symbol string
	value  int
}

var values = []romanNum{
	romanNum{"M", 1000},
	romanNum{"CM", 900},
	romanNum{"D", 500},
	romanNum{"CD", 400},
	romanNum{"C", 100},
	romanNum{"XC", 90},
	romanNum{"L", 50},
	romanNum{"XL", 40},
	romanNum{"X", 10},
	romanNum{"IX", 9},
	romanNum{"V", 5},
	romanNum{"IV", 4},
	romanNum{"I", 1},
}

// ToRomanNumeral converts an arabic number to the relative roman number
func ToRomanNumeral(arabic int) (string, error) {
	if arabic > 3000 || arabic < 1 {
		return "", errors.New("Out of bounds")
	}
	var b strings.Builder
	count := arabic
	for _, rn := range values {
		if count == 0 {
			break
		}
		for count%rn.value < count {
			b.WriteString(rn.symbol)
			count -= rn.value
		}
	}
	return b.String(), nil
}
