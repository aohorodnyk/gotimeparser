package gotimeparser

import (
	"math"
	"time"
)

// List of times in int64 formats.
const (
	// Positive values, bigger than 0 nanoseconds.
	MaxNanoseconds  = int64(math.MaxInt64)
	MaxMicroseconds = MaxNanoseconds / 1000
	MaxMilliseconds = MaxMicroseconds / 1000
	MaxSeconds      = MaxMilliseconds / 1000

	// Negative values, smaller than 0 nanoseconds.
	MinNanoseconds  = int64(math.MinInt64)
	MinMicroseconds = MinNanoseconds / 1000
	MinMilliseconds = MinMicroseconds / 1000
	MinSeconds      = MinMilliseconds / 1000
)

// ParseTimestamp parses time from int64.
// Support all time formats from nanoseconds to seconds.
// There is a small overlap around 0 between all formats, but this overlap is insignificant.
func ParseTimestamp(timestamp int64) time.Time {
	switch {
	case timestamp < MinMicroseconds:
		return time.Unix(0, timestamp) // Before 1970 in nanoseconds.
	case timestamp < MinMilliseconds:
		return time.Unix(0, timestamp*int64(time.Microsecond)) // Before 1970 in microseconds.
	case timestamp < MinSeconds:
		return time.Unix(0, timestamp*int64(time.Millisecond)) // Before 1970 in milliseconds.
	case timestamp < 0:
		return time.Unix(timestamp, 0) // Before 1970 in seconds.
	case timestamp < MaxSeconds:
		return time.Unix(timestamp, 0) // After 1970 in seconds.
	case timestamp < MaxMilliseconds:
		return time.Unix(0, timestamp*int64(time.Millisecond)) // After 1970 in milliseconds.
	case timestamp < MaxMicroseconds:
		return time.Unix(0, timestamp*int64(time.Microsecond)) // After 1970 in microseconds.
	}

	return time.Unix(0, timestamp) // After 1970 in nanoseconds.
}
