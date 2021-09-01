package isbn

// IsValidISBN function validates if the input is a valid ISBN-10
func IsValidISBN(input string) bool {
	if len(input) < 10 {
		return false
	}
	m, val, sum := 10, 0, 0
	for i, el := range input {
		if el == '-' {
			continue
		}
		if el == 'X' {
			val = 10
		} else {
			val = int(el - '0')
		}
		if m == 1 {
			if len(input) > i+1 {
				return false
			}
			return (sum+val)%11 == 0
		} else if val > 9 || val < 0 {
			return false
		} else {
			sum += val * m
			m--
		}
	}
	return false
}
