package collatzconjecture

import "errors"

// CollatzConjecture function calculates how much steps is required to reach 1
func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n <= 0 {
		return 0, errors.New("Value must be greather than zero")
	} else if n != 1 {
		for n != 1 {
			if n|1 == n { // n % 2 == 0
				n = 3*n + 1
			} else {
				n = n >> 1 // n / 2
			}
			steps++
		}
	}
	return steps, nil
}
