package gotimeparser

// EmptyInputError represents empty input in Parse function.
type EmptyInputError struct {
	message string
}

// Implementation of the error type for EmptyInputError struct.
func (e EmptyInputError) Error() string {
	return e.message
}

// UnknownFormatError returned when we did not find any supported format to parse the time.
type UnknownFormatError struct {
	message string
}

// Implementation of the error type for UnknownFormatError struct.
func (e UnknownFormatError) Error() string {
	return e.message
}

// UnknownLayoutError returned from `ParseLayouts` and `ParseAllLayouts` when supported layout was not found.
type UnknownLayoutError struct {
	message string
}

// Implementation of the error type for UnknownLayoutError struct.
func (e UnknownLayoutError) Error() string {
	return e.message
}
