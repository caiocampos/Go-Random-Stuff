package raindrops

import (
	"strconv"
	"strings"
)

// Convert function generates a string on Raindrops pattern from a integer parameter
func Convert(input int) string {
	var b strings.Builder

	if input%3 == 0 {
		b.WriteString("Pling")
	}
	if input%5 == 0 {
		b.WriteString("Plang")
	}
	if input%7 == 0 {
		b.WriteString("Plong")
	}
	if b.Len() == 0 {
		b.WriteString(strconv.Itoa(input))
	}
	return b.String()
}
