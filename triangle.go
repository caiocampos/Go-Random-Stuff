// Package triangle provides ways to validate triangles
package triangle

import "math"

// Kind type defines the types of triangles
type Kind byte

const (
	// NaT - not a triangle
	NaT Kind = iota
	// Equ - equilateral
	Equ
	// Iso - isosceles
	Iso
	// Sca - scalene
	Sca
)

// KindFromSides function verifies the type of a triangle
func KindFromSides(a, b, c float64) Kind {
	if IsInvalid(a) || IsInvalid(b) || IsInvalid(c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}

// IsInvalid function tests if n is a valid side of a triangle
func IsInvalid(n float64) bool {
	return n <= 0 || math.IsNaN(n) || math.IsInf(n, 0)
}
