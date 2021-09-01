package luhn

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Valid function verifies if the input is valid per the Luhn formula
func Valid(input string) bool {
	return ValidVD(input)
}

// ValidVA function verifies if the input is valid per the Luhn formula (Version A)
func ValidVA(input string) bool {
	vals := make([]int, 0, len(input))
	for _, el := range input {
		if unicode.IsNumber(el) {
			vals = append(vals, int(el-'0')) // int(el-'0') converts rune to int
		} else {
			if el != ' ' {
				return false
			}
		}
	}
	if len(vals) < 2 {
		return false
	}
	double := len(vals)&1 == 0 // begin doubling if len(vals) is even
	sum := 0
	for _, val := range vals {
		if double {
			if val > 4 { // if val > 4 then val * 2 > 9
				sum += (val << 1) - 9 // val * 2 - 9
			} else {
				sum += val << 1 // val * 2
			}
		} else {
			sum += val
		}
		double = !double
	}
	return sum%10 == 0
}

// ValidVB function verifies if the input is valid per the Luhn formula (Version B)
func ValidVB(input string) bool {
	count, sumEven, sumOdd, val := 0, 0, 0, 0
	even := true
	for _, el := range input {
		if unicode.IsNumber(el) {
			val = int(el - '0') // int(el-'0') converts rune to int
			sumEven += calculates(val, even)
			sumOdd += calculates(val, !even)
			even = !even
			count++
		} else {
			if el != ' ' {
				return false
			}
		}
	}
	if count < 2 {
		return false
	}
	if count&1 == 0 { // uses sumEven if count is even
		return sumEven%10 == 0
	}
	return sumOdd%10 == 0
}

func calculates(n int, double bool) int {
	if double {
		if n > 4 { // if val > 4 then val * 2 > 9
			return (n << 1) - 9 // val * 2 - 9
		}
		return n << 1 // val * 2
	}
	return n
}

var regexSpaces = regexp.MustCompile(`\s`)
var regexNotDigit = regexp.MustCompile(`\D`)

// ValidVC function verifies if the input is valid per the Luhn formula (Version C)
func ValidVC(input string) bool {
	vals := regexSpaces.ReplaceAllString(input, "")
	if len(vals) < 2 || regexNotDigit.MatchString(vals) {
		return false
	}
	double := len(vals)&1 == 0 // begin doubling if len(vals) is even
	sum, val := 0, 0
	for _, el := range vals {
		val = int(el - '0') // int(el-'0') converts rune to int
		if double {
			if val > 4 { // if val > 4 then val * 2 > 9
				sum += (val << 1) - 9 // val * 2 - 9
			} else {
				sum += val << 1 // val * 2
			}
		} else {
			sum += val
		}
		double = !double
	}
	return sum%10 == 0
}

// ValidVD function verifies if the input is valid per the Luhn formula (Version D for dummies)
func ValidVD(input string) bool {
	vals := strings.Replace(input, " ", "", -1)
	if len(vals) < 2 {
		return false
	}
	double := len(vals)%2 == 0 // begin doubling if len(vals) is even
	sum, val := 0, 0
	var err error
	for _, el := range vals {
		val, err = strconv.Atoi(string(el))
		if err != nil {
			return false
		}
		if double {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		sum += val
		double = !double
	}
	return sum%10 == 0
}
