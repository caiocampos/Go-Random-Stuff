package diffsquares

// SquareOfSum calculates the square of the sums
func SquareOfSum(n int) int {
	return SquareOfSumVB(n)
}

// SquareOfSumVA calculates the square of the sums (Version A)
func SquareOfSumVA(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SquareOfSumVB calculates the square of the sums (Version B)
func SquareOfSumVB(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares
func SumOfSquares(n int) int {
	return SumOfSquaresVB(n)
}

// SumOfSquaresVA calculates the sum of the squares (Version A)
func SumOfSquaresVA(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// SumOfSquaresVB calculates the sum of the squares (Version B)
func SumOfSquaresVB(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference of square of the sums and sum of the squares
func Difference(n int) int {
	return DifferenceVB(n)
}

// DifferenceVA calculates the difference of square of the sums and sum of the squares (Version A)
func DifferenceVA(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// DifferenceVB calculates the difference of square of the sums and sum of the squares (Version B)
func DifferenceVB(n int) int {
	return n * (n + 1) * (3*n*(n+1) - 2*(2*n+1)) / 12
}
