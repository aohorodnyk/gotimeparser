package gotimeparser

import (
	"fmt"
	"strconv"
	"time"
)

// Parse function parses time from string.
// It supports all Go-defined layouts and int64 values.
// This function supports strings with `"` and `'` quotes marks, but only valid one, like:
// * "2022-05-06T03:35:02Z"
// * '1651808102'
// * `2022-05-06T03:35:02.363368423Z`
// One quote like "1651808102 or multiple quotes like "'1651808102'" are not supported,
// because of they are not valid strings.
func Parse(input string) (time.Time, error) {
	// Trim quotes, it's important for every future step.
	inputTrimmed := removeStrQuotes(input, supportedQuotes())

	timeInt64, err := strconv.ParseInt(inputTrimmed, 0, 64)
	if err == nil {
		parsed := ParseTimestamp(timeInt64)

		return parsed, nil
	}

	// Parse the string as a time using the supported layouts.
	parsed, err := ParseAllLayouts(inputTrimmed)
	if err == nil {
		return parsed, nil
	}

	return time.Time{}, &UnknownFormatError{fmt.Sprintf("Unsupported time format for string: '%s'", input)}
}

func supportedQuotes() []rune {
	return []rune{'"', '\'', '`'}
}

func removeStrQuotes(input string, quotes []rune) string {
	const minInputLength = 3
	if len(input) < minInputLength {
		return input
	}

	for _, quote := range quotes {
		if rune(input[0]) == quote && rune(input[len(input)-1]) == quote {
			return input[1 : len(input)-1]
		}
	}

	return input
}
