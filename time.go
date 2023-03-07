package gotimeparser

import "time"

// Time is a custom Time type that can be used during a parsing process.
// To convert to default time we just need to use `time.Time(timeVariable)`.
type Time time.Time

// UnmarshalJSON implements universal parser for the custom Time type.
func (t *Time) UnmarshalJSON(input []byte) error {
	parsed, err := Parse(string(input))
	if err == nil {
		*t = Time(parsed)
	}

	return err
}
