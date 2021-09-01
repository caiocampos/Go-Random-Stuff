package clock

import (
	"fmt"
)

// Clock type defines a clock structure
type Clock struct {
	minutes int
}

// New function creates a clock
func New(hours, minutes int) Clock {
	return Clock{normalize(hours*60 + minutes)}
}

// Add method adds minutes to the clock time
func (c Clock) Add(minutes int) Clock {
	return Clock{normalize(c.minutes + minutes)}
}

// Subtract method subtracts minutes to the clock time
func (c Clock) Subtract(minutes int) Clock {
	return Clock{normalize(c.minutes - minutes)}
}

// String method generates the textual representation of the clock
func (c Clock) String() string {
	var hours = c.minutes / 60
	var minutes = c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func normalize(val int) int {
	for val < 0 {
		val += 1440
	}
	return val % 1440
}
