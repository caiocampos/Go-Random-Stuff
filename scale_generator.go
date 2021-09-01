package scale

import (
	"strings"
	"unicode"
)

// Scale function generate the musical scale by the tonic and interval
func Scale(tonic string, interval string) []string {
	flat := UseFlat(tonic)
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}
	pos := GetPosTone(tonic)
	res := []string{}
	if pos == -1 {
		return res
	}
	for _, step := range interval {
		res = append(res, GetTone(pos, flat))
		switch step {
		case 'm':
			pos++
		case 'M':
			pos += 2
		case 'A':
			pos += 3
		}
		if pos >= len(tones) {
			pos %= len(tones)
		}
	}
	return res
}

// UseFlat function verifies if will use flat or sharp notation
func UseFlat(tonic string) bool {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return true
	}
	return false
}

type tone struct {
	major string
	sharp string
	flat  string
}

// A, A#, B, C, C#, D, D#, E, F, F#, G, G#
// A, Bb, B, C, Db, D, Eb, E, F, Gb, G, Ab
var tones = [12]tone{
	tone{"A", "", ""},
	tone{"", "A#", "Bb"},
	tone{"B", "", ""},
	tone{"C", "", ""},
	tone{"", "C#", "Db"},
	tone{"D", "", ""},
	tone{"", "D#", "Eb"},
	tone{"E", "", ""},
	tone{"F", "", ""},
	tone{"", "F#", "Gb"},
	tone{"G", "", ""},
	tone{"", "G#", "Ab"}}

// GetPosTone function searches for the position of the tonic
func GetPosTone(tonic string) int {
	var b strings.Builder
	notFirst := true
	for _, letter := range tonic {
		if notFirst {
			letter = unicode.ToUpper(letter)
			notFirst = false
		}
		b.WriteRune(letter)
	}
	first := b.String()
	for pos := range tones {
		switch first {
		case tones[pos].major, tones[pos].sharp, tones[pos].flat:
			return pos
		}
	}
	return -1
}

// GetTone function get the tone in "pos"
func GetTone(pos int, flat bool) string {
	if tones[pos].major == "" {
		if flat {
			return tones[pos].flat
		}
		return tones[pos].sharp
	}
	return tones[pos].major
}
