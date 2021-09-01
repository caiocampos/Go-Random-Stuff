package protein

import (
	"errors"
	"strings"
)

var (
	// ErrStop - STOP Codon
	ErrStop = errors.New("STOP Codon")
	// ErrInvalidBase - Invalid Codon
	ErrInvalidBase = errors.New("Invalid Codon")
)

// FromCodon function converts a codon in a protein
func FromCodon(codon string) (string, error) {
	switch codon {
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "AUG":
		return "Methionine", nil
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA function converts a RNA in a list of proteins
func FromRNA(rna string) ([]string, error) {
	var res []string
	var b strings.Builder
	c := 0
	for _, nucleotide := range rna {
		b.WriteRune(nucleotide)
		if c++; c == 3 {
			protein, err := FromCodon(b.String())
			if err == nil {
				res = append(res, protein)
			} else if err == ErrStop {
				return res, nil
			} else {
				return res, ErrInvalidBase
			}
			b.Reset()
			c = 0
		}
	}
	return res, nil
}
