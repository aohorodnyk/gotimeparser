package gotimeparser

import (
	"fmt"
	"time"
)

// DefaultAllLayouts returns a default list of ALL supported layouts.
// This function returns new copy of a slice.
func DefaultAllLayouts() []string {
	return []string{
		time.RFC3339Nano,
		time.RFC3339,
		time.RFC1123Z,
		time.RFC1123,
		time.RFC850,
		time.RFC822Z,
		time.RFC822,
		time.Layout,
		time.RubyDate,
		time.UnixDate,
		time.ANSIC,
		time.StampNano,
		time.StampMicro,
		time.StampMilli,
		time.Stamp,
		time.Kitchen,
	}
}

// ParseAllLayouts uses layouts from `DefaultAllLayouts` to parse the time from a provided input string.
func ParseAllLayouts(dateTime string) (time.Time, error) {
	return ParseLayouts(DefaultAllLayouts(), dateTime)
}

// ParseLayouts parses time based on a list of provided layouts.
// If layouts is empty list or nil, the error with unknown layout will be returned.
func ParseLayouts(layouts []string, dateTime string) (time.Time, error) {
	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, dateTime)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, &UnknownLayoutError{fmt.Sprintf("could not parse time: %s", dateTime)}
}
