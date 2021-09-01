package grains

import "errors"

const chessMin int = 1
const chessMax int = 64

// Square function gives the number of grains for the inputed number
func Square(n int) (uint64, error) {
	if n < chessMin || n > chessMax {
		return 0, errors.New("Square must be between 1 and 64")
	}
	return powTwo(n), nil
}

// Total function gives the number of grains for the 64 squares of the chessboard
func Total() uint64 {
	return powTwo(chessMax+1) - 1
}

func powTwo(n int) uint64 {
	return 1 << uint(n-1) // 1 << uint(n-1) == math.Pow(2, n)
}
