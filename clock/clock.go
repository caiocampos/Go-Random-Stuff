package clock

import (
	"fmt"
	"time"
)

// Clock type defines a clock structure
type Clock struct {
	hour, minute int
}

var start = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

// New function creates a clock
func New(hours, minutes int) Clock {
	temp := start.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes))
	return Clock{temp.Hour(), temp.Minute()}
}

// Add method adds minutes to the clock time
func (c Clock) Add(minutes int) Clock {
	temp := start.Add(time.Hour*time.Duration(c.hour) + time.Minute*time.Duration(c.minute+minutes))
	return Clock{temp.Hour(), temp.Minute()}
}

// Subtract method subtracts minutes to the clock time
func (c Clock) Subtract(minutes int) Clock {
	temp := start.Add(time.Hour*time.Duration(c.hour) + time.Minute*time.Duration(c.minute-minutes))
	return Clock{temp.Hour(), temp.Minute()}
}

// String method generates the textual representation of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
