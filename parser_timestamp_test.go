package gotimeparser_test

import (
	"fmt"

	"github.com/aohorodnyk/gotimeparser"
)

func ExampleParseTimestamp() {
	seconds := int64(1651808102)
	fmt.Println(gotimeparser.ParseTimestamp(seconds))

	milliseconds := int64(1651808102363)
	fmt.Println(gotimeparser.ParseTimestamp(milliseconds))

	microseconds := int64(1651808102363368)
	fmt.Println(gotimeparser.ParseTimestamp(microseconds))

	nanoseconds := int64(1651808102363368423)
	fmt.Println(gotimeparser.ParseTimestamp(nanoseconds))

	// Output:
	// 2022-05-05 20:35:02 -0700 PDT
	// 2022-05-05 20:35:02.363 -0700 PDT
	// 2022-05-05 20:35:02.363368 -0700 PDT
	// 2022-05-05 20:35:02.363368423 -0700 PDT
}
