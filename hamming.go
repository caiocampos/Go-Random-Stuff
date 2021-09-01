package hamming

import "errors"

// Distance function calculates the Hamming difference between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be of equal length")
	}
	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
