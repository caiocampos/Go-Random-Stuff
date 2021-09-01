// Package gigasecond calculates with gigaseconds
package gigasecond

import "time"

var gigasecond = time.Second * 1000000000

// AddGigasecond function adds one gigasecond to the instant "t"
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
