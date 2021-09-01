package scrabble

import "unicode"

// Score function computes the scrabble score for a word
func Score(word string) int {
	return ScoreVB(word)
}

// ScoreVA function computes the scrabble score for a word (Version A)
func ScoreVA(word string) int {
	res := 0
	for _, letter := range word {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			res++
		case 'F', 'H', 'V', 'W', 'Y':
			res += 4
		case 'B', 'C', 'M', 'P':
			res += 3
		case 'D', 'G':
			res += 2
		case 'J', 'X':
			res += 8
		case 'Q', 'Z':
			res += 10
		default:
			res += 5
		}
	}
	return res
}

var scores = map[rune]int {
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// ScoreVB function computes the scrabble score for a word (Version B)
func ScoreVB(word string) int {
	res := 0
	for _, letter := range word {
		res += scores[unicode.ToUpper(letter)]
	}
	return res
}
