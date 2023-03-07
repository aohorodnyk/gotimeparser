package gotimeparser_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Since this package is related to time.
	// Many tests use time representation in output.
	// So we set an expected timezone for all tests.
	os.Setenv("TZ", "America/Los_Angeles")

	os.Exit(m.Run())
}
