package gotimeparser

import "time"

// This function converts `time.Time` to `gotimepareser.Time`.
func FromGoTime(t time.Time) Time {
	return Time{t}
}

// Time is a custom Time type that can be used during a parsing process.
// To convert to `time.Time“ use `timeVariable.GoTime()` function.
type Time struct {
	time.Time
}

// UnmarshalJSON implements universal parser for the custom Time type.
func (t *Time) UnmarshalJSON(input []byte) error {
	parsed, err := Parse(string(input))
	if err == nil {
		t.Time = parsed
	}

	return err
}

// Return time.Time from gotimeparser.Time.
func (t *Time) GoTime() time.Time {
	return t.Time
}
